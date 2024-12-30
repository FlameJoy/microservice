package handlers

import (
	"microAPI/data"
	logging "microAPI/logger"
	"net/http"
)

type ProductHandler struct {
	logger *logging.CustomLogger
}

func NewProductHandler(l *logging.CustomLogger) *ProductHandler {
	return &ProductHandler{
		logger: l,
	}
}

func (h *ProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func (h *ProductHandler) TestHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()

	err := products.ToJSON(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
