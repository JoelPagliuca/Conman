package conman

import (
	"fmt"
	"log"
	"reflect"
)

// Iface ...
type Iface interface {
	Hydrate(interface{}) error
}

// Conman ...
type Conman struct {
	cfg        Cfg
	strategies map[string]Strategy
}

// New ...
func New(c Cfg) (Iface, error) {
	cm := Conman{cfg: c}
	cm.strategies = make(map[string]Strategy)
	cm.strategies[SourceDefault] = DefaultStrategy
	cm.strategies[SourceEnvironment] = EnvironmentStrategy
	// make sure the given config was safe
	for _, ord := range c.SourceOrder {
		if _, ok := cm.strategies[ord]; !ok {
			return nil, fmt.Errorf("Value " + ord + " from SourceOrder doesn't have corresponding strategy")
		}
	}
	if len(c.SourceOrder) == 0 {
		cm.inform("Using the default ordering")
		cm.cfg.SourceOrder = DefaultCfg.SourceOrder
	}
	return &cm, nil
}

func (cm Conman) inform(s string) {
	if cm.cfg.LogInfo {
		log.Println(s)
	}
}

// Hydrate - populate a config object with ssm params defined by tags.
// Looks for ssmConfig path from env var
func (cm Conman) Hydrate(cfg interface{}) error {
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
		for _, src := range cm.cfg.SourceOrder {
			tag := fld.Tag.Get(src)
			if tag == "" {
				continue
			} else {
				sub, err := cm.strategies[src](tag)
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
