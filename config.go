package conman

//
const (
	SourceEnvironment = "cmenv"
	SourceDefault     = "cmdefault"
)

// Cfg configuration for Conman
type Cfg struct {
	logInfo          bool
	suppressWarnings bool
	sourceOrder      []string
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
