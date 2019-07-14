package common

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	mux *mux.Router
}

func (r *Router) Mux() *mux.Router {
	return r.mux
}

func NewRouter() *Router {
	return &Router{mux: mux.NewRouter().StrictSlash(true)}
}

func newRoute(router *mux.Router, method string, pattern string, name string, callback http.HandlerFunc) {
	router.Methods(method).Path(pattern).Name(name).Handler(callback)
}

func (r *Router) RegisterRoutes(routes Routes) {
	for _, route := range routes {
		newRoute(r.mux, route.Method, route.Pattern, route.Name, route.Handler)
	}
}
