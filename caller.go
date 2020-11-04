package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"os"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		CredentialsChainVerboseErrors: aws.Bool(true),
	})

	if err != nil {
		fmt.Println("failed to find valid aws credentials")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	svc := sts.New(sess)
	input := &sts.GetCallerIdentityInput{}
	result, err := svc.GetCallerIdentity(input)
	if err != nil {
		fmt.Println("failed to get an identity response from AWS")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	json, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal aws sts response into json")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(string(json))
}
