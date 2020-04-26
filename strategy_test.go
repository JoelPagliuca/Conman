package conman

import (
	"os"
	"testing"
)

func TestDefaultStrategy(t *testing.T) {
	cm := New(Cfg{logInfo: true, suppressWarnings: false, sourceOrder: []string{SourceEnvironment, SourceDefault}})
	var toHydrate struct {
		Token1     string `cmdefault:"token1"`
		OtherToken string
	}
	err := cm.HydrateConfig(&toHydrate)
	if err != nil {
		t.Errorf("HydrateConfig Failed: %s", err.Error())
	}
	if toHydrate.Token1 != "token1" {
		t.Errorf("cmdefault strategy failed %s", toHydrate.Token1)
	}
	if toHydrate.OtherToken != "" {
		t.Errorf("cmdefault strategy failed, got %s", toHydrate.OtherToken)
	}
}

func TestEnvironmentStrategy(t *testing.T) {
	cm := New(Cfg{logInfo: true, suppressWarnings: false, sourceOrder: []string{SourceEnvironment, SourceDefault}})
	os.Setenv("token1", "token1value")
	defer os.Unsetenv("token1")
	var toHydrate struct {
		Token1 string `cmenv:"token1"`
	}
	err := cm.HydrateConfig(&toHydrate)
	if err != nil {
		t.Errorf("HydrateConfig Failed: %s", err.Error())
	}
	if toHydrate.Token1 != "token1value" {
		t.Errorf("cmdefault strategy failed, got %s", toHydrate.Token1)
	}
}

func TestStrategyOrdering(t *testing.T) {
	cm := New(Cfg{logInfo: true, suppressWarnings: false, sourceOrder: []string{SourceEnvironment, SourceDefault}})
	os.Setenv("token1", "token1value")
	var toHydrate struct {
		Token1 string `cmenv:"token1" cmdefault:"token1"`
	}
	err := cm.HydrateConfig(&toHydrate)
	if err != nil {
		t.Errorf("HydrateConfig Failed: %s", err.Error())
	}
	if toHydrate.Token1 != "token1value" {
		t.Errorf("cmdefault strategy failed, got \"%s\"", toHydrate.Token1)
	}
}
