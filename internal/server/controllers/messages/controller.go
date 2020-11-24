package messages

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"

	"mailarc/internal/mimecontent"
	"mailarc/internal/store"
)

type controller struct {
	mailbox store.ReadStore
	handler http.Handler
}

func New(mailbox store.ReadStore, router *mux.Router) http.Handler {
	c := &controller{mailbox: mailbox, handler: router}
	router.HandleFunc("/", c.ListMessages)
	router.HandleFunc("/{uid}", c.GetMessage)
	router.HandleFunc("/{uid}/", c.GetMessage)
	router.HandleFunc("/{uid}/rendered", c.GetMessageRendered)
	return c
}

func (c *controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.handler.ServeHTTP(w, r)
}

func (c *controller) ListMessages(w http.ResponseWriter, r *http.Request) {
	messages, err := c.mailbox.List()
	if err != nil {
		panic(err)
	}
	if err := tmplListMessages.Execute(w, messages); err != nil {
		panic(err)
	}
}

func (c *controller) GetMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]
	mr, err := c.mailbox.Open(uid)
	if err != nil {
		if err == store.ErrNotFound {
			http.NotFound(w, r)
			return
		}
		renderServerError(w, r, err)
		return
	}

	// decode
	br := bufio.NewReader(mr)
	content, err := mimecontent.Read(br)
	if err != nil {
		renderServerError(w, r, err)
		return
	}

	// render html
	err = tmplGetMessage.Execute(w, NewViewMessage(uid, content))
	if err != nil {
		renderServerError(w, r, err)
		return
	}
}

func (c *controller) GetMessageRendered(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]
	mr, err := c.mailbox.Open(uid)
	if err != nil {
		renderServerError(w, r, err)
		return
	}

	// decode
	br := bufio.NewReader(mr)
	content, err := mimecontent.Read(br)
	if err != nil {
		renderServerError(w, r, err)
		return
	}

	// get document root
	for content.IsMultipart() {
		content = content.Root()
	}

	// write
	if s := content.ContentType(); s != "" {
		w.Header().Set("Content-Type", s)
	} else {
		w.Header().Set("Content-Type", "text/plain")
	}
	if _, err = io.Copy(w, bytes.NewReader(content.Body)); err != nil {
		panic(err)
	}
}

func parseView(name, tmpl string) *template.Template {
	t := template.New(name)
	t = t.Funcs(template.FuncMap{
		"decode":      tmplFuncDecode,
		"decodeSlice": tmplFuncDecodeSlice,
		"isText":      tmplFuncIsText,
		"isImage":     tmplFuncIsImage,
	})
	t = template.Must(t.Parse(tmpl))
	t = template.Must(t.Parse(ComponentHeaderTable))
	t = template.Must(t.Parse(LayoutDefault))
	return t
}

func renderServerError(w http.ResponseWriter, r *http.Request, err error) {
	status := http.StatusInternalServerError
	http.Error(w, http.StatusText(status), status)
	log.Print(err)
}

var tmplListMessages = parseView("ListMessages", ViewListMessages)
