package conman_test

import (
	"github.com/JoelPagliuca/conman"
)

// normal use case, this will try to set `myAppConfig.Port` to:
// * your `PORT` environment variable if it is set
// * your AWS SSM Parameter value set in `/Prod/app/port`
// * then default to `8080`
func Example() {
	var myAppConfig struct {
		Port string `cmssm:"/Prod/app/port" cmenv:"PORT" cmdefault:"8080"`
	}
	cm, _ := conman.New()
	cm.Hydrate(&myAppConfig)
}

func Example_with_options() {
	conman.New(
		// Change the ordering conman tries to load config values
		conman.SetOrder(conman.TagSSM, conman.TagEnvironment),
		// Help find out why your config wasn't loaded Properly
		conman.EnableLogging(),
		// Add a strategy to be used by Hydrate
		conman.AddStrategy(
			"cmadd",
			conman.Strategy(
				func(_ *conman.Conman, _ string) (*string, error) {
					return nil, nil
				},
			),
		),
		// Add your own `aws.Config` to be used by Hydrate
		conman.AddAWSConfig(nil),
	)
}
