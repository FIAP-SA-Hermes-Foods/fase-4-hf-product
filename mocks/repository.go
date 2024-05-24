package mocks

import (
	"fase-4-hf-product/internal/core/domain/entity/dto"
	"strings"
)

type MockProductRepository struct {
	WantOut     *dto.ProductDB
	WantOutList []dto.ProductDB
	WantOutNull string
	WantErr     error
}

func (m MockProductRepository) GetProductByUUID(uuid string) (*dto.ProductDB, error) {
	if m.WantErr != nil && strings.EqualFold("errGetProductByUUID", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilGetProductByUUID") {
		return nil, nil
	}
	return m.WantOut, nil
}

func (m MockProductRepository) GetProductByCategory(category string) ([]dto.ProductDB, error) {
	if m.WantErr != nil && strings.EqualFold("errGetProductByCategory", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilGetProductByCategory") {
		return nil, nil
	}
	return m.WantOutList, nil
}

func (m MockProductRepository) UpdateProductByUUID(uuid string, product dto.ProductDB) (*dto.ProductDB, error) {
	if m.WantErr != nil && strings.EqualFold("errUpdateProductByUUID", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilUpdateProductByUUID") {
		return nil, nil
	}
	return m.WantOut, nil
}

func (m MockProductRepository) DeleteProductByUUID(uuid string) error {
	if m.WantErr != nil && strings.EqualFold("errDeleteProductByUUID", m.WantErr.Error()) {
		return m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilDeleteProductByUUID") {
		return nil
	}
	return nil
}

func (m MockProductRepository) SaveProduct(product dto.ProductDB) (*dto.ProductDB, error) {
	if m.WantErr != nil && strings.EqualFold("errSaveProduct", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilSaveProduct") {
		return nil, nil
	}
	return m.WantOut, nil
}
