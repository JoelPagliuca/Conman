package conman

import (
	"os"
	"testing"
)

func TestDefaultStrategy(t *testing.T) {
	cm, _ := New()
	var toHydrate struct {
		Token1     string `cmdefault:"default"`
		OtherToken string
	}
	err := cm.Hydrate(&toHydrate)
	if err != nil {
		t.Errorf("Hydrate Failed: %s", err.Error())
	}
	if toHydrate.Token1 != "default" {
		t.Errorf("cmdefault strategy failed %s", toHydrate.Token1)
	}
	if toHydrate.OtherToken != "" {
		t.Errorf("cmdefault strategy failed, got %s", toHydrate.OtherToken)
	}
}

func TestEnvironmentStrategy(t *testing.T) {
	cm, _ := New()
	os.Setenv("token1", "token1value")
	defer os.Unsetenv("token1")
	var toHydrate struct {
		Token1 string `cmenv:"token1"`
	}
	err := cm.Hydrate(&toHydrate)
	if err != nil {
		t.Errorf("Hydrate Failed: %s", err.Error())
	}
	if toHydrate.Token1 != "token1value" {
		t.Errorf("cmenv strategy failed, got %s", toHydrate.Token1)
	}
}

func TestEnvironmentStrategy_errors(t *testing.T) {
	cm, _ := New()
	os.Unsetenv("notset")
	var toHydrate struct {
		Token1 string `cmenv:"notset"`
	}
	err := cm.Hydrate(&toHydrate)
	if err == nil {
		t.Errorf("Hydrate should have errored but didn't")
	}
}

func TestSSMStrategy(t *testing.T) {
	cm, _ := New()
	var toHydrate struct {
		Token1 string `cmssm:"/conman/test/value1"`
	}
	err := cm.Hydrate(&toHydrate)
	if err != nil {
		t.Errorf("Hydrate Failed: %s", err.Error())
	}
	if toHydrate.Token1 != "ssm-test" {
		t.Errorf("cmssm strategy failed, got %s", toHydrate.Token1)
	}
}

func TestSSMStrategy_errors(t *testing.T) {
	cm, _ := New()
	var toHydrate struct {
		Token1 string `cmssm:"notset"`
	}
	err := cm.Hydrate(&toHydrate)
	if err == nil {
		t.Errorf("Hydrate should have errored but didn't")
	}
}

func TestStrategyOrdering(t *testing.T) {
	cm, _ := New(SetOrder(TagEnvironment, TagDefault))
	var toHydrate struct {
		Token1 string `cmenv:"token1" cmdefault:"default"`
	}
	err := cm.Hydrate(&toHydrate)
	if err != nil {
		t.Errorf("Hydrate Failed: %s", err.Error())
	}
	if toHydrate.Token1 != "default" {
		t.Errorf("cmdefault strategy failed, got \"%s\"", toHydrate.Token1)
	}
	defer os.Unsetenv("token1")
	os.Setenv("token1", "token1value")
	cm.Hydrate(&toHydrate)
	if toHydrate.Token1 != "token1value" {
		t.Errorf("cmenv strategy failed, got \"%s\"", toHydrate.Token1)
	}
}
