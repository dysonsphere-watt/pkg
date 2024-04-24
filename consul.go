package pkg

import (
	"fmt"
	"net"
	"sync"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/goravel/framework/facades"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/hertz-contrib/registry/consul"
)

var (
	wg   sync.WaitGroup
	h    *server.Hertz
	info *registry.Info
	r    registry.Registry
)

// Register -- register service to consul
func Register() *server.Hertz {
	cfg := facades.Config()

	//set consul server ip
	consulHost := cfg.GetString("CONSUL_HOST", "127.0.0.1")
	consulPort := cfg.GetString("CONSUL_PORT", "8500")

	consulCfg := consulapi.DefaultConfig()
	consulCfg.Address = net.JoinHostPort(consulHost, consulPort)
	fmt.Println(consulCfg.Address)
	consulClient, err := consulapi.NewClient(consulCfg)
	if err != nil {
		facades.Log().Panic(err)
		return nil
	}
	wg.Add(2)
	defer wg.Done()

	//set localIP for serviceHost
	serviceHost := cfg.GetString("APP_URL", "http://localhost")
	addr := net.JoinHostPort(serviceHost, cfg.GetString("APP_PORT", "8000"))
	r = consul.NewConsulRegister(consulClient)
	info = &registry.Info{
		ServiceName: cfg.GetString("APP_MODULE", "Demo"),
		Addr:        utils.NewNetAddr("tcp", addr),
		Weight:      10,
		Tags:        nil,
	}

	s := "0.0.0.0:" + cfg.GetString("APP_PORT")
	h = server.Default(
		server.WithHostPorts(s),
		server.WithRegistry(r, info),
	)
	return h

}
