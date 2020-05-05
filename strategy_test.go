package conman

import (
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws/external"
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

func TestEnvironmentStrategy_prefix(t *testing.T) {
	cm, _ := New(SetEnvPrefix("TEST_"))
	os.Setenv("TEST_token1", "token1value")
	defer os.Unsetenv("_TEST_token1")
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

func TestSSMStrategy(t *testing.T) {
	cm, _ := New()
	if cm.awsConfig != nil {
		t.Errorf("Conman shouldn't have an aws config")
	}
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
	if cm.awsConfig == nil {
		t.Errorf("Conman should now have an aws config")
	}
}

func TestSSMStrategy_prefix(t *testing.T) {
	cm, _ := New(SetSSMPrefix("/conman"))
	var toHydrate struct {
		Token1 string `cmssm:"/test/value1"`
	}
	err := cm.Hydrate(&toHydrate)
	if err != nil {
		t.Errorf("Hydrate Failed: %s", err.Error())
	}
	if toHydrate.Token1 != "ssm-test" {
		t.Errorf("cmssm strategy failed, got %s", toHydrate.Token1)
	}
}

func TestSSMStrategyWithConfig(t *testing.T) {
	a, err := external.LoadDefaultAWSConfig()
	if err != nil {
		t.Errorf("Hydrate Failed: %s", err.Error())
	}
	cm, _ := New(AddAWSConfig(&a))
	if cm.awsConfig != &a {
		t.Errorf("Conman should have been given the aws config")
	}
	var toHydrate struct {
		Token1 string `cmssm:"/conman/test/value1"`
	}
	err = cm.Hydrate(&toHydrate)
	if err != nil {
		t.Errorf("Hydrate Failed: %s", err.Error())
	}
	if cm.awsConfig != &a {
		t.Errorf("Conman should still have the aws config")
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
