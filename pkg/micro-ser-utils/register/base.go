package register

// @Description
// @Author 代码小学生王木木
// @Date 2023/11/21 13:53
//

// ServiceRegistry 定义了微服务注册和注销的接口
type ServiceRegistry interface {
	RegisterService(serviceName string, serviceAddress string, servicePort int) error
	DeregisterService(serviceName string) error
}
