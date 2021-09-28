package main

import (
	"fmt"
	"go-mysql/db"
	"go-mysql/model"
	"go-mysql/setting"
)

func init() {
	setting.Setup()
	db.Setup()
}

func main() {
	fmt.Println("success")
	mysql := db.GetManager().InstanceDB()
	u1 := model.User{Name: "aaa", Age: 18}
	u2 := model.User{Name: "bbb", Age: 19}
	u3 := model.User{Name: "ccc", Age: 20}
	mysql.Create(&u1)
	mysql.Create(&u2)
	mysql.Create(&u3)
	d1 := model.User{}
	d1.ID = 1
	mysql.Debug().Delete(&model.User{}, "name = ?", d1.ID)
}
