package reponosql

import (
	"fase-4-hf-product/internal/core/db"
	"fase-4-hf-product/internal/core/domain/entity/dto"
	"fase-4-hf-product/internal/core/domain/repository"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var _ repository.ProductRepository = (*productDB)(nil)

type productDB struct {
	Database  db.NoSQLDatabase
	tableName string
}

func NewProductRepository(database db.NoSQLDatabase, tableName string) *productDB {
	return &productDB{Database: database, tableName: tableName}
}

func (p *productDB) GetProductByUUID(uuid string) (*dto.ProductDB, error) {
	filter := "uuid = :value"
	attrSearch := map[string]*dynamodb.AttributeValue{
		":value": {
			S: aws.String(uuid),
		},
	}

	input := &dynamodb.ScanInput{
		TableName:                 aws.String(p.tableName),
		FilterExpression:          aws.String(filter),
		ExpressionAttributeValues: attrSearch,
	}

	result, err := p.Database.Scan(input)
	if err != nil {
		return nil, err
	}

	var productList = make([]dto.ProductDB, 0)
	for _, item := range result.Items {
		var pDb dto.ProductDB
		if err := dynamodbattribute.UnmarshalMap(item, &pDb); err != nil {
			return nil, err
		}
		productList = append(productList, pDb)
	}

	if len(productList) > 0 {
		return &productList[0], nil
	}

	return nil, nil
}

func (p *productDB) SaveProduct(product dto.ProductDB) (*dto.ProductDB, error) {

	priceString := strconv.FormatFloat(product.Price, 'f', -1, 64)
	putItem := map[string]*dynamodb.AttributeValue{
		"uuid": {
			S: aws.String(product.UUID),
		},
		"name": {
			S: aws.String(product.Name),
		},
		"category": {
			S: aws.String(product.Category),
		},
		"image": {
			S: aws.String(product.Image),
		},
		"description": {
			S: aws.String(product.Description),
		},
		"price": {
			N: aws.String(priceString),
		},
		"created_at": {
			S: aws.String(product.CreatedAt),
		},
		"deactivated_at": {
			S: aws.String(product.DeactivatedAt),
		},
	}

	inputPutItem := &dynamodb.PutItemInput{
		Item:      putItem,
		TableName: aws.String(p.tableName),
	}

	putOut, err := p.Database.PutItem(inputPutItem)

	if err != nil {
		return nil, err
	}

	var out *dto.ProductDB

	if err := dynamodbattribute.UnmarshalMap(putOut.Attributes, &out); err != nil {
		return nil, err
	}

	return out, nil
}

func (p *productDB) UpdateProductByUUID(uuid string, product dto.ProductDB) (*dto.ProductDB, error) {

	priceString := strconv.FormatFloat(product.Price, 'f', 2, 64)

	action := "SET name = :newName, category = :newCategory, image = :newImage, description = :newDescription, price = :newPrice, created_at = :newCreated_at, deactivated_at = :newDeactivated_at"
	key := map[string]*dynamodb.AttributeValue{
		"uuid": {
			S: aws.String(uuid),
		},
	}

	updateItem := map[string]*dynamodb.AttributeValue{
		":newName": {
			S: aws.String(product.Name),
		},
		":newCategory": {
			S: aws.String(product.Category),
		},
		":newImage": {
			S: aws.String(product.Image),
		},
		":newDescription": {
			S: aws.String(product.Description),
		},
		":newPrice": {
			N: aws.String(priceString),
		},
		":newCreated_at": {
			S: aws.String(product.CreatedAt),
		},
		":newDeactivated_at": {
			S: aws.String(product.DeactivatedAt),
		},
	}

	inputUpdateItem := &dynamodb.UpdateItemInput{
		TableName:                 aws.String(p.tableName),
		ExpressionAttributeValues: updateItem,
		Key:                       key,
		UpdateExpression:          aws.String(action),
	}

	pOut, err := p.Database.UpdateItem(inputUpdateItem)

	if err != nil {
		return nil, err
	}

	var out *dto.ProductDB

	if err := dynamodbattribute.UnmarshalMap(pOut.Attributes, &out); err != nil {
		return nil, err
	}

	return out, nil
}

func (p *productDB) GetProductByCategory(category string) ([]dto.ProductDB, error) {
	filter := "category = :value"
	attrSearch := map[string]*dynamodb.AttributeValue{
		":value": {
			S: aws.String(category),
		},
	}

	input := &dynamodb.ScanInput{
		TableName:                 aws.String(p.tableName),
		FilterExpression:          aws.String(filter),
		ExpressionAttributeValues: attrSearch,
	}

	result, err := p.Database.Scan(input)
	if err != nil {
		return nil, err
	}

	var productList = make([]dto.ProductDB, 0)
	for _, item := range result.Items {
		var pDb dto.ProductDB
		if err := dynamodbattribute.UnmarshalMap(item, &pDb); err != nil {
			return nil, err
		}
		productList = append(productList, pDb)
	}

	if len(productList) > 0 {
		return productList, nil
	}

	return nil, nil
}

func (p *productDB) DeleteProductByUUID(uuid string) error {
	key := map[string]*dynamodb.AttributeValue{
		"uuid": {S: aws.String("uuid")},
	}

	inputUpdateItem := &dynamodb.DeleteItemInput{
		TableName: aws.String(p.tableName),
		Key:       key,
	}

	if _, err := p.Database.DeleteItem(inputUpdateItem); err != nil {
		return err
	}

	return nil
}
