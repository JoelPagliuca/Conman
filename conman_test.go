package conman

import (
	"testing"
)

func TestNew(t *testing.T) {
	New(Cfg{LogInfo: true, SourceOrder: []string{SourceEnvironment, SourceDefault}})
}

func TestNew_errors(t *testing.T) {
	cm, err := New(Cfg{SourceOrder: []string{"cmdoesnt-exist"}})
	if cm != nil || err == nil {
		t.Errorf("New should have returned (nil, error)")
	}
}
