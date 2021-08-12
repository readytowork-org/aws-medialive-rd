package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/medialive"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-northeast-1"), config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(os.Getenv("AWS_KEY"), os.Getenv("AWS_SECRET"), "")))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Using the Config value, create the DynamoDB client
	svc := medialive.NewFromConfig(cfg)

	// start channel
	/*out, err := svc.StartChannel(context.TODO(), &medialive.StartChannelInput{
		ChannelId: aws.String("5867858"),
	})
	if err != nil {
		log.Fatalf("unable to load channel start: ", err)
	}
	*/

	// stop channel
	out, err := svc.StopChannel(context.TODO(), &medialive.StopChannelInput{
		ChannelId: aws.String("5867858"),
	})
	if err != nil {
		log.Fatalln("unable to load channel stop: ", err)
	}

	fmt.Println(out)

}
