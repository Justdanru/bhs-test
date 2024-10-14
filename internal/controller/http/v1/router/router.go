package router

import (
	"github.com/Justdanru/bhs-test/internal/controller/http/v1/handler"
	"github.com/gorilla/mux"
	"net/http"
	"sync/atomic"
)

type Router struct {
	mux         *mux.Router
	rootHandler *handler.RootHandler
	requestId   atomic.Uint64
}

func NewRouter(
	rootHandler *handler.RootHandler,
) *Router {
	router := &Router{
		rootHandler: rootHandler,
		requestId:   atomic.Uint64{},
	}

	newMux := mux.NewRouter()

	newMux.Use(router.initMiddleware)

	newMux.HandleFunc("/users/{user_id}", rootHandler.User.Handle).Methods(http.MethodGet)

	newMux.HandleFunc("/users", rootHandler.Register.Handle).Methods(http.MethodPost)

	router.mux = newMux

	return router
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
