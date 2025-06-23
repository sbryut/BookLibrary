package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Book struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		log.Printf("error decoding request body into Book struct %v", err)
		return
	}
	if err := db.QueryRow("INSERT INTO books (title, description, author) VALUES ($1, $2, $3)", book.Title, book.Description, book.Author).Err(); err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []Book

	rows, err := db.Query("SELECT id, title, description, author FROM books;")
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.Id, &book.Title, &book.Description, &book.Author); err != nil {
			http.Error(
				w,
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		}
		books = append(books, book)
	}

	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(books)
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		log.Printf("error marshalling books into json %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		log.Printf("error parsing %d from string to integer %v", id, err)
		return
	}

	var book Book

	if err := db.QueryRow("SELECT id, title, description, author FROM books WHERE id = $1", id).Scan(&book.Id, &book.Title, &book.Description, &book.Author); err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(book)
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		log.Printf("error parsing %d from string to integer %v", id, err)
		return
	}

	if err := db.QueryRow("DELETE	FROM books WHERE id = $1", id).Err(); err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
