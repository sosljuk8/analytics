package svc

import (
    "context"
    "github.com/estatie/id/pkg/auth/jwt"
    "github.com/zeromicro/go-zero/rest"
    "[[.LazyDepsImport]]"
)

type ServiceContext struct {
	Config {{.config}}
    Deps app.Deps
	{{.middleware}}
}

func NewServiceContext(c {{.config}}, d app.Deps) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Deps: d,
		{{.middlewareAssignment}}
	}
}

// Scheduler returns the scheduler instance for the service-as-worker mode, see zero.ModeConf.
func (s *ServiceContext) Scheduler(_ context.Context) *gocron.Scheduler {
	return nil
}

// Shutdown gracefully shuts down the service.
func (s *ServiceContext) Shutdown(ctx context.Context) error {

sch := s.Scheduler(ctx)
if sch != nil {
sch.Stop()
fmt.Println("Scheduler stopped.")
}

fmt.Println("Service stopped.")

return nil
}
