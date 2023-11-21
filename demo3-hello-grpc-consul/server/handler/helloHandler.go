package handler

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/21 10:39
//

import (
	"context"
	"fmt"
	"grpc-study/demo3-hello-grpc-consul/server/pb/hello"
)

type HelloService struct {
	proto.UnimplementedHelloServer
}

func (h *HelloService) SayHello(ctx context.Context, req *proto.HelloReq) (res *proto.HelloResp, err error) {
	res = new(proto.HelloResp)
	fmt.Println("发起了调用！")
	res.Value = "你好呀" + req.Key
	return res, nil
}
