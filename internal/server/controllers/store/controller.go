package store

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/cavaliercoder/mailarc/internal/mimecontent"
	"github.com/cavaliercoder/mailarc/internal/store"
	"github.com/gorilla/mux"
)

type controller struct {
	mailbox store.ReadStore
	handler http.Handler
}

func New(mailbox store.ReadStore, router *mux.Router) http.Handler {
	c := &controller{mailbox: mailbox, handler: router}
	router.HandleFunc("/messages/{uid}", c.GetMessage)
	router.HandleFunc("/messages/{uid}/parts/{path:[0-9]+(?:/[0-9]+)*}", c.GetMessagePart)
	return c
}

func (c *controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.handler.ServeHTTP(w, r)
}

func (c *controller) GetMessage(w http.ResponseWriter, r *http.Request) {
	c.getMessage(w, r, false)
}

func (c *controller) getMessage(w http.ResponseWriter, r *http.Request, asAttachment bool) {
	vars := mux.Vars(r)
	uid := vars["uid"]

	mr, err := c.mailbox.Open(uid)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
		panic(err)
	}
	defer mr.Close()

	if asAttachment {
		w.Header().Set("Content-Type", "message/rfc822")
		w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s.eml"`, uid))
	} else {
		w.Header().Set("Content-Type", "text/plain")
	}

	if _, err = io.Copy(w, mr); err != nil {
		panic(err) // TODO
	}
}

func (c *controller) GetMessagePart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]
	path := vars["path"]

	mr, err := c.mailbox.Open(uid)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
		// renderServerError(w, r, err)
		panic(err)
	}
	defer mr.Close()

	// parse message
	content, err := mimecontent.Read(bufio.NewReader(mr))
	if err != nil {
		// renderServerError(w, r, err)
		panic(err)
	}

	// extract part
	content = content.GetPart(path)
	if content == nil {
		log.Print("not found: ", r.URL)
		http.NotFound(w, r)
		return
	}

	// print
	if s := content.ContentType(); s != "" {
		w.Header().Set("Content-Type", s)
	}
	if _, err := w.Write(content.Body); err != nil {
		panic(err) // TODO
	}
}
