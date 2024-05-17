package mocks

import (
	"strings"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Mock DB NoSQL
type MockDbNoSQL struct {
	WantResultScan       *dynamodb.ScanOutput
	WantResultPutItem    *dynamodb.PutItemOutput
	WantResultUpdateItem *dynamodb.UpdateItemOutput
	WantResultDeleteItem *dynamodb.DeleteItemOutput
	WantErr              error
}

func (m MockDbNoSQL) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	if m.WantErr != nil && strings.EqualFold("errScan", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	return m.WantResultScan, nil
}

func (m MockDbNoSQL) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	if m.WantErr != nil && strings.EqualFold("errPutItem", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	return m.WantResultPutItem, nil
}

func (m MockDbNoSQL) UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	if m.WantErr != nil && strings.EqualFold("errUpdateItem", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	return m.WantResultUpdateItem, nil
}

func (m MockDbNoSQL) DeleteItem(input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	if m.WantErr != nil && strings.EqualFold("errDeleteItem", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	return m.WantResultDeleteItem, nil
}
