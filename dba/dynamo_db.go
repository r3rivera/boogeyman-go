package dba

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const USER_CREDENTIAL_TBL = "b_user_credential"

type DDBUserCredential struct {
	email string
	hash  string
}

func NewDDBUserCredential(email, hash string) DDBUserCredential {
	return DDBUserCredential{
		email, hash,
	}
}

func (c *DDBUserCredential) WriteDB() error {
	ddbClient := getDynamodbClient()
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
