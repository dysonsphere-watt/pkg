package pkg

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/network"
	"github.com/goravel/framework/facades"
	"github.com/hertz-contrib/cors"
)

func Boot(bodyMaxSize int) *server.Hertz {

	// Load the server config
	con := facades.Config()
	s := "0.0.0.0:" + con.GetString("APP_PORT")

	if facades.Config().GetString("CONSUL_HOST") != "" {
		fmt.Println("Registering services to Consul")
		h = Register(bodyMaxSize)
	} else {
		fmt.Println("Skipping Consul service registration")
		h = server.Default(
			server.WithHostPorts(s),
			server.WithRedirectTrailingSlash(false),
			server.WithOnConnect(svrconn),
			server.WithMaxRequestBodySize(bodyMaxSize),
		)
	}

	// set cors
	var corsc = con.Get("corscfg")

	// 转换
	cor, ok := corsc.(cors.Config)
	if ok {
		// 跨域等等头部设置在配置文件 /config/cors.go中
		h.Use(cors.New(cor))
	}

	// Write to log
	facades.Log().Info(strings.ToUpper(facades.Config().GetString("APP_MODULE")) + " Server Started:")
	return h
}

// svrconn -- just used self
func svrconn(c context.Context, _ network.Conn) context.Context {
	return c
}
