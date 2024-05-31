package useCase

import "fase-4-hf-product/internal/core/domain/entity/dto"

type ProductUseCase interface {
	SaveProduct(reqProduct dto.RequestProduct) error
	GetProductByID(uuid string) error
	UpdateProductByID(uuid string, product dto.RequestProduct) error
	GetProductByCategory(category string) error
	DeleteProductByID(uuid string) error
}
