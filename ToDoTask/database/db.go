package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
)

// Database contains access to database
type Database struct {
	*gorm.DB
}

// OpenDatabase opens database
func OpenDatabase() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "user=postgres password=vasil1306 dbname=toDoItemsDatabase sslmode=disable")
	if err != nil {
		return nil, errors.Wrap(err, "Unable to connect to database")
	}
	db.LogMode(true)
	return db, err
}

// CreateTodo creates a todo item
func CreateTodo(db *gorm.DB, todo *ToDoItem) error {
	db.LogMode(true)
	return errors.Wrap(db.Create(todo).Error, "unable to create todo")
}

// DeleteTodo deletes a todo item
func DeleteTodo(db *gorm.DB, id int) error {
	return errors.Wrap(db.Delete(&ToDoItem{}, id).Error, "unable to delete todo")
}

// ListTodoItems lists all existing items
func ListTodoItems(db *gorm.DB) (*gorm.DB, error) {
	return db.Find(&ToDoItems), nil
}
