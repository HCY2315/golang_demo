package models

type BookList struct {
	Model
	BookName string `gorm:"type:varchar(16)" json:"book_name"`
	Price    string `gorm:"type:varchar(16)" json:"price"`
	BookType string `gorm:"type:varchar(16)" json:"book_type"`
}

func (bl *BookList) TableName() string {
	return "book_list"
}
