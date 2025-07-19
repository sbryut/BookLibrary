package service

import (
	"context"
	"library/internal/domain/entity"
)

type BookStorage interface {
	GetOne(id string) entity.Book
	GetAll(ctx context.Context) []entity.Book
	Create(book entity.Book) entity.Book
	Delete(book entity.Book) error
}

type bookService struct {
	storage BookStorage
}

func NewBookService(storage BookStorage) *bookService {
	return &bookService{storage: storage}
}

func (s bookService) GetByID(ctx context.Context, id string) entity.Book {
	return s.storage.GetOne(id)
}

func (s bookService) GetAll(ctx context.Context) []entity.Book {
	return s.storage.GetAll(ctx)
}

func (s bookService) GetAllForList(ctx context.Context) []entity.BookView {
	// TODO implement
	return nil
}

func (s bookService) Create(ctx context.Context) entity.Book {
	return entity.Book{} // dto - связующая часть слоя бизнес логики со слоем контроллеров
}
