package conman

import (
	"fmt"
	"os"
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
