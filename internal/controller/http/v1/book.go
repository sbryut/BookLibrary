package v1

import (
	"context"
	"encoding/json"
	"library/internal/controller/http/dto"
	"library/internal/domain/entity"
	"library/internal/domain/usecase/book"
	"net/http"
)

const (
	bookURL  = "/books/:book_id"
	booksURL = "/books"
)

type BookUseCase interface {
	CreateBook(ctx context.Context, dto book_usecase.CreateBookDTO) (string, error)
	ListAllBooks(ctx context.Context) []entity.BookView
	GetFullBook(ctx context.Context, id string) entity.FullBook
}

type bookHandler struct {
	bookUseCase BookUseCase
}

func NewBookHandler(bookUseCase BookUseCase) *bookHandler {
	return &bookHandler{bookUseCase: bookUseCase}
}

func (h *bookHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc(booksURL, h.GetAllBooks)
}

func (h *bookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	//books := h.bookService.GetAllBooks(context.Background(), 0, 0)
	//получаем из сервиса, потом unmarshall в байты и отдаем
	w.Write([]byte("books"))
	w.WriteHeader(http.StatusOK)
}

func (h *bookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var d dto.CreateAuthorDTO // это ожидаем в транспорте с  джсоном
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil { // получили данные
		return // error
	}
	// валидация

	// маппинг dto.CreateAuthorDTO --> book_usecase.CreateBookDTO
	usecaseDTO := book_usecase.CreateBookDTO{
		"",
		"",
		"",
		"",
		0,
	}
	book, err := h.bookUseCase.CreateBook(r.Context(), usecaseDTO)
	if err != nil {
		// JSON RPS: 200, error: {msg, ..., dev_msg} (not rest api)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(book))
}
