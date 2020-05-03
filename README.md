# Conman
```diff
+ Tag a config to remind them to stay ｈｙｄｒａｔｅｄ
```
[![Godoc reference](https://godoc.org/github.com/JoelPagliuca/conman?status.svg)](http://godoc.org/github.com/JoelPagliuca/conman)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/JoelPagliuca/conman)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/JoelPagliuca/conman?color=green)](https://github.com/JoelPagliuca/Conman/releases/latest)
![GitHub milestone](https://img.shields.io/github/milestones/progress-percent/JoelPagliuca/conman/2?color=yellow)

### Quickstart
```go

import "github.com/JoelPagliuca/conman"

var myAppConfig struct {
	Port string `cmenv:"PORT" cmdefault:"8080"`
}

func init() {
	cm, _ := conman.New(c)
	cm.Hydrate(&myAppConfig)
}
```
This will set `myAppConfig.Port` to your `PORT` environment variable if it is set or default to `8080`

### Coming soon
* AWS SSM Params
* Custom strategies


Uses `reflect` so watch out
