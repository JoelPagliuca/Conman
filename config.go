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
	// Stop the logs that are helpful
	SuppressWarnings bool
	// Order to check for config values in
	SourceOrder []string
}

// DefaultCfg sensible defaults
var DefaultCfg = Cfg{
	LogInfo:          false,
	SuppressWarnings: false,
	SourceOrder: []string{
		SourceEnvironment,
		SourceDefault,
	},
}
