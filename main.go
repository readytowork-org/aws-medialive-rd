package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/medialive"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-northeast-1"), config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("", "", "")))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Using the Config value, create the DynamoDB client
	svc := medialive.NewFromConfig(cfg)
	out, err := svc.StartChannel(context.TODO(), &medialive.StartChannelInput{
		ChannelId: aws.String("5867858"),
	})
	if err != nil {
		log.Fatalf("unable to load channel start", err)
	}

	fmt.Println(out)

}
