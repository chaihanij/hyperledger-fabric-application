package router

import (
	"github.com/chaihanij/hyperledger-fabric-application/back-end/models"
	"github.com/gorilla/mux"
)

type Router struct {
	RawRouter *mux.Router
}

func (r Router) GetRawRouter() *mux.Router {
	return r.RawRouter
}

func (r *Router) AttachSubRouterWithMiddleware(path string, subroutes models.Routes, middleware mux.MiddlewareFunc) *mux.Router {

	SubRouter := r.RawRouter.PathPrefix(path).Subrouter()
	SubRouter.Use(middleware)

	for _, route := range subroutes {
		SubRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return SubRouter
}
