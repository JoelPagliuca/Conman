package conman

import (
	"os"
)

// Strategy iface for a config getting func
type Strategy func(string) *string

// DefaultStrategy for SourceDefault
func DefaultStrategy(in string) *string {
	return &in
}

// EnvironmentStrategy for SourceEnvironment
func EnvironmentStrategy(in string) *string {
	val, ok := os.LookupEnv(in)
	if ok {
		return &val
	}
	return nil
}
