package postgresdb

import (
	"context"
	"database/sql"
	"library/internal/domain/entity"
)

type authorStorage struct { // реализует интерфейс
	db *sql.DB
}

func NewAuthorStorage(db *sql.DB) *authorStorage {
	return &authorStorage{db: db}
}

func (as authorStorage) GetOne(id string) entity.Author {
	return entity.Author{}
}

func (as authorStorage) GetAll(ctx context.Context) []entity.Author {
	return []entity.Author{}
}

func (as authorStorage) Create(author entity.Author) entity.Author {
	return entity.Author{}
}

func (as authorStorage) Delete(author entity.Author) error {
	return nil
}
