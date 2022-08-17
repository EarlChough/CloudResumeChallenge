package main

import (
	"fmt"
	"log"
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	
	cfg, err := config.LoadDefaultConfig(context.TODO(), 
	config.WithRegion("us-east-1"),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	svc := dynamodb.NewFromConfig(cfg)

	
    out, err := svc.GetItem(context.TODO(), &dynamodb.GetItemInput{
        TableName: aws.String("cloud-resume-challenge8417"),
        Key: map[string]types.AttributeValue{
            "ID": &types.AttributeValueMemberS{Value: "visitors"},
        },
    })
	if err != nil {
        panic(err)
    }

	type Count struct {
		ID string `json:"id"`
		Visitors string `json:"visitors"`
	}

	count := Count{}

	err = attributevalue.UnmarshalMap(out.Item, &count)

	if err != nil {
		log.Fatalf("Got error calling UpdateItem: %s", err)
	}


	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "*",
			"Access-Control-Allow-Headers": "*",
		},
		Body:       fmt.Sprintf("{ \"count\": \"%s\" }", count.Visitors),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
