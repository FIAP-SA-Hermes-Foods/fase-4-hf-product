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

func (m MockProductRepository) GetProductByID(uuid string) (*dto.ProductDB, error) {
	if m.WantErr != nil && strings.EqualFold("errGetProductByID", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilGetProductByID") {
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

func (m MockProductRepository) UpdateProductByID(uuid string, product dto.ProductDB) (*dto.ProductDB, error) {
	if m.WantErr != nil && strings.EqualFold("errUpdateProductByID", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilUpdateProductByID") {
		return nil, nil
	}
	return m.WantOut, nil
}

func (m MockProductRepository) DeleteProductByID(uuid string) error {
	if m.WantErr != nil && strings.EqualFold("errDeleteProductByID", m.WantErr.Error()) {
		return m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilDeleteProductByID") {
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
