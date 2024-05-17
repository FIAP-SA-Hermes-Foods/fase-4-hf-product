package repository

import (
	"fase-4-hf-product/internal/core/domain/entity/dto"
)

type ProductRepository interface {
	GetProductByUUID(uuid string) (*dto.ProductDB, error)
	SaveProduct(product dto.ProductDB) (*dto.ProductDB, error)
	UpdateProductByUUID(uuid string, product dto.ProductDB) (*dto.ProductDB, error)
	GetProductByCategory(category string) ([]dto.ProductDB, error)
	DeleteProductByUUID(uuid string) error
}
