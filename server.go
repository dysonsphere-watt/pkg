package pkg

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/network"
	"github.com/goravel/framework/facades"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/swagger"

	swaggerfiles "github.com/swaggo/files"
)

// Create a Hertz server instance with our preset parameters
func CreateHzInstance(bodyMaxSize int) *server.Hertz {
	// Load the server config
	cfg := facades.Config()
	s := "0.0.0.0:" + cfg.GetString("APP_PORT")

	serverOptions := []config.Option{
		server.WithHostPorts(s),
		server.WithRedirectTrailingSlash(false),
		server.WithOnConnect(svrconn),
		server.WithMaxRequestBodySize(bodyMaxSize),
	}

	if facades.Config().GetBool("CONSUL_REGISTER", false) {
		fmt.Println("Registering services to Consul...")

		consulOpt, err := ConsulOption()
		if err != nil {
			facades.Log().Errorf("Error with Consul registration: %s", err.Error())
			return nil
		}

		serverOptions = append(serverOptions, consulOpt)
	} else {
		facades.Log().Info("Skipping Consul service registration...")
	}

	h := server.Default(serverOptions...)

	// Setup CORS
	var corsc = cfg.Get("corscfg")
	cor, ok := corsc.(cors.Config)
	if ok {
		// CORS configuration is set in the microservice's /config/cors.go file
		// 跨域等等头部设置在配置文件 /config/cors.go中
		h.Use(cors.New(cor))
	}

	// Write to log
	facades.Log().Info(strings.ToUpper(facades.Config().GetString("APP_MODULE")) + " Server Started:")
	return h
}

// Attach swagger documentation to the Hertz instance
func LinkSwagger(h *server.Hertz, pathPrefix string) {
	url := swagger.URL(pathPrefix + "/docs/doc.json")

	h.GET(pathPrefix+"/docs/*any", swagger.WrapHandler(swaggerfiles.Handler, url))

	h.GET(pathPrefix+"/docs", func(_ context.Context, c *app.RequestContext) {
		c.Redirect(301, []byte(pathPrefix+"/docs/index.html"))
	})
}

// svrconn -- just used self
func svrconn(c context.Context, _ network.Conn) context.Context {
	return c
}
