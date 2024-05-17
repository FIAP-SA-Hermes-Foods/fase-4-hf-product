package dynamo

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type dynamoDB struct {
	session *session.Session
	client  *dynamodb.DynamoDB
}

func NewDynamoDB(session *session.Session) *dynamoDB {
	return &dynamoDB{session: session}
}

func (d *dynamoDB) clientDynamo() {
	d.client = dynamodb.New(d.session)
}

func (d *dynamoDB) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	if d.client == nil {
		d.clientDynamo()
	}
	return d.client.Scan(input)
}

func (d *dynamoDB) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	if d.client == nil {
		d.clientDynamo()
	}
	return d.client.PutItem(input)
}

func (d *dynamoDB) UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	if d.client == nil {
		d.clientDynamo()
	}
	return d.client.UpdateItem(input)
}

func (d *dynamoDB) DeleteItem(input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	if d.client == nil {
		d.clientDynamo()
	}
	return d.client.DeleteItem(input)
}
