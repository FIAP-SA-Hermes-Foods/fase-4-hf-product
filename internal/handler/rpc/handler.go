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

func (h *handlerGRPC) Create(ctx context.Context, req *cp.CreateRequest) (*cp.CreateResponse, error) {
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

	out := &cp.CreateResponse{
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

func (h *handlerGRPC) GetByCategory(ctx context.Context, req *cp.GetByCategoryRequest) (*cp.GetByCategoryResponse, error) {

	c, err := h.app.GetProductByCategory(req.Category)

	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, nil
	}

	outList := make([]*cp.GetBycategoryItem, 0)

	for i := range c {
		item := &cp.GetBycategoryItem{
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

	out := &cp.GetByCategoryResponse{
		Items: outList,
	}

	return out, nil
}

func (h *handlerGRPC) Update(ctx context.Context, req *cp.UpdateRequest) (*cp.UpdateResponse, error) {
	input := dto.RequestProduct{
		Name:          req.Name,
		Category:      req.Category,
		Image:         req.Image,
		Description:   req.Description,
		Price:         float64(req.Price),
		CreatedAt:     req.CreatedAt,
		DeactivatedAt: req.DeactivatedAt,
	}

	c, err := h.app.UpdateProductByUUID(req.Uuid, input)

	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, nil
	}

	out := &cp.UpdateResponse{
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

func (h *handlerGRPC) DeleteByUUID(ctx context.Context, req *cp.DeleteByUUIDRequest) (*cp.DeleteByUUIDResponse, error) {

	if err := h.app.DeleteProductByUUID(req.Uuid); err != nil {
		return nil, err
	}

	out := &cp.DeleteByUUIDResponse{
		Message: "OK",
	}

	return out, nil

}
