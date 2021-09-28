package main

import (
	"flag"
	"ping/app"
)

var (
	limitwork int
)

func main() {
	flag.IntVar(&limitwork, "limit", 1024, "限制运行的进程数")
	flag.Parse()
	app.PingInit(limitwork)
}
