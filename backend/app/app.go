package main

import (
	"flag"
	"fmt"

	"github.com/xiao-en-5970/Goodminton/backend/app/internal/config"
	"github.com/xiao-en-5970/Goodminton/backend/app/internal/handler"
	"github.com/xiao-en-5970/Goodminton/backend/app/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/app-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
