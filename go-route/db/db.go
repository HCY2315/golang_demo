package db

import (
	"book-server/db/models"
	"book-server/pkg/setting"
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MysqlMgr struct {
	db      *gorm.DB
	initOne sync.Once
	models  []models.TableInterface
}

//InstanceDB 获取db实例
func (m *MysqlMgr) InstanceDB() *gorm.DB {
	return m.db
}

// CreateManager 创建manager
func CreateManager() (*MysqlMgr, error) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.MysqlSetting.User,
		setting.MysqlSetting.Password,
		setting.MysqlSetting.Address,
		setting.MysqlSetting.Port,
		setting.MysqlSetting.Database))
	if err != nil {
		log.Fatalf("book database models.setup err: %v", err)
		return nil, err
	}
	db.SingularTable(true)
	db.Callback().Create().Replace("gorm:update_time_stamp", models.UpdateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", models.UpdateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", models.DeleteCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	mysqlmgr := &MysqlMgr{
		db:      db,
		initOne: sync.Once{},
	}
	// db.SetLogger(mysqlmgr)
	mysqlmgr.RegisterTableModel()
	mysqlmgr.CheckTable()
	log.Println("mysql db driver created.")
	return mysqlmgr, nil
}

// RegisterTableModel 注册表结构
func (m *MysqlMgr) RegisterTableModel() {
	m.models = append(m.models, &models.BookList{})
}

// CheckTable 检测表结构
func (m *MysqlMgr) CheckTable() {
	m.initOne.Do(func() {
		for _, md := range m.models {
			if !m.db.HasTable(md) {
				log.Println("insert", md.TableName())
				m.db.CreateTable(md)
			} else {
				log.Println(md.TableName())
			}
		}
	})
}

var defaultMysqlMgr *MysqlMgr

//Setup 初始化sql
func Setup() {
	var err error
	if defaultMysqlMgr == nil {
		defaultMysqlMgr, err = CreateManager()
		if err != nil {
			fmt.Printf("setup db manager error, %v\n", err)
		}
	}
	fmt.Println("db manager has already init.")
}

// GetManager 获取db控制实例
func GetManager() *MysqlMgr {
	if defaultMysqlMgr == nil {
		Setup()
	}
	return defaultMysqlMgr
}
