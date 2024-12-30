package main

import (
	"microAPI/handlers"
	logging "microAPI/logger"
	"net/http"
	"strings"
)

func registerHandler(mux *http.ServeMux, logger *logging.CustomLogger) {
	ph := handlers.NewProductHandler(logger)
	mux.HandleFunc("/product", ph.TestHandler)

	products := NewRouteGroup("/products", mux, logger)
	products.POST("/list", ph.GetProducts)
}

type RouteGroup struct {
	prefix string
	parent *RouteGroup
	mws    []Middleware
	mux    *http.ServeMux
	logger *logging.CustomLogger
}

func NewRouteGroup(prefix string, mux *http.ServeMux, logger *logging.CustomLogger) *RouteGroup {
	return &RouteGroup{
		prefix: strings.TrimRight(prefix, "/"),
		mux:    mux,
		logger: logger,
	}
}

func (rg *RouteGroup) Use(mws ...Middleware) {
	rg.mws = append(mws, mws...)
}

func (rg *RouteGroup) Group(pattern string) *RouteGroup {
	return &RouteGroup{
		prefix: strings.TrimRight(rg.prefix+pattern, "/"),
		mux:    rg.mux,
		parent: rg,
	}
}

func (rg *RouteGroup) GET(pattern string, handler http.HandlerFunc) {
	rg.Handle(http.MethodGet, pattern, handler)
}

func (rg *RouteGroup) POST(pattern string, handler http.HandlerFunc) {
	rg.Handle(http.MethodPost, pattern, handler)
}

func (rg *RouteGroup) PUT(pattern string, handler http.HandlerFunc) {
	rg.Handle(http.MethodPut, pattern, handler)
}

func (rg *RouteGroup) DELETE(pattern string, handler http.HandlerFunc) {
	rg.Handle(http.MethodDelete, pattern, handler)
}

func (rg *RouteGroup) CollectMW() []Middleware {
	mws := []Middleware{}
	current := rg

	for current != nil {
		mws = append(mws, current.mws...)
		current = current.parent
	}

	return mws
}

func (rg *RouteGroup) Handle(method, pattern string, handler http.Handler) {
	finalPath := strings.TrimRight(rg.prefix+pattern, "/")
	rg.mux.HandleFunc(finalPath, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			rg.logger.Warn("Method not allowed")
			return
		}

		mws := rg.CollectMW()

		finalHandler := handler

		for i := len(mws) - 1; i >= 0; i-- {
			finalHandler = mws[i](finalHandler)
		}

		finalHandler.ServeHTTP(w, r)
	})
}
