package main

import (
	"flag"
	"fmt"
	"github.com/sosljuk8/analytics/orm"

	"github.com/sosljuk8/analytics/internal/config"
	"github.com/sosljuk8/analytics/internal/handler"
	"github.com/sosljuk8/analytics/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/analytics.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	store := orm.NewStore(c.Mysql)
	if err := store.Migrate(); err != nil {
		panic(err)
	}

	ctx := svc.NewServiceContext(c, store)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
