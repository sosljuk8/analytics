package svc

import (
	"github.com/sosljuk8/analytics/internal/config"
	"github.com/sosljuk8/analytics/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config  config.Config
	APIAuth rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		APIAuth: middleware.NewAPIAuthMiddleware().Handle,
	}
}
