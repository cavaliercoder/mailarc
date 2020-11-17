package s3

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"mime"
	"net/url"
	"strings"

	aws "github.com/aws/aws-sdk-go/aws"
	awserr "github.com/aws/aws-sdk-go/aws/awserr"
	aws_session "github.com/aws/aws-sdk-go/aws/session"
	aws_s3 "github.com/aws/aws-sdk-go/service/s3"

	"github.com/cavaliercoder/mailarc/internal/store"
)

const suffix = ".eml.gz"

type s3Store struct {
	session *aws_session.Session
	s3      *aws_s3.S3
	bucket  string
	prefix  string
}

func New(s3url string) (store.Store, error) {
	session, err := aws_session.NewSessionWithOptions(aws_session.Options{
		SharedConfigState: aws_session.SharedConfigEnable,
	})
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(s3url)
	if err != nil {
		return nil, err
	}
	prefix := u.Path
	if strings.HasPrefix(prefix, "/") {
		prefix = prefix[1:]
	}
	if !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}
	c := &s3Store{
		session: session,
		s3:      aws_s3.New(session),
		bucket:  u.Hostname(),
		prefix:  prefix,
	}
	c.Logf("Connected to %s", s3url)
	return c, nil
}

func (c *s3Store) Logf(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	log.Printf("[S3] %s", s)
}

func (c *s3Store) getKey(name string) string {
	return fmt.Sprintf("%s%s%s", c.prefix, name, suffix)
}

func (c *s3Store) getS3URL(name string) string {
	return fmt.Sprintf("s3://%s/%s", c.bucket, c.getKey(name))
}

func (c *s3Store) List() ([]string, error) {
	req := &aws_s3.ListObjectsInput{
		Bucket: aws.String(c.bucket),
		Prefix: aws.String(c.prefix),
	}
	a := make([]string, 0, 64)
	err := c.s3.ListObjectsPages(
		req,
		func(resp *aws_s3.ListObjectsOutput, last bool) (shouldContinue bool) {
			for _, obj := range resp.Contents {
				key := *obj.Key
				if !strings.HasSuffix(key, suffix) {
					continue
				}
				key = key[len(c.prefix):]        // strip prefix
				key = key[:len(key)-len(suffix)] // strip suffix
				a = append(a, key)
			}
			return true
		})
	if err != nil {
		panic(err)
	}
	c.Logf("List() -> %d messages", len(a))
	return a, nil
}

func (c *s3Store) Exists(name string) (bool, error) {
	s3 := aws_s3.New(c.session)
	_, err := s3.HeadObject(
		&aws_s3.HeadObjectInput{
			Bucket: aws.String(c.bucket),
			Key:    aws.String(c.getKey(name)),
		})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			if aerr.Code() == "NotFound" {
				return false, nil
			}
		}
		return false, err
	}
	return true, nil
}

func (c *s3Store) Open(name string) (r io.ReadCloser, err error) {
	c.Logf("Fetching msg %s from %s", name, c.getS3URL(name))
	resp, err := c.s3.GetObject(
		&aws_s3.GetObjectInput{
			Bucket: aws.String(c.bucket),
			Key:    aws.String(c.getKey(name)),
		})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			if aerr.Code() == aws_s3.ErrCodeNoSuchKey {
				return nil, store.ErrNotFound
			}
		}
		return
	}
	return resp.Body, nil // s3 client will gunzip
}

func md5sum(r io.Reader) (s string, err error) {
	h := md5.New()
	if _, err = io.Copy(h, r); err != nil {
		return "", err
	}
	b := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(b), nil
}

func (c *s3Store) Store(name string, r io.Reader) (err error) {
	c.Logf("Uploading msg %s to %s", name, c.getS3URL(name))

	// S3 needs to be able to call r.Seek, and we need an MD5 sum, so copy and
	// compress the whole message into a buffer.
	buf := &bytes.Buffer{}
	gzw, err := gzip.NewWriterLevel(buf, gzip.BestCompression)
	if err != nil {
		return err
	}
	if _, err = io.Copy(gzw, r); err != nil {
		return err
	}
	if err = gzw.Close(); err != nil {
		return err
	}
	body := buf.Bytes()

	// compute md5
	h, err := md5sum(bytes.NewReader(body))
	if err != nil {
		return err
	}

	// upload
	contentDisposition := mime.FormatMediaType(
		"attachment",
		map[string]string{"filename": c.getKey(name)},
	)
	_, err = c.s3.PutObject(
		&aws_s3.PutObjectInput{
			Bucket:             aws.String(c.bucket),
			Key:                aws.String(c.getKey(name)),
			Body:               bytes.NewReader(body),
			ACL:                aws.String("private"),
			ContentLength:      aws.Int64(int64(len(body))),
			ContentType:        aws.String("message/rfc822"),
			ContentEncoding:    aws.String("gzip"),
			ContentDisposition: aws.String(contentDisposition),
			ContentMD5:         aws.String(h),
		})
	return err
}

func (c *s3Store) Delete(name string) (err error) {
	c.Logf("Deleting msg %s from %s", name, c.getS3URL(name))
	_, err = c.s3.DeleteObject(&aws_s3.DeleteObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(c.getKey(name)),
	})
	return
}
