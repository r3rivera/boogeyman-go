package domain

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/r3rivera/boogeyman/adapters"
)

type UserRegistration struct {
	FirstName string
	LastName  string
	Email     string
}

func (u *UserRegistration) Register() error {
	ddbClient := adapters.GetDynamodbClient()
	_, err := ddbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("b_user_registration"),
		Item: map[string]types.AttributeValue{
			"email":    &types.AttributeValueMemberS{Value: u.Email},
			"username": &types.AttributeValueMemberS{Value: u.Email},
			"details": &types.AttributeValueMemberM{
				Value: map[string]types.AttributeValue{
					"firstname": &types.AttributeValueMemberS{Value: u.FirstName},
					"lastname":  &types.AttributeValueMemberS{Value: u.LastName},
				},
			},
		},
	})

	if err != nil {
		return err
	}
	return nil
}
