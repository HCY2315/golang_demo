package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloServer struct{}

func (p *HelloServer) Hello(request string, reply *string) error {
	*reply = "test:" + request
	return nil
}

type HelloServer1 struct{}

func (p *HelloServer1) Hello1(request string, reply *string) error {
	*reply = "test1:" + request
	return nil
}

// curl localhost:1234/jsonrpc --data '{"method":"chaoyue.Hello","params":["hello"],"id":0}'
// 返回信息： {"id":0,"result":"test:hello","error":null}

// curl localhost:1234/jsonrpc --data '{"method":"chaoyue1.Hello1","params":["hello"],"id":0}'
// 返回信息： {"id":0,"result":"test1:hello","error":null}

func main() {
	rpc.RegisterName("chaoyue", new(HelloServer))
	rpc.RegisterName("chaoyue1", new(HelloServer1))
	http.HandleFunc("/jsonrpc", HttpHello)
	http.ListenAndServe(":1234", nil)
}

func HttpHello(w http.ResponseWriter, r *http.Request) {
	var conn io.ReadWriteCloser = struct {
		io.Writer
		io.ReadCloser
	}{
		ReadCloser: r.Body,
		Writer:     w,
	}
	rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
}
