package register

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/21 13:53
//
import (
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/consul/api"
)

// ConsulServiceRegistry 是基于 Consul 的微服务注册和注销实现
type ConsulServiceRegistry struct {
	client *api.Client
}

// NewConsulServiceRegistry 创建一个 ConsulServiceRegistry 实例
func NewConsulServiceRegistry(consulAddress string) (*ConsulServiceRegistry, error) {
	config := api.DefaultConfig()
	config.Address = consulAddress

	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &ConsulServiceRegistry{client: client}, nil
}

// RegisterService 在 Consul 中注册微服务
func (c *ConsulServiceRegistry) RegisterService(serviceName string, serviceAddress string, servicePort int) error {
	serviceID := serviceName + "-" + strconv.Itoa(int(time.Now().Unix()))

	registration := &api.AgentServiceRegistration{
		ID:      serviceID,
		Name:    serviceName,
		Address: serviceAddress,
		Port:    servicePort,
	}

	err := c.client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Printf("Failed to register service: %v", err)
		return err
	}

	log.Printf("Service registered: %s", serviceName)
	return nil
}

// DeregisterService 从 Consul 注销微服务
func (c *ConsulServiceRegistry) DeregisterService(serviceName string) error {
	serviceID := serviceName + "-" + strconv.Itoa(int(time.Now().Unix()))

	err := c.client.Agent().ServiceDeregister(serviceID)
	if err != nil {
		log.Printf("Failed to deregister service: %v", err)
		return err
	}

	log.Printf("Service deregistered: %s", serviceName)
	return nil
}
