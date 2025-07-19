package service

import (
	"context"
	"library/internal/domain/entity"
)

type GenreStorage interface {
	GetOne(id string) entity.Genre
	GetAll(ctx context.Context) []entity.Genre
	Create(genre entity.Genre) entity.Genre
	Delete(genre entity.Genre) error
}

type genreService struct {
	storage GenreStorage
}

func NewGenreService(storage GenreStorage) *genreService {
	return &genreService{storage: storage}
}

func (s genreService) GetByID(ctx context.Context, id string) entity.Genre {
	return s.storage.GetOne(id)
}

func (s genreService) GetAll(ctx context.Context) []entity.Genre {
	return s.storage.GetAll(ctx)
}

func (s genreService) Create(ctx context.Context) entity.Genre {
	return entity.Genre{} // dto - связующая часть слоя бизнес логики со слоем контроллеров
}
