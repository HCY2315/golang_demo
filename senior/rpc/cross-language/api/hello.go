package api

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

const ServerName = "chaoyue"

type ServerInterface interface {
	Hello(request string, replay *string) error
}

type Client struct {
	*rpc.Client
}

// 客户端需要调用的方法，这里指需要再server端实现哪些方法
func (p *Client) Hello(request string, replay *string) error {
	return p.Client.Call(ServerName+".Hello", request, replay)
}

// 客户端连接服务端
func DialHelloServer(network, address string) (*Client, error) {
	c, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(c))
	return &Client{Client: client}, nil
}

func RegisterName(svc ServerInterface) error {
	return rpc.RegisterName(ServerName, svc)
}
