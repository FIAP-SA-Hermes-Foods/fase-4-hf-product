package application

import "fase-4-hf-product/internal/core/domain/entity/dto"

func (app application) GetProductByIDUseCase(cpf string) error {
	return app.productUC.GetProductByID(cpf)
}

func (app application) SaveProductUseCase(product dto.RequestProduct) error {
	return app.productUC.SaveProduct(product)
}

func (app application) GetProductByCategoryUseCase(category string) error {
	return app.productUC.GetProductByCategory(category)
}

func (app application) UpdateProductByIDUseCase(uuid string, product dto.RequestProduct) error {
	return app.productUC.UpdateProductByID(uuid, product)
}

func (app application) DeleteProductByIDUseCase(uuid string) error {
	return app.productUC.DeleteProductByID(uuid)
}
