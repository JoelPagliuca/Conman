package conman

import "testing"

func TestConman_HydrateConfig(t *testing.T) {
	cm := New(Cfg{logInfo: true, suppressWarnings: false})
	var toHydrate struct {
		GithubToken string
		OtherConfig string
		topicArn    string
	}
	err := cm.HydrateConfig(&toHydrate)
	if err != nil {
		t.Errorf("HydrateConfig Failed: %s", err.Error())
	}
}
