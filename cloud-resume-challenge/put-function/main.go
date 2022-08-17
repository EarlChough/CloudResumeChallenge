package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)


func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO(), 
	config.WithRegion("us-east-1"),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	svc := dynamodb.NewFromConfig(cfg)

    _, err = svc.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
        TableName: aws.String("cloud-resume-challenge8417"),
        Key: map[string]types.AttributeValue{
            "ID": &types.AttributeValueMemberS{Value: "visitors"},
        },
        UpdateExpression: aws.String("set visitors = visitors + :value"),
        ExpressionAttributeValues: map[string]types.AttributeValue{
            ":value": &types.AttributeValueMemberN{Value: "1"},
        },
    })

    if err != nil {
        log.Fatalf("Got error calling UpdateItem: %s", &err)
    }

	

//	return attributeMap, err

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "*",
			"Access-Control-Allow-Headers": "*",
		},
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
