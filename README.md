Instalation: 
1. Make sure Go already [installed](https://golang.org/doc/install) in your machine and setup your [Go Workspace](https://golang.org/doc/code.html#Workspaces)
2. Install [Govendor](https://github.com/kardianos/govendor) `go get github.com/kardianos/govendor`
3. Go to your $GOPATH/src/. Run `go get github.com/501army/go-tv-show`
4. Go to inside $GOPATH/src/github.com/501army/go-tv-show
5. Run `govendor sync` to pull all package needed and wait until finish
6. Adjust your config.json file
7. You're ready to go. Run `go run main.go`.
8. Now your application is running in your specific port in config (default 2323)

Route list:
* `v1/schedule` : list of all tv show
* `v1/search?q=[string]` : search tv show

Depedencies :
- [Viper](https://github.com/spf13/viper)
- [CORS](https://github.com/gin-contrib/cors)