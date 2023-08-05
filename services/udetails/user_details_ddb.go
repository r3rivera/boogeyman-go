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
	Firstname string    `dynamodbav:"firstname" json:"firstname"`
	Lastname  string    `dynamodbav:"lastname" json:"lastname"`
	Dob       time.Time `dynamodbav:"dob" json:"dob"`
}

type UserAddress struct {
	Street1 string `dynamodbav:"street1" json:"street1"`
	Street2 string `dynamodbav:"street2" json:"street2"`
	City    string `dynamodbav:"city" json:"city"`
	State   string `dynamodbav:"state" json:"state"`
	Zip     string `dynamodbav:"zip" json:"zip"`
}

type DDBUserDetail struct {
	Email   string      `dynamodbav:"email" json:"email"`
	Info    UserDetail  `dynamodbav:"details" json:"details"`
	Address UserAddress `dynamodbav:"address" json:"address"`
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

	details, err := attributevalue.MarshalMap(c.Info)
	if err != nil {
		return err
	}

	address, err := attributevalue.MarshalMap(c.Address)
	if err != nil {
		return err
	}
	//Create the record in the credential table
	_, err = ddbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(USER_DETAIL_TBL),
		Item: map[string]types.AttributeValue{
			"email":   &types.AttributeValueMemberS{Value: c.Email},
			"details": &types.AttributeValueMemberM{Value: details},
			"address": &types.AttributeValueMemberM{Value: address},
		},
	})

	if err != nil {
		return err
	}
	return nil
}

func (c *DDBUserDetail) ReadDB() (DDBUserDetail, error) {
	ddbClient := dba.GetDynamodbClient()

	result, err := ddbClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(USER_DETAIL_TBL),
		Key: map[string]types.AttributeValue{
			"email": &types.AttributeValueMemberS{Value: c.Email},
		},
	})
	if err != nil {
		return DDBUserDetail{}, err
	}
	output := &DDBUserDetail{}
	attributevalue.UnmarshalMap(result.Item, output)
	return *output, nil
}

func (c *DDBUserDetail) DeleteDB() error {
	ddbClient := dba.GetDynamodbClient()
	_, err := ddbClient.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(USER_DETAIL_TBL),
		Key: map[string]types.AttributeValue{
			"email": &types.AttributeValueMemberS{Value: c.Email},
		},
	})

	if err != nil {
		return err
	}
	return nil
}
