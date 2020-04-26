package conman

// Cfg configuration for Conman
type Cfg struct {
	logInfo          bool
	suppressWarnings bool
}

// DefaultCfg sensible defaults
var DefaultCfg = Cfg{
	logInfo:          false,
	suppressWarnings: false,
}
