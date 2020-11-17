package server

import (
	"net/http"

	messages_controller "github.com/cavaliercoder/mailarc/internal/server/controllers/messages"
	store_controller "github.com/cavaliercoder/mailarc/internal/server/controllers/store"
	"github.com/cavaliercoder/mailarc/internal/store"
	"github.com/gorilla/mux"
)

func New(mailbox store.Store) http.Handler {
	r := mux.NewRouter()
	r.Handle("/", http.RedirectHandler("/messages/", http.StatusTemporaryRedirect))
	store_controller.New(mailbox, r.PathPrefix("/store").Subrouter())
	messages_controller.New(mailbox, r.PathPrefix("/messages").Subrouter())
	return r
}
