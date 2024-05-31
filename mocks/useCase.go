package mocks

import (
	"fase-4-hf-product/internal/core/domain/entity/dto"
	"strings"
)

type MockProductUseCase struct {
	WantOutNull string
	WantErr     error
}

func (m MockProductUseCase) GetProductByID(cpf string) error {
	if m.WantErr != nil && strings.EqualFold("errGetProductByID", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m MockProductUseCase) SaveProduct(product dto.RequestProduct) error {
	if m.WantErr != nil && strings.EqualFold("errSaveProduct", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m MockProductUseCase) GetProductByCategory(category string) error {
	if m.WantErr != nil && strings.EqualFold("errGetProductByCategory", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m MockProductUseCase) UpdateProductByID(uuid string, product dto.RequestProduct) error {
	if m.WantErr != nil && strings.EqualFold("errUpdateProductByID", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m MockProductUseCase) DeleteProductByID(uuid string) error {
	if m.WantErr != nil && strings.EqualFold("errDeleteProductByID", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}
