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
			Addr:               envx.ValueByEnv("MYSQL_ADDR", "localhost:3306").String(),
			User:               envx.ValueByEnv("MYSQL_USER", "root").String(),
			Password:           envx.ValueByEnv("MYSQL_PASSWORD", "").String(),
			DbName:             envx.ValueByEnv("MYSQL_DATABASE", "DATABASE").String(),
			MaxOpenConnections: 64,
			MaxIdleConnections: 8,
			MaxLifetime:        600,
			ReadConfig: []gormx.Config{
				{
					Addr: envx.ValueByEnv("MYSQL_READ_ADDR", "localhost:3306").String(),
				},
			},
		},
		Redis: redisx.Config{
			Addr: envx.ValueByEnv("REDIS_ADDR", "localhost:6379").String(),
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
