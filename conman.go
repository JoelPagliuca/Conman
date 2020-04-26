package conman

import "log"

// Iface ...
type Iface interface {
	HydrateConfig(interface{}) error
}

// Conman ...
type Conman struct {
	cfg Cfg
}

// New ...
func New(c Cfg) Iface {
	return Conman{cfg: c}
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
	// TODO: implement
	cm.whinge("NOT IMPLEMENTED")
	return nil
}
