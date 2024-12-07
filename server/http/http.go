package http

import (
	"github.com/lascape/gopkg/server"
	httpServer "github.com/lascape/gopkg/server/http"
	"github.com/lascape/goweb/config"
)

func New() server.Server {
	hServer := httpServer.NewServer(func(s *httpServer.Server) {
	},
		httpServer.WithConfig(config.Config.HttpServer),
	)
	return hServer
}
