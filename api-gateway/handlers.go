package main

import (
	"microsvc/common/utils"
	"net/http"
)

type handler struct {
	logger *utils.CustomLogger
}

func NewHandler(l *utils.CustomLogger) *handler {
	return &handler{
		logger: l,
	}
}

func (h *handler) testHandler(w http.ResponseWriter, r *http.Request) {
	// err := r.ParseMultipartForm()
	// if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {

	// }
}
