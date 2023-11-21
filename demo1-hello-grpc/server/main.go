package server

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/21 10:23
//

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"grpc-study/demo1-hello-grpc/server/handler"
	proto "grpc-study/demo1-hello-grpc/server/pb/hello"
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

	// 注册服务
	Register(IP, Port, "go测试", []string{"test", "fuck"}, "wll15248was", "192.168.1.52:8500")
	err = server.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}

func Register(serverIP string, serverPort int, serverName string, serverTags []string, serverID string, consulConnStr string) {
	cfg := api.DefaultConfig()
	cfg.Address = consulConnStr // 这里是Consul的IP和地址
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	// 生成对应的检查对象
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", serverIP, serverPort), // 服务端的健康检查 支持HTTP
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "3s",
	}
	// 生成注册对象
	reg := new(api.AgentServiceRegistration)
	reg.Address = serverIP
	reg.ID = serverID
	reg.Port = serverPort
	reg.Name = serverName
	reg.Tags = serverTags
	reg.Check = check

	err = client.Agent().ServiceRegister(reg)
	if err != nil {
		panic(err)
	}
}
