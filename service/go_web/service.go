package go_web

import (
	"context"
	"github.com/lascape/goweb/server/http/params"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GoWeb(ctx context.Context, req params.ReqGoWeb) error {
	return nil
}
