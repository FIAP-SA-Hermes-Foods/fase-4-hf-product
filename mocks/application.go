package mocks

import (
	"fase-4-hf-product/internal/core/domain/entity/dto"
	"strings"
)

type MockApplication struct {
	WantOut     *dto.OutputProduct
	WantOutList []dto.OutputProduct
	WantErr     error
	WantOutNull string
}

func (m MockApplication) GetProductByID(uuid string) (*dto.OutputProduct, error) {
	if m.WantErr != nil && strings.EqualFold("errGetProductByCPF", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilGetProductByCPF") {
		return nil, nil
	}
	return m.WantOut, nil
}

func (m MockApplication) SaveProduct(reqProduct dto.RequestProduct) (*dto.OutputProduct, error) {
	if m.WantErr != nil && strings.EqualFold("errSaveProduct", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilSaveProduct") {
		return nil, nil
	}
	return m.WantOut, nil
}

func (m MockApplication) GetProductByCategory(category string) ([]dto.OutputProduct, error) {
	if m.WantErr != nil && strings.EqualFold("errGetProductByCategory", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilGetProductByCategory") {
		return nil, nil
	}
	return m.WantOutList, nil
}

func (m MockApplication) UpdateProductByID(uuid string, product dto.RequestProduct) (*dto.OutputProduct, error) {
	if m.WantErr != nil && strings.EqualFold("errUpdateProductByID", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilUpdateProductByID") {
		return nil, nil
	}
	return m.WantOut, nil
}

func (m MockApplication) DeleteProductByID(uuid string) error {
	if m.WantErr != nil && strings.EqualFold("errDeleteProductByID", m.WantErr.Error()) {
		return m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilDeleteProductByID") {
		return nil
	}
	return nil
}

// Repository Callers
type MockApplicationRepostoryCallers struct {
	WantOut     *dto.ProductDB
	WantOutList []dto.ProductDB
	WantErr     error
}

func (m MockApplicationRepostoryCallers) GetProductByIDRepository(uuid string) (*dto.ProductDB, error) {
	if m.WantErr != nil && strings.EqualFold("errGetProductByIDRepository", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	return m.WantOut, nil
}

func (m MockApplicationRepostoryCallers) GetProductByCategoryRepository(category string) ([]dto.ProductDB, error) {
	if m.WantErr != nil && strings.EqualFold("errGetProductByCategoryRepository", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	return m.WantOutList, nil
}

func (m MockApplicationRepostoryCallers) UpdateProductByIDRepository(uuid string, product dto.ProductDB) (*dto.ProductDB, error) {
	if m.WantErr != nil && strings.EqualFold("errUpdateProductByIDRepository", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	return m.WantOut, nil
}

func (m MockApplicationRepostoryCallers) DeleteProductByIDRepository(uuid string) error {
	if m.WantErr != nil && strings.EqualFold("errDeleteProductByIDRepository", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m MockApplicationRepostoryCallers) SaveProductRepository(product dto.ProductDB) (*dto.ProductDB, error) {
	if m.WantErr != nil && strings.EqualFold("errSaveProductRepository", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	return m.WantOut, nil
}

// UseCase callers
type MockApplicationUseCaseCallers struct {
	WantErr error
}

func (m MockApplicationUseCaseCallers) GetProductByIDUseCase(cpf string) error {
	if m.WantErr != nil && strings.EqualFold("errGetProductByIDUseCase", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m MockApplicationUseCaseCallers) SaveProductUseCase(product dto.RequestProduct) error {
	if m.WantErr != nil && strings.EqualFold("errSaveProductUseCase", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m MockApplicationUseCaseCallers) GetProductByCategoryUseCase(category string) error {
	if m.WantErr != nil && strings.EqualFold("errGetProductByCategoryUseCase", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m MockApplicationUseCaseCallers) UpdateProductByIDUseCase(uuid string, product dto.RequestProduct) error {
	if m.WantErr != nil && strings.EqualFold("errUpdateProductByIDUseCase", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m MockApplicationUseCaseCallers) DeleteProductByIDUseCase(uuid string) error {
	if m.WantErr != nil && strings.EqualFold("errDeleteProductByIDUseCase", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}
