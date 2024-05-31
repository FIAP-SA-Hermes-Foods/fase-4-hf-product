package rpc

import (
	"context"
	"fase-4-hf-product/internal/core/application"
	"fase-4-hf-product/internal/core/domain/entity/dto"
	cp "fase-4-hf-product/product_proto"
)

type HandlerGRPC interface {
	Handler() *handlerGRPC
}

type handlerGRPC struct {
	app application.Application
	cp.UnimplementedProductServer
}

func NewHandler(app application.Application) HandlerGRPC {
	return &handlerGRPC{app: app}
}

func (h *handlerGRPC) Handler() *handlerGRPC {
	return h
}

func (h *handlerGRPC) GetProductByID(ctx context.Context, req *cp.GetProductByIDRequest) (*cp.GetProductByIDResponse, error) {
	c, err := h.app.GetProductByID(req.Uuid)

	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, nil
	}

	out := &cp.GetProductByIDResponse{
		Uuid:          c.UUID,
		Name:          c.Name,
		Category:      c.Category,
		Image:         c.Image,
		Description:   c.Description,
		Price:         float32(c.Price),
		CreatedAt:     c.CreatedAt,
		DeactivatedAt: c.DeactivatedAt,
	}

	return out, nil
}

func (h *handlerGRPC) CreateProduct(ctx context.Context, req *cp.CreateProductRequest) (*cp.CreateProductResponse, error) {
	input := dto.RequestProduct{
		Name:          req.Name,
		Category:      req.Category,
		Image:         req.Image,
		Description:   req.Description,
		Price:         float64(req.Price),
		CreatedAt:     req.CreatedAt,
		DeactivatedAt: req.DeactivatedAt,
	}

	c, err := h.app.SaveProduct(input)

	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, nil
	}

	out := &cp.CreateProductResponse{
		Uuid:          c.UUID,
		Name:          c.Name,
		Category:      c.Category,
		Image:         c.Image,
		Description:   c.Description,
		Price:         float32(c.Price),
		CreatedAt:     c.CreatedAt,
		DeactivatedAt: c.DeactivatedAt,
	}

	return out, nil
}

func (h *handlerGRPC) GetProductByCategory(ctx context.Context, req *cp.GetProductByCategoryRequest) (*cp.GetProductByCategoryResponse, error) {
	c, err := h.app.GetProductByCategory(req.Category)

	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, nil
	}

	outList := make([]*cp.GetProductByCategoryItem, 0)

	for i := range c {
		item := &cp.GetProductByCategoryItem{
			Uuid:          c[i].UUID,
			Name:          c[i].Name,
			Category:      c[i].Category,
			Image:         c[i].Image,
			Description:   c[i].Description,
			Price:         float32(c[i].Price),
			CreatedAt:     c[i].CreatedAt,
			DeactivatedAt: c[i].DeactivatedAt,
		}

		outList = append(outList, item)
	}

	out := &cp.GetProductByCategoryResponse{
		Items: outList,
	}

	return out, nil
}

func (h *handlerGRPC) UpdateProduct(ctx context.Context, req *cp.UpdateProductRequest) (*cp.UpdateProductResponse, error) {
	input := dto.RequestProduct{
		Name:          req.Name,
		Category:      req.Category,
		Image:         req.Image,
		Description:   req.Description,
		Price:         float64(req.Price),
		CreatedAt:     req.CreatedAt,
		DeactivatedAt: req.DeactivatedAt,
	}

	c, err := h.app.UpdateProductByID(req.Uuid, input)

	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, nil
	}

	out := &cp.UpdateProductResponse{
		Uuid:          c.UUID,
		Name:          c.Name,
		Category:      c.Category,
		Image:         c.Image,
		Description:   c.Description,
		Price:         float32(c.Price),
		CreatedAt:     c.CreatedAt,
		DeactivatedAt: c.DeactivatedAt,
	}

	return out, nil
}

func (h *handlerGRPC) DeleteProductByID(ctx context.Context, req *cp.DeleteProductByIDRequest) (*cp.DeleteProductByIDResponse, error) {

	if err := h.app.DeleteProductByID(req.Uuid); err != nil {
		return nil, err
	}

	out := &cp.DeleteProductByIDResponse{
		Message: "OK",
	}

	return out, nil

}
