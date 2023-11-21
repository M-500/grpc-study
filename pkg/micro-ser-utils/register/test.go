package register

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/21 13:57
//

func main() {
	// 使用 Consul
	consulRegistry, err := NewConsulServiceRegistry("localhost:8500")
	if err != nil {
		log.Fatal("Failed to create Consul registry: %v", err)
	}

	// 或者使用 etcd
	etcdRegistry, err := NewEtcdServiceRegistry("localhost:2379")
	if err != nil {
		log.Fatal("Failed to create etcd registry: %v", err)
	}

	// 注册服务
	err = consulRegistry.RegisterService("example-service", "127.0.0.1", 8080)
	if err != nil {
		log.Fatal("Failed to register service: %v", err)
	}

	// 或者使用 etcd
	err = etcdRegistry.RegisterService("example-service", "127.0.0.1", 8080)
	if err != nil {
		log.Fatal("Failed to register service: %v", err)
	}

	// 在程序结束时注销服务
	defer func() {
		err := consulRegistry.DeregisterService("example-service")
		if err != nil {
			log.Printf("Failed to deregister service: %v", err)
		}

		// 或者使用 etcd
		err = etcdRegistry.DeregisterService("example-service")
		if err != nil {
			log.Printf("Failed to deregister service: %v", err)
		}
	}()

	// 等待程序终止信号
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)
	<-stopChan
}
