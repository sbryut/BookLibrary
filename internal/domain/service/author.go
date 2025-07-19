package service

import (
	"context"
	"library/internal/domain/entity"
)

type AuthorStorage interface {
	GetOne(id string) entity.Author
	GetAll(ctx context.Context) []entity.Author
	Create(author entity.Author) entity.Author
	Delete(author entity.Author) error
}

type authorService struct {
	storage AuthorStorage
}

func NewAuthorService(storage AuthorStorage) *authorService {
	return &authorService{storage: storage}
}

func (s authorService) GetByID(ctx context.Context, id string) entity.Author {
	return s.storage.GetOne(id)
}

func (s authorService) GetAll(ctx context.Context) []entity.Author {
	return s.storage.GetAll(ctx)
}

func (s authorService) Create(ctx context.Context) entity.Author {
	return entity.Author{} // dto - связующая часть слоя бизнес логики со слоем контроллеров
}
