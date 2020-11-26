package index

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"mailarc/internal/index"
)

type controller struct {
	index   index.Index
	handler http.Handler
}

func New(ix index.Index, router *mux.Router) http.Handler {
	c := &controller{index: ix, handler: router}
	router.HandleFunc("/", c.Search)
	return c
}

func (c *controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.handler.ServeHTTP(w, r)
}

func (c *controller) Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// page offset
	s := r.URL.Query().Get("offset")
	if s == "" {
		s = "0"
	}
	offset, _ := strconv.Atoi(s)

	// page limit
	s = r.URL.Query().Get("limit")
	if s == "" {
		s = "100000"
	}
	limit, _ := strconv.Atoi(s)

	// search
	q := r.URL.Query().Get("q") // default ""
	results, err := c.index.Search(context.TODO(), q, offset, limit)
	if err != nil {
		panic(err) // TODO
	}
	enc := json.NewEncoder(w)
	enc.Encode(results)
}
