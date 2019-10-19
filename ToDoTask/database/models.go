package database

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	ID        int
	UserName  string
	FirstName string
	LastName  string
	Email     string
}

type ToDoItem struct {
	gorm.Model

	Id       int `gorm:"primary_key;column:id"`
	Title    string
	Category string
	UserName string
}

type Category struct {
	gorm.Model

	ID    int
	Title string
}

var ToDoItems []*ToDoItem
