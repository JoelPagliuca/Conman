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
type Strategy func(*Conman, string) (*string, error)

// DefaultStrategy sets the default defined by "in"
func DefaultStrategy(cm *Conman, in string) (*string, error) {
	return &in, nil
}

// EnvironmentStrategy gets the value of the environment variable "in"
func EnvironmentStrategy(cm *Conman, in string) (*string, error) {
	val, ok := os.LookupEnv(in)
	if ok {
		return &val, nil
	}
	return nil, fmt.Errorf("no env var \"%s\" found", in)
}

// SSMStrategy gets the value stored in the SSM Parameter "in"
func SSMStrategy(cm *Conman, in string) (*string, error) {
	if cm.awsConfig == nil {
		a, err := external.LoadDefaultAWSConfig()
		cm.awsConfig = &a
		if err != nil {
			return nil, err
		}
		cm.inform("Getting default AWS config")
	}
	ssmClient := ssm.New(*cm.awsConfig)
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
