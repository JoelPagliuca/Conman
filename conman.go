package conman

import (
	"fmt"
	"log"
	"reflect"
)

// Iface ...
type Iface interface {
	HydrateConfig(interface{}) error
}

// Conman ...
type Conman struct {
	cfg        Cfg
	strategies map[string]Strategy
}

// New ...
func New(c Cfg) Iface {
	cm := Conman{cfg: c}
	cm.strategies = make(map[string]Strategy)
	cm.strategies[SourceDefault] = DefaultStrategy
	cm.strategies[SourceEnvironment] = EnvironmentStrategy
	// make sure the given config was safe
	for _, ord := range c.sourceOrder {
		if _, ok := cm.strategies[ord]; !ok {
			cm.whinge("Value " + ord + " from sourceOrder doesn't have corresponding strategy")
		}
	}
	if len(c.sourceOrder) == 0 {
		cm.inform("Using the default ordering")
		cm.cfg.sourceOrder = DefaultCfg.sourceOrder
	}
	return cm
}

func (cm Conman) inform(s string) {
	if cm.cfg.logInfo {
		log.Println(s)
	}
}

func (cm Conman) whinge(s string) {
	if !cm.cfg.suppressWarnings {
		log.Println("\033[1;33m", s, "\033[0m")
	}
}

// HydrateConfig - populate a config object with ssm params defined by tags.
// Looks for ssmConfig path from env var
func (cm Conman) HydrateConfig(cfg interface{}) error {
	defer func() {
		if r := recover(); r != nil {
			cm.whinge(fmt.Sprintf("Had a panic: %s", r))
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
		for _, src := range cm.cfg.sourceOrder {
			tag := fld.Tag.Get(src)
			if tag == "" {
				continue
			} else {
				sub := cm.strategies[src](tag)
				if sub != nil {
					cm.inform("Setting " + fld.Name + " using " + src)
					val.SetString(*sub)
					break
				}
			}
		}
	}
	return nil
}
