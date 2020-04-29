package conman

import (
	"testing"
)

func TestNew(t *testing.T) {
	New(Cfg{LogInfo: true, SuppressWarnings: false, SourceOrder: []string{SourceEnvironment, SourceDefault}})
}
