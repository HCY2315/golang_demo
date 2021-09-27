package main

import (
	"fmt"
	"log"
	"rpc-hello/api"
)

func main() {
	client, err := api.DialHelloService("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatalln("dialing:", err)
		return
	}
	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatalln("eror", err)
		return
	}
	fmt.Println(reply)
}
