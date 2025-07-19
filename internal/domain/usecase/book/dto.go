package book_usecase

type CreateBookDTO struct {
	Title       string
	Description string
	AuthorID    string
	GenreID     string
	Year        int
}
