package entity

import (
	"fmt"
)

type BookView struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	AuthorName  string `json:"author_name"`
	GenreName   string `json:"genre_name"`
	Year        int    `json:"year"`
	Busy        bool   `json:"busy"`
}

type FullBook struct {
	Book
	Author Author `json:"author"`
	Genre  Genre  `json:"genre"`
}

type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	AuthorID    string `json:"author_id"`
	GenreID     string `json:"genre_id"`
	Year        int    `json:"year"`
	Busy        bool   `json:"busy"`
	Owner       string `json:"owner"`
}

func (b *Book) Take(owner string) error {
	if b.Busy {
		return fmt.Errorf("book is busy")
	}

	b.Owner = owner
	b.Busy = true
	return nil
}
