package main

import (
	"cross-language/api"
	"fmt"
	"log"
)

func main() {
	conn, err := api.DialHelloServer("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatalln("net Dial error:", err)
		return
	}
	var reply string
	err = conn.Hello("hello", &reply)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(reply)
}
