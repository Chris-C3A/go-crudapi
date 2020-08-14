package book

import (
	"encoding/json"
	"fmt"
	"httpserver/src/db"
	"httpserver/src/models"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)

	// create book
	bookSql := "INSERT INTO books (title, content) VALUES ($1, $2)"
	_, err := db.DB.Exec(bookSql, book.Title, book.Content)

	if err != nil {
		json.NewEncoder(w).Encode(models.Msg{Text: "FAILED error"})
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(models.Msg{Text: "succesfully created"})
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	// get url params
	vars := mux.Vars(r)
	id := vars["id"]

	// Create an empty book and make the sql query (using $1 for the parameter)
	var book models.Book
	bookSql := "SELECT id, title, content FROM books WHERE id = $1"

	err := db.DB.QueryRow(bookSql, id).Scan(&book.ID, &book.Title, &book.Content)
	// rows := db.DB.QueryContext()
	if err != nil {
		json.NewEncoder(w).Encode(models.Msg{Text: "FAILED error"})
		fmt.Println(err)
		return
	}

	// return json
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// get url params
	vars := mux.Vars(r)
	id := vars["id"]

	// Create an empty book and make the sql query (using $1 for the parameter)
	bookSql := "DELETE FROM books WHERE id = $1"
	deleted, err := db.DB.Exec(bookSql, id)

	if err != nil {
		json.NewEncoder(w).Encode(models.Msg{Text: "FAILED error"})
		fmt.Println(err)
		return
	}

	rows_affected, _ := deleted.RowsAffected()
	if rows_affected == 0 {
		json.NewEncoder(w).Encode(models.Msg{Text: "No records has been deleted"})
		return
	}

	json.NewEncoder(w).Encode(models.Msg{Text: "successfully deleted"})
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// get url params
	vars := mux.Vars(r)
	id := vars["id"]

	var book models.Book

	// load json decode data into book object
	json.NewDecoder(r.Body).Decode(&book)

	// update book
	bookSql := "UPDATE books SET title = $1, content = $2 WHERE id = $3"
	updated, err := db.DB.Exec(bookSql, book.Title, book.Content, id)

	if err != nil {
		json.NewEncoder(w).Encode(models.Msg{Text: "FAILED error"})
		fmt.Println(err)
		return
	}

	rows_affected, _ := updated.RowsAffected()
	if rows_affected == 0 {
		json.NewEncoder(w).Encode(models.Msg{Text: "No records have been updated"})
		return
	}

	json.NewEncoder(w).Encode(models.Msg{Text: "successfully updated"})
}
