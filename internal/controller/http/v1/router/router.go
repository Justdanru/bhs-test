package router

import (
	"github.com/Justdanru/bhs-test/internal/controller/http/v1/handler"
	"github.com/Justdanru/bhs-test/internal/controller/http/v1/middleware"
	"github.com/gorilla/mux"
	"net/http"
	"sync/atomic"
)

type Router struct {
	mux            *mux.Router
	rootHandler    *handler.RootHandler
	rootMiddleware *middleware.RootMiddleware
	requestId      atomic.Uint64
}

func NewRouter(
	rootHandler *handler.RootHandler,
	rootMiddleware *middleware.RootMiddleware,
) *Router {
	router := &Router{
		rootHandler:    rootHandler,
		rootMiddleware: rootMiddleware,
		requestId:      atomic.Uint64{},
	}

	baseMux := mux.NewRouter()
	baseMux.Use(rootMiddleware.Init.Handle)

	baseMux.HandleFunc("/check_username", rootHandler.CheckUsername.Handle).Methods(http.MethodPost)
	baseMux.HandleFunc("/users", rootHandler.Register.Handle).Methods(http.MethodPost)
	baseMux.HandleFunc("/login", rootHandler.Login.Handle).Methods(http.MethodPost)

	authMux := baseMux.PathPrefix("/").Subrouter()
	authMux.Use(rootMiddleware.Auth.Handle)

	authMux.HandleFunc("/users/{user_id}", rootHandler.User.Handle).Methods(http.MethodGet)

	router.mux = baseMux

	return router
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
