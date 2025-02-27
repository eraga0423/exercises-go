package router

import (
	"github.com/my/repo/internal/rest/handler"
	"net/http"
)

type Router struct {
	router  *http.ServeMux
	handler *handler.Handler
}

func NewRouter(handler *handler.Handler) *Router {
	mux := http.NewServeMux()
	return &Router{
		router:  mux,
		handler: handler,
	}
}
