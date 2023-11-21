package server

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/21 10:23
//

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients/naming"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"grpc-study/demo4-hello-grpc-nacos/server/handler"
	proto "grpc-study/demo4-hello-grpc-nacos/server/pb/hello"
	"net"
)

// @Description
// @Author 代码小学生王木木
// @Date 2023/11/17 16:03
func main() {
	IP := "192.168.1.51"
	Port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", IP, Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	server := grpc.NewServer()
	proto.RegisterHelloServer(server, &handler.HelloService{})
	reflection.Register(server) // 这一行代码很重要！！！！
	// 服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	// 创建 Nacos 客户端配置
	clientConfig := constant.ClientConfig{
		NamespaceId:         "your-namespace-id",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "192.168.1.52",
			Port:   8848,
		},
	}
	// 创建 Nacos 服务发现客户端
	client, err := naming.NewNamingClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	// 注册服务
	Register(IP, Port, "go测试", []string{"test", "fuck"}, "wll15248was", "192.168.1.52:8500")
	err = server.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}
