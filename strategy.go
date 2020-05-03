package conman

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

// Strategy iface for a config getting func
type Strategy func(string) (*string, error)

// DefaultStrategy sets the default defined by "in"
func DefaultStrategy(in string) (*string, error) {
	return &in, nil
}

// EnvironmentStrategy gets the value of the environment variable "in"
func EnvironmentStrategy(in string) (*string, error) {
	val, ok := os.LookupEnv(in)
	if ok {
		return &val, nil
	}
	return nil, fmt.Errorf("no env var \"%s\" found", in)
}

// SSMStrategy gets the value stored in the SSM Parameter "in"
func SSMStrategy(in string) (*string, error) {
	awsCfg, _ := external.LoadDefaultAWSConfig()
	ssmClient := ssm.New(awsCfg)
	input := ssm.GetParameterInput{
		Name:           aws.String(in),
		WithDecryption: aws.Bool(true),
	}
	req := ssmClient.GetParameterRequest(&input)
	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}
	return resp.Parameter.Value, nil
}
