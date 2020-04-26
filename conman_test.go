package conman

import (
	"testing"
)

func TestNew(t *testing.T) {
	New(Cfg{logInfo: true, suppressWarnings: false, sourceOrder: []string{SourceEnvironment, SourceDefault}})
}
