package main

import (
	"log"
	"net"
	"net/rpc"
	"protobuf-introduction/hello"
)

// 基于新的 string 类型，重新实现HelloString 服务
type HelloService struct{}

// 其中 Hello() 方法中的输入参数和输出参数均适用protobuf定义的String类型表达
func (p *HelloService) Hello(request *hello.String, reply *hello.String) error {
	reply.Value = "hello:" + request.GetValue()
	return nil
}

func main() {
	rpc.RegisterName("chaoyue", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
			return
		}
		go rpc.ServeConn(conn)
	}
}
