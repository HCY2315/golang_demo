package api

import "net/rpc"

const HelloServiceName = "HelloService"

type HelloServer interface {
	Hello(request string, reply *string) error
}

type HelloClient struct {
	*rpc.Client
}

// 服务端需要实现 Hello 方法的函数
func (p *HelloClient) Hello(request string, reply *string) error {
	return p.Call(HelloServiceName+".Hello", request, &reply)
}

// 注册函数
func RegisterHelloService(svc HelloServer) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

// 客户端连接方法
func DialHelloService(network string, address string) (*HelloClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloClient{Client: c}, nil
}
