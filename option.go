package conman

import (
	"github.com/aws/aws-sdk-go-v2/aws"
)

// Tags for the Hydrater to look for
const (
	TagEnvironment = "cmenv"
	TagSSM         = "cmssm"
	TagDefault     = "cmdefault"
)

// Option applies some sort of option to a Conman
type Option func(*Conman) error

// AddStrategy add a strategy to be used by the Hydrater
// will add it as first in the ordering
func AddStrategy(tag string, str Strategy) Option {
	return func(cm *Conman) error {
		cm.strategies[tag] = str
		cm.order = append([]string{tag}, cm.order...)
		return nil
	}
}

// SetOrder choose an order to check for config in
func SetOrder(src ...string) Option {
	return func(cm *Conman) error {
		cm.order = src[:]
		return nil
	}
}

// EnableLogging - Log kinda interesting info
func EnableLogging() Option {
	return func(cm *Conman) error {
		cm.logInfo = true
		return nil
	}
}

// AddAWSConfig add your own AWS config to Conman if you don't want default
func AddAWSConfig(c *aws.Config) Option {
	return func(cm *Conman) error {
		cm.awsConfig = c
		return nil
	}
}
