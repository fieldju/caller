package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"os"
)

func main() {
	svc := sts.New(session.New())
	input := &sts.GetCallerIdentityInput{}
	result, err := svc.GetCallerIdentity(input)
	if err != nil {
		fmt.Errorf("failed to get an identity response from AWS")
		fmt.Errorf(err.Error())
		os.Exit(1)
	}
	arn := *result.Arn
	fmt.Println(arn)
}
