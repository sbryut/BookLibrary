package postgresdb

import (
	"context"
	"database/sql"
	"library/internal/domain/entity"
)

type genreStorage struct { // реализует интерфейс
	db *sql.DB
}

func NewGenreStorage(db *sql.DB) *genreStorage {
	return &genreStorage{db: db}
}

func (as genreStorage) GetOne(id string) entity.Genre {
	return entity.Genre{}
}

func (as genreStorage) GetAll(ctx context.Context) []entity.Genre {
	return []entity.Genre{}
}

func (as genreStorage) Create(genre entity.Genre) entity.Genre {
	return entity.Genre{}
}

func (as genreStorage) Delete(genre entity.Genre) error {
	return nil
}
