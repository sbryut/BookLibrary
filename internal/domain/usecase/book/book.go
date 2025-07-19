package book_usecase

import (
	"context"
	"library/internal/controller/http/dto"
	"library/internal/domain/entity"
)

type Service interface {
	GetByID(ctx context.Context, id string) entity.Book
	GetAllForList(ctx context.Context) []entity.BookView
}

type AuthorService interface {
	GetByID(ctx context.Context, id string) entity.Author
}

type GenreService interface {
	GetByID(ctx context.Context, id string) entity.Genre
}

type bookUseCase struct {
	bookService   Service
	authorService AuthorService
	genreService  GenreService
}

func NewBookUseCase(bookService Service, authorService AuthorService, genreService GenreService) *bookUseCase {
	return &bookUseCase{bookService: bookService, authorService: authorService, genreService: genreService}
}

func (u *bookUseCase) CreateBook(ctx context.Context, dto dto.CreateBookDTO) (string, error) {
	return "", nil
}

func (u *bookUseCase) ListAllBooks(ctx context.Context) []entity.BookView {
	//отображаем список книг с именем жанра и именем автора
	return u.bookService.GetAllForList(ctx)
}

func (u *bookUseCase) GetFullBook(ctx context.Context, id string) entity.FullBook {
	book := u.bookService.GetByID(ctx, id)
	genre := u.genreService.GetByID(ctx, book.GenreID)
	author := u.authorService.GetByID(ctx, book.AuthorID)
	return entity.FullBook{book, author, genre}
}
