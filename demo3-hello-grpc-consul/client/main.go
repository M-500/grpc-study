package main

import (
	"context"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	proto "grpc-study/demo3-hello-grpc-consul/client/pb/hello"
	"log"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/21 10:24
//

var consulAddress = "consul://192.168.1.52:8500" // 实际替换自己的Consul地址
var serviceName = "grpc-consul-test"             // 这个要和Server端保持一致

func main() {
	// "consul://192.168.10.130:8500/user_srv?wait=14s&tag=srv",
	target := fmt.Sprintf("%s/%s?wait=14s", consulAddress, serviceName)
	conn, err := grpc.Dial(
		//consul网络必须是通的   user_srv表示服务 wait:超时 tag是consul的tag  可以不填
		target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), //轮询法   必须这样写   grpc在向consul发起请求时会遵循轮询法
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 发起10次请求
	for i := 0; i < 10; i++ {
		// 创建 gRPC 客户端
		client := proto.NewHelloClient(conn)

		// 准备请求数据
		request := &proto.HelloReq{
			Key: fmt.Sprintf("傻逼 %d 号", i),
		}

		// 调用 gRPC 服务
		response, err := client.SayHello(context.Background(), request)
		if err != nil {
			log.Fatalf("Failed to call gRPC service: %v", err)
		}

		// 处理响应
		fmt.Printf("Response from server: %v\n", response)
	}
}
