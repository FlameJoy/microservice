package handlers

import (
	"microAPI/data"
	logging "microAPI/logger"
	"microAPI/router"
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

func (h *ProductHandler) RegisterRoutes(rg *router.RouteGroup) {
	products := rg.Group("/products")
	products.GET("", h.GetProducts)
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()

	err := products.ToJSON(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
