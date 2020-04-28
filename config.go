package conman

//
const (
	SourceEnvironment = "cmenv"
	SourceDefault     = "cmdefault"
)

// Cfg configuration for Conman
type Cfg struct {
	// loginfo - Log kinda interesting info
	logInfo bool
	// Stop the logs that are helpful
	suppressWarnings bool
	// Order to check for config values in
	sourceOrder []string
}

// DefaultCfg sensible defaults
var DefaultCfg = Cfg{
	logInfo:          false,
	suppressWarnings: false,
	sourceOrder: []string{
		SourceEnvironment,
		SourceDefault,
	},
}
