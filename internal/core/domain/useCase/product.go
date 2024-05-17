package useCase

import "fase-4-hf-product/internal/core/domain/entity/dto"

type ProductUseCase interface {
	SaveProduct(reqProduct dto.RequestProduct) error
	GetProductByUUID(uuid string) error
	UpdateProductByUUID(uuid string, product dto.RequestProduct) error
	GetProductByCategory(category string) error
	DeleteProductByUUID(uuid string) error
}
