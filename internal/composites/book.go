package composites

import (
	book_storage "library/internal/adapters/db/postgresdb"
	book_handler "library/internal/controller/http/v1"
	book_service "library/internal/domain/service"
	book_usecase "library/internal/domain/usecase/book"
	"net/http"
))

type BookHandler interface {
	Register(mux *http.ServeMux)
}

type AuthorServiceProvider interface {
	BookService() book_usecase.AuthorService
}

type GenreServiceProvider interface {
	GenreService() book_usecase.GenreService
}
type BookComposite struct {
	BookHandler BookHandler
}

func NewBookComposite(postgresDBComposite PostgresDBComposite) (*BookComposite, error) {
	bookStorage := book_storage.NewBookStorage(postgresDBComposite.DB())
	bookService := book_service.NewBookService(bookStorage)
	bookUseCase := book_usecase.NewBookUseCase(bookService, authorService, genreService)
	bookHandler := book_handler.NewBookHandler(bookUseCase)
	return &BookComposite{bookHandler}, nil
}

