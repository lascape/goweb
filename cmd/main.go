package main

import (
	"github.com/lascape/gopkg/gormx"
	"github.com/lascape/gopkg/logx"
	"github.com/lascape/gopkg/redisx"
	"github.com/lascape/gopkg/server"
	"github.com/lascape/goweb/config"
	"github.com/lascape/goweb/server/http"
	"github.com/lascape/goweb/utils/global"
)

func init() {
	logx.Init(config.Config.HttpServer.Name)
	global.SetRB(redisx.Must(redisx.WithConfig(config.Config.Redis)))
	global.SetDB(gormx.Must(gormx.WithConfig(config.Config.Mysql), gormx.WithDebug()))
}

// @title Web System
// @version 1.0
// @description This Is A Web System
// @termsOfService http://example.com/terms/

// @host localhost:8080
// @BasePath /
func main() {
	server.Run(server.WithServer(http.New()))
}
