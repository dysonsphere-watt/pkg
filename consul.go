package pkg

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/goravel/framework/facades"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/hertz-contrib/registry/consul"
	"net"
	"sync"
)

var (
	wg           sync.WaitGroup
	localIP      = "127.0.0.1"
	h            *server.Hertz
	consulClient *consulapi.Client
	info         *registry.Info
	r            registry.Registry
)

// Register -- register service to consul
func Register() *server.Hertz {

	//set consul server ip
	consulhost := facades.Config().GetString("host", "127.0.0.1")
	consulport := facades.Config().GetString("port", "8500")

	consulcfg := consulapi.DefaultConfig()
	consulcfg.Address = net.JoinHostPort(consulhost, consulport)
	fmt.Println(consulcfg.Address)
	conclient, err := consulapi.NewClient(consulcfg)
	if err != nil {
		facades.Log().Panic(err)
		return nil
	}
	wg.Add(2)

	defer wg.Done()

	//set localIP for host
	localIP = Myip()
	addr := net.JoinHostPort(localIP, facades.Config().GetString("APP_PORT", "8000"))
	r = consul.NewConsulRegister(conclient)
	info = &registry.Info{
		ServiceName: facades.Config().GetString("APP_MODULE", "Demo"),
		Addr:        utils.NewNetAddr("tcp", addr),
		Weight:      10,
		Tags:        nil,
	}

	h = server.Default(
		server.WithHostPorts(addr),
		server.WithRegistry(r, info),
	)
	return h

}

// Myip  -- get myip
func Myip() string {

	// 获取本机所有网络接口的信息
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("获取网络接口信息失败:", err)
		return "0.0.0.0"
	}

	// 遍历每个网络接口，查找非回环接口并获取其 IPv4 地址
	for _, iface := range interfaces {
		// 排除回环接口和无效接口
		if iface.Flags&net.FlagLoopback == 0 && iface.Flags&net.FlagUp != 0 {
			// 获取接口的地址列表
			addrs, err := iface.Addrs()
			if err != nil {
				fmt.Println("获取接口地址失败:", err)
				continue
			}
			// 遍历地址列表，查找 IPv4 地址
			for _, addr := range addrs {
				ipNet, ok := addr.(*net.IPNet)
				if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
					//fmt.Println("本机 IPv4 地址:", ipNet.IP)
					return ipNet.IP.String()
				}
			}
		}
	}
	return "0.0.0.0"
}
