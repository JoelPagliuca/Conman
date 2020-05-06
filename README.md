# Conman
```diff
+ Tag a config to remind them to stay ｈｙｄｒａｔｅｄ
```
[![Godoc reference](https://godoc.org/github.com/JoelPagliuca/conman?status.svg)](http://godoc.org/github.com/JoelPagliuca/conman)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/JoelPagliuca/conman)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/JoelPagliuca/conman?color=green)](https://github.com/JoelPagliuca/Conman/releases/latest)

### Quickstart
```go

import "github.com/JoelPagliuca/conman"

var myAppConfig struct {
	Port string `cmssm:"/Prod/app/port" cmenv:"PORT" cmdefault:"8080"`
}

func init() {
	cm, _ := conman.New()
	cm.Hydrate(&myAppConfig)
}
```
This will try to set `myAppConfig.Port` to:
* your `PORT` environment variable if it is set
* your AWS SSM Parameter value set in `/Prod/app/port`
* then default to `8080`

Check the [godoc](http://godoc.org/github.com/JoelPagliuca/conman) for more
