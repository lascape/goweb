package config

import (
	"github.com/lascape/gopkg/envx"
	"github.com/lascape/gopkg/gormx"
	"github.com/lascape/gopkg/redisx"
	httpServer "github.com/lascape/gopkg/server/http"
)

var Config = func() config {
	var c = config{
		HttpServer: httpServer.Config{
			Name:         "goweb",
			Addr:         "0.0.0.0:8080",
			ReadTimeout:  30,
			WriteTimeout: 30,
		},
		Mysql: gormx.Config{
			Addr:               envx.ValueByEnv("localhost:3306", "MYSQL_ADDR"),
			User:               envx.ValueByEnv("root", "MYSQL_USER"),
			Password:           envx.ValueByEnv("", "MYSQL_PASSWORD"),
			DbName:             envx.ValueByEnv("paymentengine", "MYSQL_DATABASE"),
			MaxOpenConnections: 64,
			MaxIdleConnections: 8,
			MaxLifetime:        600,
			ReadConfig: []gormx.Config{
				{
					Addr: envx.ValueByEnv("localhost:3306", "MYSQL_READ_ADDR"),
				},
			},
		},
		Redis: redisx.Config{
			Addr: envx.ValueByEnv("localhost:6379", "REDIS_ADDR"),
			DB:   0,
		},
	}

	return c
}()

type config struct {
	HttpServer httpServer.Config `json:"http_server" yaml:"http_server"`
	Mysql      gormx.Config      `json:"mysql" yaml:"mysql"`
	Redis      redisx.Config     `json:"redis" yaml:"redis"`
}
