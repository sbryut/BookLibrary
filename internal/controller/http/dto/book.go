package dto

type CreateBookDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	AuthorID    string `json:"author_id"`
	GenreID     string `json:"genre_id"`
	Year        int    `json:"year"`
}

type UpdateBookDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	AuthorID    string `json:"author_id"`
	Year        int    `json:"year"`
	Busy        bool   `json:"busy"`
	Owner       string `json:"owner"`
}
