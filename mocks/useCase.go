package mocks

import (
	"fase-4-hf-product/internal/core/domain/entity/dto"
	"strings"
)

type MockProductUseCase struct {
	WantOutNull string
	WantErr     error
}

func (m MockProductUseCase) GetProductByUUID(cpf string) error {
	if m.WantErr != nil && strings.EqualFold("errGetProductByUUID", m.WantErr.Error()) {
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

func (m MockProductUseCase) UpdateProductByUUID(uuid string, product dto.RequestProduct) error {
	if m.WantErr != nil && strings.EqualFold("errUpdateProductByUUID", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m MockProductUseCase) DeleteProductByUUID(uuid string) error {
	if m.WantErr != nil && strings.EqualFold("errDeleteProductByUUID", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}
