package main

import (
	"fmt"
	"log"
	"net/rpc"
	"protobuf-introduction/hello"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	s := hello.String{Value: "hello"}
	err = client.Call("chaoyue.Hello", s, &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
