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

// @title PaymentEngine
// @version 1.0
// @description 这是一个管理银行以及银行卡的一体化系统
// @termsOfService http://example.com/terms/

// @host localhost:7068
// @BasePath /
func main() {
	server.Run(server.WithServer(http.New()))
}
