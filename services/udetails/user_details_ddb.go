package udetails

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/r3rivera/boogeyman/dba"
)

const USER_DETAIL_TBL = "b_user_info"

type UserDetail struct {
	firstname string    `dynamodbav:"firstname" json:"firstname"`
	lastname  string    `dynamodbav:"lastname" json:"lastname"`
	dob       time.Time `dynamodbav:"dob" json:"dob"`
}

type UserAddress struct {
	street1 string `dynamodbav:"street1" json:"street1"`
	street2 string `dynamodbav:"street2" json:"street2"`
	city    string `dynamodbav:"city" json:"city"`
	state   string `dynamodbav:"state" json:"state"`
	zip     string `dynamodbav:"zip" json:"zip"`
}

type DDBUserDetail struct {
	email   string
	info    UserDetail
	address UserAddress
}

func NewUserDetail(email string, info UserDetail, address UserAddress) *DDBUserDetail {
	return &DDBUserDetail{
		email,
		info,
		address,
	}
}

func (c *DDBUserDetail) WriteDB() error {
	ddbClient := dba.GetDynamodbClient()

	details, err := attributevalue.MarshalMap(c.info)
	if err != nil {
		return err
	}

	address, err := attributevalue.MarshalMap(c.address)
	if err != nil {
		return err
	}
	//Create the record in the credential table
	_, err = ddbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(USER_DETAIL_TBL),
		Item: map[string]types.AttributeValue{
			"email":   &types.AttributeValueMemberS{Value: c.email},
			"details": &types.AttributeValueMemberM{Value: details},
			"address": &types.AttributeValueMemberM{Value: address},
		},
	})

	if err != nil {
		return err
	}
	return nil
}
