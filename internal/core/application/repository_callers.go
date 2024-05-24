package application

import "fase-4-hf-product/internal/core/domain/entity/dto"

func (app application) GetProductByUUIDRepository(uuid string) (*dto.ProductDB, error) {
	return app.productRepo.GetProductByUUID(uuid)
}

func (app application) SaveProductRepository(product dto.ProductDB) (*dto.ProductDB, error) {
	return app.productRepo.SaveProduct(product)
}

func (app application) GetProductByCategoryRepository(category string) ([]dto.ProductDB, error) {
	return app.productRepo.GetProductByCategory(category)
}

func (app application) UpdateProductByUUIDRepository(uuid string, product dto.ProductDB) (*dto.ProductDB, error) {
	return app.productRepo.UpdateProductByUUID(uuid, product)
}

func (app application) DeleteProductByUUIDRepository(uuid string) error {
	return app.productRepo.DeleteProductByUUID(uuid)
}
