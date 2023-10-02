package config

import {{.authImport}}
import "github.com/estatie/auth/jwt"
import "github.com/estatie/pet/pkg/ant"
import "github.com/estatie/pet/zero"

type Config struct {
App app.Config
Jwt jwt.Config
Mysql ant.Config
rest.RestConf
Instance zero.InstanceConfig
{{.auth}}
{{.jwtTrans}}
}

// A RestConf is a http service config.
func (c Config) Config() app.Config {
return c.App
}

// A MySQL/sqlite ORM configuration.
func (c Config) MysqlConfig() ant.Config {
return c.Mysql
}


// A RestConf is a http service config.
func (c Config) Rest() rest.RestConf {
return c.RestConf
}

// A InstanceConfig is a zero instance config.
func (c Config) InstanceConfig() zero.InstanceConfig {
return c.Instance
}