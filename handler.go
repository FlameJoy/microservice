package main

import "net/http"

type handler struct {
	logger *CustomLogger
}

func NewHandler(l *CustomLogger) *handler {
	return &handler{
		logger: l,
	}
}

func (h *handler) testHandler(w http.ResponseWriter, r *http.Request) {

}
