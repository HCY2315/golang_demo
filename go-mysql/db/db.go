package db

import (
	"fmt"
	"go-mysql/model"
	"go-mysql/setting"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Manager struct {
	db      *gorm.DB
	initOne sync.Once
}

var defaultManager *Manager

func CreateManager() (*Manager, error) {
	var db *gorm.DB
	if setting.DatabaseSetting.Type == "mysql" {
		var err error
		db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			setting.DatabaseSetting.User,
			setting.DatabaseSetting.Passwd,
			setting.DatabaseSetting.Host,
			setting.DatabaseSetting.Port,
			setting.DatabaseSetting.Table_name))
		if err != nil {
			log.Fatalf("database model.Setup is err %v", err)
			return nil, err
		}
	}
	manager := &Manager{
		db:      db,
		initOne: sync.Once{},
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.AutoMigrate(&model.User{})
	// log.Fatalf("mysql db driver created !")  使用 Fatalf 方法后程序会自动退出
	fmt.Println("mysql db driver created !")
	return manager, nil
}

//初始化数据库
func Setup() {
	var err error
	if defaultManager == nil {
		defaultManager, err = CreateManager()
		fmt.Println("aaa")
		if err != nil {
			fmt.Printf("mysql setup is error:%v\n", err)
		}
	}
	fmt.Printf("db manager has already !\n")
}

//GetManager 获取db控制实例
func GetManager() *Manager {
	if defaultManager == nil {
		Setup()
	}
	return defaultManager
}

//InstanceDB 获取db实例
func (m *Manager) InstanceDB() *gorm.DB {
	return m.db
}
