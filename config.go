package conman

//
const (
	SourceEnvironment = "cmenv"
	SourceDefault     = "cmdefault"
)

// Cfg configuration for Conman
type Cfg struct {
	// Log kinda interesting info
	LogInfo bool
	// Order to check for config values in
	SourceOrder []string
}

// DefaultCfg sensible defaults
var DefaultCfg = Cfg{
	LogInfo: false,
	SourceOrder: []string{
		SourceEnvironment,
		SourceDefault,
	},
}
