package main

import (
	"ToDoTask/database"
	"ToDoTask/entities"
	customTemplates "ToDoTask/templates"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/valyala/quicktemplate/examples/basicserver/templates"
)

var db *gorm.DB
var err error

func main() {

	fmt.Println("starting the server at http://localhost:8080 ...")

	db, err = database.OpenDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)

	http.HandleFunc("/", displayMain)
	http.HandleFunc("/item-create", createItem)
	http.HandleFunc("/item-delete", deleteItem)

	if err = http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func displayMain(w http.ResponseWriter, r *http.Request) {
	db, err = database.ListTodoItems(db)
	if err != nil {
		fmt.Printf("Error listing todo items: %s", err.Error())
		panic(err)
	}

	p := &customTemplates.StartPage{
		R:    r,
		Data: database.ToDoItems,
	}

	templates.WritePageTemplate(w, p)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	p := &customTemplates.CreatePage{
		R: r,
	}

	templates.WritePageTemplate(w, p)

	err := entities.CreateNewItem(w, r, db)
	if err != nil {
		panic(err)
	}
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	p := &customTemplates.StartPage{
		R: r,
	}

	templates.WritePageTemplate(w, p)

	err := entities.DeleteItem(w, r, db)
	if err != nil {
		panic(err)
	}
}
