package conman

import (
	"os"
	"testing"
)

func TestDefaultStrategy(t *testing.T) {
	cm := New(Cfg{logInfo: true, suppressWarnings: false, sourceOrder: []string{SourceEnvironment, SourceDefault}})
	var toHydrate struct {
		Token1     string `cmdefault:"default"`
		OtherToken string
	}
	err := cm.HydrateConfig(&toHydrate)
	if err != nil {
		t.Errorf("HydrateConfig Failed: %s", err.Error())
	}
	if toHydrate.Token1 != "default" {
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
		t.Errorf("cmenv strategy failed, got %s", toHydrate.Token1)
	}
}

func TestStrategyOrdering(t *testing.T) {
	cm := New(Cfg{logInfo: true, suppressWarnings: false, sourceOrder: []string{SourceEnvironment, SourceDefault}})
	var toHydrate struct {
		Token1 string `cmenv:"token1" cmdefault:"default"`
	}
	err := cm.HydrateConfig(&toHydrate)
	if err != nil {
		t.Errorf("HydrateConfig Failed: %s", err.Error())
	}
	if toHydrate.Token1 != "default" {
		t.Errorf("cmdefault strategy failed, got \"%s\"", toHydrate.Token1)
	}
	defer os.Unsetenv("token1")
	os.Setenv("token1", "token1value")
	cm.HydrateConfig(&toHydrate)
	if toHydrate.Token1 != "token1value" {
		t.Errorf("cmenv strategy failed, got \"%s\"", toHydrate.Token1)
	}
}
