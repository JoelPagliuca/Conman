package conman_test

import (
	"github.com/JoelPagliuca/conman"
)

// A normal use case. Will set myAppConfig.Port to:
// 	* your PORT environment variable if it is set
// 	* your AWS SSM Parameter value set in /Prod/app/port
// 	* then default to 8080
//
// Make sure your config fields are Exported
// reflect can't set private fields ¯\_(ツ)_/¯
func Example() {
	var myAppConfig struct {
		Port string `cmssm:"/Prod/app/port" cmenv:"PORT" cmdefault:"8080"`
	}

	cm, _ := conman.New()
	cm.Hydrate(&myAppConfig)
}

func Example_options() {
	conman.New(
		// Use APP_ as a prefix for all env values
		conman.SetEnvPrefix("APP_"),
		// Use /app-name as a prefix for all ssm names
		conman.SetSSMPrefix("/app-name"),
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
