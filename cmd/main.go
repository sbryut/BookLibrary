package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"library/internal/config"
	"log"
	"net/http"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %v", err)
		return
	}

	log.Println("config initializing")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("cannot load config %v", err)
	}

	ctx := context.Background()

	log.Println("router initializing")
	mux := http.NewServeMux()

	log.Println("postgresdb composite initializing")
	postgresDBC, err := composites2.NewPostgresqlDBComposite(ctx, cfg)
	if err != nil {
		log.Fatalf("postgresqldb composite failed %v", err)
	}

	log.Println("author composite initializing")
	authorComposite, err := composites2.NewAuthorComposite(postgresDBC)
	if err != nil {
		log.Fatalf("author composite failed %v", err)
	}
	authorComposite.Handler.Register(mux)

	log.Println("book composite initializing")
	bookComposite, err := composites2.NewBookComposite(postgresDBC)
	if err != nil {
		log.Fatalf("book composite failed %v", err)
	}
	bookComposite.Handler.Register(mux)

	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("POST /books", createBook)
	mux.HandleFunc("GET /books", getAllBooks)
	mux.HandleFunc("GET /books/{id}", getBook)
	mux.HandleFunc("DELETE /books/{id}", deleteBook)

	fmt.Println("Server is listening on port", config.Port)
	http.ListenAndServe(":"+config.Port, mux)
}
