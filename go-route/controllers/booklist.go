package controllers

import (
	"book-server/dao"
	"book-server/db"
	"book-server/db/models"
)

//BookList BookList
type BookList struct {
	Result interface{}
	// Maps     interface{}
	// Mps      map[string]interface{}
	ID       int `json:"id"`
	PageNum  int `json:"page_num"`
	PageSize int `json:"page_size"`
}

//ADD ADD
func (bl *BookList) ADD(mo models.TableInterface) error {
	return dao.BookListDao(db.GetManager()).ADD(mo)
}

//Update Update
func (bl *BookList) Update(mo models.TableInterface) error {
	return dao.BookListDao(db.GetManager()).Update(mo)
}

//GetBookList 通过post方法中传入的id向数据库查找整条sql,并返回到 Result 中
func (bl *BookList) GetBookList() error {
	mc, err := dao.BookListDao(db.GetManager()).GetBookId(bl.ID)
	if err != nil {
		return err
	}
	bl.Result = mc
	return nil
}
