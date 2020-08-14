package main

// basic CRUD api with golang

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"httpserver/src/api/book"
	"httpserver/src/api/home"
	"httpserver/src/db"
	"httpserver/src/middlewares"
)

func main() {
	// load env variables
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect to postgres db
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))

	db.DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.DB.Close()

	err = db.DB.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("succesfully connect to database")

	// mux router
	r := mux.NewRouter()

	// middlewares
	r.Use(middlewares.LoggingMiddleware)

	// Home route
	r.HandleFunc("/", home.Home).Methods("GET")

	// CRUD operations routes
	r.HandleFunc("/books/", book.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", book.GetBook).Methods("GET")
	r.HandleFunc("/books/{id}", book.DeleteBook).Methods("DELETE")
	r.HandleFunc("/books/{id}", book.UpdateBook).Methods("PUT")

	// start server
	fmt.Println("server listening at http://localhost:3000/")
	http.ListenAndServe(":3000", r)
}
