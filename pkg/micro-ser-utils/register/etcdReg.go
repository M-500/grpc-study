package register

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/21 13:54
//

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"log"
	"strconv"
	"time"
)

// EtcdServiceRegistry 是基于 etcd 的微服务注册和注销实现
type EtcdServiceRegistry struct {
	client *clientv3.Client
}

// NewEtcdServiceRegistry 创建一个 EtcdServiceRegistry 实例
func NewEtcdServiceRegistry(etcdAddress string) (*EtcdServiceRegistry, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{etcdAddress},
	})
	if err != nil {
		return nil, err
	}

	return &EtcdServiceRegistry{client: client}, nil
}

// RegisterService 在 etcd 中注册微服务
func (e *EtcdServiceRegistry) RegisterService(serviceName string, serviceAddress string, servicePort int) error {
	serviceID := serviceName + "-" + strconv.Itoa(int(time.Now().Unix()))

	key := fmt.Sprintf("/services/%s/%s", serviceName, serviceID)
	value := fmt.Sprintf("%s:%d", serviceAddress, servicePort)

	_, err := e.client.Put(context.TODO(), key, value)
	if err != nil {
		log.Printf("Failed to register service: %v", err)
		return err
	}

	log.Printf("Service registered: %s", serviceName)
	return nil
}

// DeregisterService 从 etcd 注销微服务
func (e *EtcdServiceRegistry) DeregisterService(serviceName string) error {
	serviceID := serviceName + "-" + strconv.Itoa(int(time.Now().Unix()))

	key := fmt.Sprintf("/services/%s/%s", serviceName, serviceID)

	_, err := e.client.Delete(context.TODO(), key)
	if err != nil {
		log.Printf("Failed to deregister service: %v", err)
		return err
	}

	log.Printf("Service deregistered: %s", serviceName)
	return nil
}
