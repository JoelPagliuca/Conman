// Package conman aims to make your config setup less verbose.
//
// This package uses `reflect` so watch out
package conman

import (
	"fmt"
	"log"
	"reflect"
)

// Conman ...
type Conman struct {
	logInfo    bool
	order      []string
	strategies map[string]Strategy
	envPrefix  string
}

// New sets the defaults then applies all the options
func New(opts ...Option) (*Conman, error) {
	cm := Conman{}
	cm.strategies = make(map[string]Strategy)
	AddStrategy(TagDefault, DefaultStrategy)(&cm)
	AddStrategy(TagEnvironment, EnvironmentStrategy)(&cm)
	for _, o := range opts {
		err := o(&cm)
		if err != nil {
			return nil, err
		}
	}
	return &cm, nil
}

func (cm Conman) inform(s string) {
	if cm.logInfo {
		log.Println(s)
	}
}

// Hydrate - populate a config object with params defined by tags.
func (cm *Conman) Hydrate(cfg interface{}) error {
	defer func() {
		if r := recover(); r != nil {
			cm.inform(fmt.Sprintf("Had a panic: %s", r))
		}
	}()
	val := reflect.ValueOf(cfg)
	// the struct itself
	str := val.Elem()
	// the structure of cfg
	typ := str.Type()
	for i := 0; i < str.NumField(); i++ {
		// fld should be the field on the stuct we were given
		fld := typ.Field(i)
		val := str.Field(i)
		if !val.CanSet() {
			continue
		}
		var finalError error
		for _, src := range cm.order {
			tag := fld.Tag.Get(src)
			if tag == "" {
				continue
			} else {
				sub, err := cm.strategies[src](cm, tag)
				if sub != nil {
					cm.inform("Setting " + fld.Name + " using " + src)
					val.SetString(*sub)
					finalError = nil
					break
				}
				if err != nil {
					finalError = fmt.Errorf("%s with tag %s: %w", fld.Name, src, err)
				}
			}
		}
		if finalError != nil {
			return finalError
		}
	}
	return nil
}
