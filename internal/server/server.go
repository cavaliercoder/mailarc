package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"mailarc/internal/index"
	messages_controller "mailarc/internal/server/controllers/messages"
	store_controller "mailarc/internal/server/controllers/store"
	"mailarc/internal/store"
)

func New(mailbox store.Store, ix index.Index) http.Handler {
	r := mux.NewRouter()
	r.Handle("/", http.RedirectHandler("/messages/", http.StatusTemporaryRedirect))
	store_controller.New(mailbox, r.PathPrefix("/store").Subrouter())
	messages_controller.New(mailbox, ix, r.PathPrefix("/messages").Subrouter())
	return r
}
