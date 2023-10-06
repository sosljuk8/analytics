package svc

import (
	"github.com/sosljuk8/analytics/internal/config"
	"github.com/sosljuk8/analytics/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config  config.Config
	APIAuth rest.Middleware
	Store   Store
}

func NewServiceContext(c config.Config, s Store) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Store:   s,
		APIAuth: middleware.NewAPIAuthMiddleware().Handle,
	}
}

type Store interface {
	Load(query RangeQuery) ([]PrimeTime, error)
}
