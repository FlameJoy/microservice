package main

import (
	"microsvc/common/utils"
	"microsvc/middleware"
	"net/http"
	"strings"
)

type RouteGroup struct {
	mux    *http.ServeMux
	prefix string
	parent *RouteGroup
	mws    []middleware.Middleware
	logger *utils.CustomLogger
}

func NewRouteGroup(mux *http.ServeMux, prefix string, logger *utils.CustomLogger) *RouteGroup {
	return &RouteGroup{
		mux:    mux,
		prefix: strings.TrimRight(prefix, "/"),
		logger: logger,
	}
}

func (rg *RouteGroup) Use(mw middleware.Middleware) {
	rg.mws = append(rg.mws, mw)
}

func (rg *RouteGroup) Group(pattern string) *RouteGroup {
	return &RouteGroup{
		mux:    rg.mux,
		prefix: strings.TrimRight(rg.prefix+pattern, "/"),
		parent: rg,
		logger: rg.logger,
	}
}

func (rg *RouteGroup) CollectMW() []middleware.Middleware {
	current := rg
	mws := []middleware.Middleware{}

	for current != nil {
		mws = append(mws, current.mws...)
		current = current.parent
	}

	return mws
}

func (rg *RouteGroup) Handle(method, pattern string, h http.Handler) {
	fullPath := strings.TrimRight(rg.prefix+pattern, "/")
	rg.mux.HandleFunc(fullPath, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			rg.logger.Error("Method %s not allowed", method)
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		finalHandler := h

		mws := rg.CollectMW()

		for i := len(mws) - 1; i >= 0; i-- {
			finalHandler = mws[i](finalHandler)
		}

		finalHandler.ServeHTTP(w, r)
	})
}

func (rg *RouteGroup) GET(pattern string, h http.HandlerFunc) {
	rg.Handle(http.MethodGet, pattern, h)
}

func (rg *RouteGroup) POST(prefix string, handler http.HandlerFunc) {
	rg.Handle(http.MethodPost, prefix, handler)
}

func (rg *RouteGroup) PUT(prefix string, handler http.HandlerFunc) {
	rg.Handle(http.MethodPut, prefix, handler)
}

func (rg *RouteGroup) DELETE(prefix string, handler http.HandlerFunc) {
	rg.Handle(http.MethodDelete, prefix, handler)
}

func registerHandlers(h *handler, mux *http.ServeMux) {

	// API gatewai
	// api := NewRouteGroup(mux, "/api/v1", h.logger)

	// Auth/Reg
	// auth := api.Group("/auth")
	// auth.POST("", h.ProxyAuthReq)
	// auth.POST("/register", h.ProxyRegReq)
	// auth.POST("/recovery", h.ProxyRecoveryReq)
}
