package router

import (
	"github.com/Justdanru/bhs-test/internal/controller/http/v1/handler"
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	mux         *mux.Router
	rootHandler *handler.RootHandler
}

func NewRouter(
	rootHandler *handler.RootHandler,
) *Router {
	newMux := mux.NewRouter()

	newMux.HandleFunc("/users/{user_id}", rootHandler.User.Handle).Methods(http.MethodGet)

	return &Router{
		mux:         newMux,
		rootHandler: rootHandler,
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
