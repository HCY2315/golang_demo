package dao

import (
	"book-server/db"
	"book-server/db/models"
	"fmt"

	"github.com/jinzhu/gorm"
)

type BookListDaoImpl struct {
	DB *gorm.DB
}

func BookListDao(m *db.MysqlMgr) *BookListDaoImpl {
	return &BookListDaoImpl{
		DB: m.InstanceDB(),
	}
}

// ADD 添加数据
func (m *BookListDaoImpl) ADD(mo models.TableInterface) error {
	books := mo.(*models.BookList)
	var oldRecord models.BookList
	if ok := m.DB.Where("book_name = ?", books.BookName).Find(&oldRecord).RecordNotFound(); ok {
		if err := m.DB.Create(books).Error; err != nil {
			fmt.Println(fmt.Sprintf("Add book append error: %v", err))
			return err
		}
	} else {
		fmt.Println(fmt.Sprintf("书籍[ %s ]已存在", books.BookName))
		return fmt.Errorf("书籍[ %s ]已存在", books.BookName)
	}
	return nil
}

//Update 修改数据
func (m *BookListDaoImpl) Update(mo models.TableInterface) error {
	books := mo.(*models.BookList)
	var oldRecord models.BookList
	if ok := m.DB.Where("book_name = ?", books.BookName).Find(&oldRecord).RecordNotFound(); ok {
		if err := m.DB.Save(books).Error; err != nil {
			fmt.Println(fmt.Sprintf("update book err, %v", err))
			return fmt.Errorf("update book err, %v", err)
		}
	} else {
		fmt.Println(fmt.Sprintf("书籍[ %s ]已存在", books.BookName))
		return fmt.Errorf("书籍[ %s ]已存在", books.BookName)
	}
	return nil
}

//GetBookId GetBookId
func (m *BookListDaoImpl) GetBookId(id int) (*models.BookList, error) {
	var booklist models.BookList
	err := m.DB.Where("id = ?", id).Find(&booklist).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println(fmt.Sprintf("get book by id err :%v", err))
		return nil, err
	}
	return &booklist, nil
}
