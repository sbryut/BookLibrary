package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		log.Printf("cannot load config %v", err)
	}

	err = OpenDatabase(config.DB)
	if err != nil {
		log.Printf("error connection to postgresql database %v", err)
	}
	defer CloseDatabase()

	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("POST /books", createBook)
	mux.HandleFunc("GET /books", getAllBooks)
	mux.HandleFunc("GET /books/{id}", getBook)
	mux.HandleFunc("DELETE /books/{id}", deleteBook)

	fmt.Println("Server is listening on port", config.Port)
	http.ListenAndServe(":"+config.Port, mux)
}
