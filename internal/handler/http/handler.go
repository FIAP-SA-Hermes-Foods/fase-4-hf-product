package http

import (
	"bytes"
	"encoding/json"
	ps "fase-4-hf-product/external/strings"
	"fase-4-hf-product/internal/core/application"
	"fase-4-hf-product/internal/core/domain/entity/dto"
	"fmt"
	"net/http"
	"time"
)

type ProductHandler interface {
	Handler(rw http.ResponseWriter, req *http.Request)
	HealthCheck(rw http.ResponseWriter, req *http.Request)
}

type productHandler struct {
	app application.Application
}

func NewHandler(app application.Application) ProductHandler {
	return productHandler{app: app}
}

func (h productHandler) Handler(rw http.ResponseWriter, req *http.Request) {

	var routeProducts = map[string]http.HandlerFunc{
		"get hermes_foods/product":           h.getProductByCategory,
		"post hermes_foods/product":          h.saveProduct,
		"put hermes_foods/product/{uuid}":    h.UpdateProductByUUID,
		"delete hermes_foods/product/{uuid}": h.deleteProductByUUID,
	}

	handler, err := router(req.Method, req.URL.Path, routeProducts)

	if err == nil {
		handler(rw, req)
		return
	}

	rw.WriteHeader(http.StatusNotFound)
	rw.Write([]byte(`{"error": "route ` + req.Method + " " + req.URL.Path + ` not found"} `))
}

func (h productHandler) HealthCheck(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte(`{"error": "method not allowed"}`))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{"status": "OK"}`))
}

func (h *productHandler) saveProduct(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte(`{"error": "method not allowed"} `))
		return
	}

	var buff bytes.Buffer

	var reqProduct dto.RequestProduct

	if _, err := buff.ReadFrom(req.Body); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to read data body: %v"} `, err)
		return
	}

	if err := json.Unmarshal(buff.Bytes(), &reqProduct); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to Unmarshal: %v"} `, err)
		return
	}

	p, err := h.app.SaveProduct(reqProduct)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to save product: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(ps.MarshalString(p)))
}

func (h *productHandler) UpdateProductByUUID(rw http.ResponseWriter, req *http.Request) {
	id := getID("product", req.URL.Path)

	var buff bytes.Buffer

	var reqProduct dto.RequestProduct

	if _, err := buff.ReadFrom(req.Body); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to read data body: %v"} `, err)
		return
	}

	if err := json.Unmarshal(buff.Bytes(), &reqProduct); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to Unmarshal: %v"} `, err)
		return
	}

	product := reqProduct.Product()

	if len(reqProduct.DeactivatedAt) > 0 {
		product.DeactivatedAt.Value = new(time.Time)
		if err := product.DeactivatedAt.SetTimeFromString(reqProduct.DeactivatedAt); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(rw, `{"error": "error to update product: %v"} `, err)
			return
		}
	}

	reqProduct.DeactivatedAt = product.DeactivatedAt.Format()

	p, err := h.app.UpdateProductByUUID(id, reqProduct)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to update product: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(ps.MarshalString(p)))
}

func (h *productHandler) deleteProductByUUID(rw http.ResponseWriter, req *http.Request) {
	id := getID("product", req.URL.Path)

	if req.Method != http.MethodDelete {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte(`{"error": "method not allowed"} `))
		return
	}

	if err := h.app.DeleteProductByUUID(id); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to delete product: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{"status":"OK"}`))
}

func (h *productHandler) getProductByCategory(rw http.ResponseWriter, req *http.Request) {
	category := req.URL.Query().Get("category")

	pList, err := h.app.GetProductByCategory(category)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to get product by category: %v"} `, err)
		return
	}

	if pList == nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte(`{"error": "product not found"}`))
		return
	}

	b, err := json.Marshal(pList)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to get product by category: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(b)
}
