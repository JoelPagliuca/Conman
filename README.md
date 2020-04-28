# Conman
```diff
+ Tag a config to remind them to stay hydrated
```
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/JoelPagliuca/conman)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/JoelPagliuca/conman)

### Quickstart
```go

import "github.com/JoelPagliuca/conman"

var myAppConfig struct {
	Port string `cmenv:"PORT" cmdefault:"8080"`
}

func init() {
	c := conman.DefaultCfg
	cm := conman.New(c)
	cm.HydrateConfig(&myAppConfig)
}
```
This will set `myAppConfig.Port` to your `PORT` environment variable if it is set or default to `8080`

### Coming soon
* AWS SSM Params
* Custom strategies


Uses `reflect` so watch out
