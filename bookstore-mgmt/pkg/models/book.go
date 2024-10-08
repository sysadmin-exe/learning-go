package models

import (
	"github.com/jinzhu/gorm"
	"github.com/sysadmin-exe/learning-go/bookstore-mgmt/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (NewBook *Book) CreateBook() *Book {
	db.NewRecord(NewBook)
	db.Create(&NewBook)
	return NewBook
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(ID int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID = ?", ID).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID = ?", ID).Delete(&book)
	return book
}
