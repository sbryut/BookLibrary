package v1

import (
	"net/http"
)

const (
	authorURL  = "/authors/:author_id"
	authorsURL = "/authors"
)

type AuthorUseCase interface{}

type authorHandler struct {
	authorService AuthorUseCase
}

func NewAuthorHandler(service AuthorUseCase) *authorHandler {
	return &authorHandler{authorService: service}
}

func (h *authorHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc(authorsURL, h.GetAllAuthors)
}

func (h *authorHandler) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	//authors := h.bookService.GetAllBooks(context.Background(), 0, 0)
	//получаем из сервиса, потом unmarshall в байты и отдаем
	w.Write([]byte("authors"))
	w.WriteHeader(http.StatusOK)
}
