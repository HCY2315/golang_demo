package main

import (
	"cross-language/api"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Server struct{}

func (p *Server) Hello(request string, replay *string) error {
	*replay = "test:" + request
	return nil
}

func main() {
	api.RegisterName(new(Server))

	listenr, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalln("listen error:", err)
		return
	}
	for {
		conn, err := listenr.Accept()
		if err != nil {
			log.Fatalln("accept error:", err)
			return
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
