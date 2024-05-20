package pkg

import (
	"fmt"
	"net"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/goravel/framework/facades"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/hertz-contrib/registry/consul"
)

var (
	info *registry.Info
	r    registry.Registry
)

func ConsulOption() (config.Option, error) {
	consulOption := config.Option{}

	cfg := facades.Config()

	consulHost := cfg.GetString("CONSUL_HOST", "127.0.0.1")
	consulPort := cfg.GetString("CONSUL_PORT", "8500")

	consulCfg := consulapi.DefaultConfig()
	consulCfg.Address = net.JoinHostPort(consulHost, consulPort)
	fmt.Println(consulCfg.Address)

	consulClient, err := consulapi.NewClient(consulCfg)
	if err != nil {
		facades.Log().Panic(err)
		return consulOption, nil
	}

	// Set localIP for serviceHost
	serviceHost := cfg.GetString("APP_URL", "http://localhost")
	servicePort := cfg.GetString("APP_PUBLIC_PORT", "80")
	addr := net.JoinHostPort(serviceHost, servicePort)
	r = consul.NewConsulRegister(consulClient, consul.WithCheck(&consulapi.AgentServiceCheck{
		Interval:                       "10s",
		Timeout:                        "5s",
		DeregisterCriticalServiceAfter: "1m",
	}))

	info = &registry.Info{
		ServiceName: cfg.GetString("APP_MODULE", "Watt-Generic"),
		Addr:        utils.NewNetAddr("tcp", addr),
		Weight:      10,
		Tags:        nil,
	}

	return server.WithRegistry(r, info), nil
}
