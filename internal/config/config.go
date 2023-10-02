package config

import (
	"github.com/estatie/auth/jwt"
	"github.com/estatie/pet/pkg/ant"
	"github.com/estatie/pet/zero"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	Jwt   jwt.Config
	Mysql ant.Config
	rest.RestConf
	Instance zero.InstanceConfig
}
