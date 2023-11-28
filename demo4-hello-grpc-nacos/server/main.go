package main

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/21 10:23
//

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
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
	//clientConfig := constant.ClientConfig{
	//	TimeoutMs            uint64 // 请求Nacos服务端的超时时间，默认是10000ms
	//	NamespaceId          string // ACM的命名空间Id
	//	Endpoint             string // 当使用ACM时，需要该配置. https://help.aliyun.com/document_detail/130146.html
	//	RegionId             string // ACM&KMS的regionId，用于配置中心的鉴权
	//	AccessKey            string // ACM&KMS的AccessKey，用于配置中心的鉴权
	//	SecretKey            string // ACM&KMS的SecretKey，用于配置中心的鉴权
	//	OpenKMS              bool   // 是否开启kms，默认不开启，kms可以参考文档 https://help.aliyun.com/product/28933.html
	//	// 同时DataId必须以"cipher-"作为前缀才会启动加解密逻辑
	//	CacheDir             string // 缓存service信息的目录，默认是当前运行目录
	//	UpdateThreadNum      int    // 监听service变化的并发数，默认20
	//	NotLoadCacheAtStart  bool   // 在启动的时候不读取缓存在CacheDir的service信息
	//	UpdateCacheWhenEmpty bool   // 当service返回的实例列表为空时，不更新缓存，用于推空保护
	//	Username             string // Nacos服务端的API鉴权Username
	//	Password             string // Nacos服务端的API鉴权Password
	//	LogDir               string // 日志存储路径
	//	RotateTime           string // 日志轮转周期，比如：30m, 1h, 24h, 默认是24h
	//	MaxAge               int64  // 日志最大文件数，默认3
	//	LogLevel             string // 日志默认级别，值必须是：debug,info,warn,error，默认值是info
	//}
	clientConfig := constant.ClientConfig{
		NamespaceId:         "5d55b29c-129d-405c-95cc-089acc9d95b5",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	// 创建服务注册客户端
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig: &clientConfig,
			ServerConfigs: []constant.ServerConfig{
				{
					IpAddr:      "192.168.1.52",
					ContextPath: "/nacos",
					Port:        8848,
				},
			},
		},
	)
	// 注册服务
	params := vo.RegisterInstanceParam{
		Ip:          IP,
		Port:        uint64(Port),
		ServiceName: "test.grpc.service",
		Weight:      10,
		ClusterName: "cluster-a",
		Metadata:    map[string]string{"idc": "shenzhen"},
	}
	success, err := namingClient.RegisterInstance(params)
	if err != nil {
		panic(err)
	}
	fmt.Println(success)

	// 启动服务
	err = server.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}
