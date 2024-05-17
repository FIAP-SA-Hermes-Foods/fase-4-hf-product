package application

import "fase-4-hf-product/internal/core/domain/entity/dto"

func (app application) GetProductByUUIDUseCase(cpf string) error {
	return app.productUC.GetProductByUUID(cpf)
}

func (app application) SaveProductUseCase(product dto.RequestProduct) error {
	return app.productUC.SaveProduct(product)
}

func (app application) GetProductByCategoryUseCase(category string) error {
	return app.productUC.GetProductByCategory(category)
}

func (app application) UpdateProductByUUIDUseCase(uuid string, product dto.RequestProduct) error {
	return app.productUC.UpdateProductByUUID(uuid, product)
}

func (app application) DeleteProductByUUIDUseCase(uuid string) error {
	return app.productUC.DeleteProductByUUID(uuid)
}
