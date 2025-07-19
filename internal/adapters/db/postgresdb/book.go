package postgresdb

import (
	"context"
	"database/sql"
	"library/internal/domain/entity"
)

// db -это слой доступа к данным
type bookStorage struct { // реализует интерфейс Storage в domain/book/author.go
	db *sql.DB
}

func NewBookStorage(db *sql.DB) *bookStorage {
	return &bookStorage{db: db}
}

func (bs bookStorage) GetOne(id string) entity.Book {
	return entity.Book{}
}

func (bs bookStorage) GetAll(ctx context.Context) []entity.Book {
	return []entity.Book{}
}

func (bs bookStorage) Create(book entity.Book) entity.Book {
	return entity.Book{}
}

func (bs bookStorage) Delete(book entity.Book) error {
	return nil
}
