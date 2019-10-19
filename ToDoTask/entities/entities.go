package entities

import (
	"ToDoTask/database"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
)

// CreateNewItem receives data from the front-end
func CreateNewItem(w http.ResponseWriter, r *http.Request, db *gorm.DB) error {
	if r.Method == http.MethodPost {
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			panic(err)
		}
		title := r.PostFormValue("title")
		if len(title) == 0 {
			return errors.New("title cannot be empty value")
		}
		category := r.PostFormValue("category")
		if len(category) == 0 {
			return errors.New("category cannot be empty value")
		}
		username := r.PostFormValue("username")
		if len(username) == 0 {
			return errors.New("username cannot be empty value")
		}

		item := &database.ToDoItem{
			UserName: username,
			Category: category,
			Title:    title,
		}
		err := database.CreateTodo(db, item)
		if err != nil {
			fmt.Printf("Error creating item: %s", err.Error())
		}
		db.Save(item)
		return nil
	}

	return nil
}

// DeleteItem deletes item from database found by id
func DeleteItem(w http.ResponseWriter, r *http.Request, db *gorm.DB) error {

	fmt.Println(r)

	idData := r.URL.Query()["id"]
	if len(idData) == 0 {
		panic(errors.New("id cannot be empty value"))
	}

	idAsInt, _ := strconv.Atoi(idData[0])
	err := database.DeleteTodo(db, idAsInt)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
