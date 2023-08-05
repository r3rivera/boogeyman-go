package ucreds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/r3rivera/boogeyman/dba"
)

const USER_CREDENTIAL_TBL = "b_user_credential"

type DDBUserCredential struct {
	email string
	hash  string
}

func NewDDBUserCredential(email, hash string) *DDBUserCredential {
	return &DDBUserCredential{
		email, hash,
	}
}

func (c *DDBUserCredential) WriteDB() error {
	ddbClient := dba.GetDynamodbClient()

	//Create the record in the credential table
	_, err := ddbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(USER_CREDENTIAL_TBL),
		Item: map[string]types.AttributeValue{
			"email": &types.AttributeValueMemberS{Value: c.email},
			"hash":  &types.AttributeValueMemberS{Value: c.hash},
		},
	})

	if err != nil {
		return err
	}
	return nil
}

func (c *DDBUserCredential) ReadDB() (string, error) {
	ddbClient := dba.GetDynamodbClient()

	result, err := ddbClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(USER_CREDENTIAL_TBL),
		Key: map[string]types.AttributeValue{
			"email": &types.AttributeValueMemberS{Value: c.email},
		},
	})
	if err != nil {
		return "", err
	}

	if len(result.Item) == 0 {
		return "", nil
	}

	o := &struct {
		Email string `dynamodbav:"email" json:"email"`
		Hash  string `dynamodbav:"hash" json:"hash"`
	}{}

	attributevalue.UnmarshalMap(result.Item, o)
	return o.Hash, nil
}

func (c *DDBUserCredential) DeleteDB() error {
	ddbClient := dba.GetDynamodbClient()
	_, err := ddbClient.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(USER_CREDENTIAL_TBL),
		Key: map[string]types.AttributeValue{
			"email": &types.AttributeValueMemberS{Value: c.email},
		},
	})

	if err != nil {
		return err
	}
	return nil
}
