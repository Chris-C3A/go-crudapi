package home

import (
	"encoding/json"
	"fmt"
	"httpserver/src/db"
	"httpserver/src/models"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	SQLquery := "SELECT * FROM books"
	rows, err := db.DB.Query(SQLquery)

	if err != nil {
		json.NewEncoder(w).Encode(models.Msg{Text: "FAILED error"})
		fmt.Println(err)
		return
	}

	defer rows.Close()

	var books []models.Book

	for rows.Next() {
		// empty book object
		var book models.Book

		if err := rows.Scan(&book.ID, &book.Title, &book.Content); err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)
}
