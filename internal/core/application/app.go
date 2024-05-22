package application

import (
	"errors"
	l "fase-4-hf-product/external/logger"
	ps "fase-4-hf-product/external/strings"
	"fase-4-hf-product/internal/core/domain/entity/dto"
	"fase-4-hf-product/internal/core/domain/repository"
	"fase-4-hf-product/internal/core/domain/useCase"
	"time"

	"github.com/google/uuid"
)

type Application interface {
	GetProductByUUID(uuid string) (*dto.OutputProduct, error)
	SaveProduct(reqProduct dto.RequestProduct) (*dto.OutputProduct, error)
	GetProductByCategory(category string) ([]dto.OutputProduct, error)
	UpdateProductByUUID(uuid string, product dto.RequestProduct) (*dto.OutputProduct, error)
	DeleteProductByUUID(uuid string) error
}

type application struct {
	productRepo repository.ProductRepository
	productUC   useCase.ProductUseCase
}

func NewApplication(productRepo repository.ProductRepository, productUC useCase.ProductUseCase) Application {
	return application{productRepo: productRepo, productUC: productUC}
}

func (app application) GetProductByUUID(uuid string) (*dto.OutputProduct, error) {
	l.Infof("GetProductByUUIDApp: ", " | ", uuid)
	if err := app.GetProductByUUIDUseCase(uuid); err != nil {
		l.Errorf("GetProductByUUIDApp error: ", " | ", err)
		return nil, err
	}

	cOutDb, err := app.GetProductByUUIDRepository(uuid)

	if err != nil {
		l.Errorf("GetProductByUUIDApp error: ", " | ", err)
		return nil, err
	}

	if cOutDb == nil {
		l.Infof("GetProductByUUIDApp output: ", " | ", cOutDb)
		return nil, nil
	}

	out := &dto.OutputProduct{
		UUID:          cOutDb.UUID,
		Name:          cOutDb.Name,
		Category:      cOutDb.Category,
		Image:         cOutDb.Image,
		Description:   cOutDb.Description,
		Price:         cOutDb.Price,
		CreatedAt:     cOutDb.CreatedAt,
		DeactivatedAt: cOutDb.DeactivatedAt,
	}

	l.Infof("GetProductByUUIDApp output: ", " | ", ps.MarshalString(out))
	return out, err
}

func (app application) SaveProduct(product dto.RequestProduct) (*dto.OutputProduct, error) {
	l.Infof("SaveProductApp: ", " | ", ps.MarshalString(product))

	if err := app.SaveProductUseCase(product); err != nil {
		l.Errorf("SaveProductApp error: ", " | ", err)
		return nil, err
	}

	createdAtFmt := time.Now().Format(`02-01-2006 15:04:05`)

	productDB := dto.ProductDB{
		UUID:          uuid.NewString(),
		Name:          product.Name,
		Category:      product.Category,
		Image:         product.Image,
		Description:   product.Description,
		Price:         product.Price,
		CreatedAt:     createdAtFmt,
		DeactivatedAt: product.DeactivatedAt,
	}

	cOutDb, err := app.SaveProductRepository(productDB)

	if err != nil {
		l.Errorf("SaveProductApp error: ", " | ", err)
		return nil, err
	}

	if cOutDb == nil {
		l.Infof("SaveProductApp output: ", " | ", nil)
		return nil, nil
	}

	out := &dto.OutputProduct{
		UUID:          cOutDb.UUID,
		Name:          cOutDb.Name,
		Category:      cOutDb.Category,
		Image:         cOutDb.Image,
		Description:   cOutDb.Description,
		Price:         cOutDb.Price,
		CreatedAt:     cOutDb.CreatedAt,
		DeactivatedAt: cOutDb.DeactivatedAt,
	}

	l.Infof("SaveProductApp output: ", " | ", ps.MarshalString(out))

	return out, nil
}

func (app application) GetProductByCategory(category string) ([]dto.OutputProduct, error) {
	l.Infof("GetProductByCategoryApp: ", " | ", category)

	if err := app.GetProductByCategoryUseCase(category); err != nil {
		l.Errorf("GetProductByCategoryApp error: ", " | ", err)
		return nil, err
	}

	productList := make([]dto.OutputProduct, 0)
	products, err := app.GetProductByCategoryRepository(category)

	if err != nil {
		l.Errorf("GetProductByCategoryApp error: ", " | ", err)
		return nil, err
	}

	if products == nil {
		l.Infof("GetProductByCategoryApp output: ", " | ", nil)
		return nil, nil
	}

	for i := range products {
		product := dto.OutputProduct{
			UUID:          products[i].UUID,
			Name:          products[i].Name,
			Category:      products[i].Category,
			Image:         products[i].Image,
			Description:   products[i].Description,
			Price:         products[i].Price,
			CreatedAt:     products[i].CreatedAt,
			DeactivatedAt: products[i].CreatedAt,
		}
		productList = append(productList, product)
	}

	l.Infof("GetProductByCategoryApp output: ", " | ", productList)
	return productList, nil
}

func (app application) UpdateProductByUUID(id string, newProduct dto.RequestProduct) (*dto.OutputProduct, error) {
	l.Infof("UpdateProductByUUIDApp: ", " | ", id, " | ", ps.MarshalString(newProduct))

	if err := app.UpdateProductByUUIDUseCase(id, newProduct); err != nil {
		l.Errorf("UpdateProductByUUIDApp error: ", " | ", err)
		return nil, err
	}

	product, err := app.GetProductByUUID(id)

	if err != nil {
		l.Errorf("UpdateProductByUUIDApp error: ", " | ", err)
		return nil, err
	}

	if product == nil {
		l.Errorf("UpdateProductByUUIDApp error: ", " | ", "product with this uuid was not found")
		return nil, err
	}

	var (
		name          = product.Name
		category      = product.Category
		image         = product.Image
		description   = product.Description
		price         = product.Price
		createdAt     = product.CreatedAt
		deactivatedAt = product.DeactivatedAt
	)

	if len(newProduct.Name) > 0 {
		name = newProduct.Name
	}

	if len(newProduct.Category) > 0 {
		category = newProduct.Category
	}

	if len(newProduct.Image) > 0 {
		image = newProduct.Image
	}

	if len(newProduct.Description) > 0 {
		description = newProduct.Description
	}

	if newProduct.Price != 0 {
		price = newProduct.Price
	}

	if len(newProduct.CreatedAt) > 0 {
		createdAt = newProduct.CreatedAt
	}

	if len(newProduct.DeactivatedAt) > 0 {
		deactivatedAt = newProduct.DeactivatedAt
	}

	productDB := dto.ProductDB{
		Name:          name,
		Category:      category,
		Image:         image,
		Description:   description,
		Price:         price,
		CreatedAt:     createdAt,
		DeactivatedAt: deactivatedAt,
	}

	cOutDb, err := app.UpdateProductByUUIDRepository(id, productDB)

	if err != nil {
		l.Errorf("UpdateProductByUUIDApp error: ", " | ", err)
		return nil, err
	}

	if cOutDb == nil {
		l.Infof("UpdateProductByUUIDApp output: ", " | ", nil)
		return nil, nil
	}

	out := &dto.OutputProduct{
		UUID:          cOutDb.UUID,
		Name:          cOutDb.Name,
		Category:      cOutDb.Category,
		Image:         cOutDb.Image,
		Description:   cOutDb.Description,
		Price:         cOutDb.Price,
		CreatedAt:     cOutDb.CreatedAt,
		DeactivatedAt: cOutDb.DeactivatedAt,
	}

	l.Infof("UpdateProductByUUIDApp output: ", " | ", ps.MarshalString(out))
	return out, nil
}

func (app application) DeleteProductByUUID(id string) error {
	l.Infof("DeleteProductByUUIDApp: ", " | ", id)

	pByUUID, err := app.GetProductByUUID(id)

	if err != nil {
		l.Errorf("DeleteProductByUUIDApp error: ", " | ", err)
		return err
	}

	if pByUUID == nil {
		productNullErr := errors.New("was not found any product with this id")
		l.Infof("DeleteProductByUUIDApp output: ", " | ", productNullErr)
		return productNullErr
	}

	if err := app.DeleteProductByUUIDUseCase(id); err != nil {
		l.Errorf("DeleteProductByUUIDApp error: ", " | ", err)
		return err
	}

	l.Infof("DeleteProductByUUIDApp output: ", " | ", nil)
	return app.DeleteProductByUUIDRepository(id)
}
