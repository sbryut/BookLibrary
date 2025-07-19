package author_usecase

import (
	"context"
	"library/internal/domain/entity"
)

type Service interface {
	GetByID(ctx context.Context, id string) entity.Author
}

type BookService interface {
	GetByID(ctx context.Context, id string) entity.Book
}

type GenreService interface {
	GetByID(ctx context.Context, id string) entity.Genre
}

type authorUseCase struct {
	authorService Service
	bookService   BookService
	genreService  GenreService
}

func NewAuthorUseCase(authorService Service, bookService BookService, genreService GenreService) *authorUseCase {
	return &authorUseCase{authorService: authorService, bookService: bookService, genreService: genreService}
}
