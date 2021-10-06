package main

import (
	"log"
	"net"
	"net/rpc"
	"rpc-hello/api"
)

type HelloServer struct{}

func (p *HelloServer) Hello(request string, reply *string) error {
	*reply = "helloWold:" + request
	return nil
}

func main() {
	api.RegisterHelloService(new(HelloServer))

	listenr, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalln("listen port error:", err)
		return
	}

	for {
		conn, err := listenr.Accept()
		if err != nil {
			log.Fatalln("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}
