package main

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/21 10:23
//

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"grpc-study/demo3-hello-grpc-consul/server/handler"
	proto "grpc-study/demo3-hello-grpc-consul/server/pb/hello"
)

// @Description
// @Author 代码小学生王木木
// @Date 2023/11/17 16:03
func main() {
	IP := flag.String("ip", "192.168.1.51", "IP地址")
	Port := flag.Int("port", 50051, "端口号")
	flag.Parse()
	fmt.Println("运行端口号为:", *Port)
	srvID := UUID()
	serverName := "grpc-consul-test"
	consulAddr := "192.168.1.52:8500" // 本地consul的地址
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	server := grpc.NewServer()
	proto.RegisterHelloServer(server, &handler.HelloService{})
	// 服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	// 注册服务到Consul
	Register(*IP, *Port, serverName, []string{"test", "fuck"}, srvID, "192.168.1.52:8500")

	go func() {
		err = server.Serve(lis)
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()
	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	// 注销Consul的服务
	if err = DeregisterService(srvID, consulAddr); err != nil {
		fmt.Println("error：服务注销失败")
	}
	fmt.Println("服务注销成功")
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

func DeregisterService(srvID string, consulConnStr string) error {
	config := api.DefaultConfig()
	config.Address = consulConnStr // 这里是Consul的IP和地址
	client, err := api.NewClient(config)
	if err != nil {
		return err
	}

	agent := client.Agent()

	err = agent.ServiceDeregister(srvID)
	if err != nil {
		return err
	}

	fmt.Println("Service deregistered successfully!")
	return nil
}

func UUID() string {
	uid := uuid.NewV4()
	var err error
	u := uuid.Must(uid, err)
	return strings.ReplaceAll(u.String(), "-", "")
}
