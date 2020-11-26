package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"mailarc/internal/index"
	json_search_controller "mailarc/internal/server/controllers/json-rpc/search"
	messages_controller "mailarc/internal/server/controllers/messages"
	store_controller "mailarc/internal/server/controllers/store"
	"mailarc/internal/store"
)

func New(mailbox store.Store, ix index.Index) http.Handler {
	r := mux.NewRouter()
	r.Handle("/", http.RedirectHandler("/messages/", http.StatusTemporaryRedirect))
	store_controller.New(mailbox, r.PathPrefix("/store").Subrouter())
	messages_controller.New(mailbox, ix, r.PathPrefix("/messages").Subrouter())

	// json-roc
	jsonRpcRouter := r.PathPrefix("/json-rpc").Subrouter()
	json_search_controller.New(ix, jsonRpcRouter.PathPrefix("/search").Subrouter())

	// r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	// 	t, err := route.GetPathTemplate()
	// 	if err != nil {
	// 		return err
	// 	}
	// 	fmt.Println(t)
	// 	return nil
	// })
	return r
}
