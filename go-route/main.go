package main

import (
	"book-server/db"
	"book-server/pkg/setting"
	"book-server/routers"
)

func main() {
	setting.Setup()
	db.Setup()
	r := routers.InitRouter()
	r.Run(":80")
}
