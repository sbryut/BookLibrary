package composites

import (
	author_storage "library/internal/adapters/db/postgresdb"
	author_handler "library/internal/controller/http/v1"
	author_service "library/internal/domain/service"
	author_usecase "library/internal/domain/usecase/author"
	"net/http"
)

type AuthorHandler interface {
	Register(mux *http.ServeMux)
}

type BookServiceProvider interface {
	BookService() author_usecase.BookService
}

type GenreServiceProvider interface {
	GenreService() author_usecase.GenreService
}

type AuthorComposite struct {
	authorHandler AuthorHandler
}

func NewAuthorComposite(
	postgresDBComposite PostgresDBComposite,
	BookServiceProvider BookServiceProvider,
	GenreServiceProvider GenreServiceProvider,
) (*AuthorComposite, error) {
	bookService := BookServiceProvider.BookService()
	genreService := GenreServiceProvider.GenreService()

	authorStorage := author_storage.NewAuthorStorage(postgresDBComposite.DB())
	authorService := author_service.NewAuthorService(authorStorage)
	authorUseCase := author_usecase.NewAuthorUseCase(authorService, bookService, genreService)
	authorHandler := author_handler.NewAuthorHandler(authorUseCase)
	return &AuthorComposite{authorHandler}, nil
}
