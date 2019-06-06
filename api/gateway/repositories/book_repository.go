package repositories

import (
	"clean_architecture_sample/api/domain"
	"clean_architecture_sample/api/usecases"
)

type BookRepository struct {
	Connection DBConnection
}

func NewBookRepository(conn DBConnection) usecases.BookRepository {
	return &BookRepository{Connection: conn}
}

func (b *BookRepository) FindAll() (*domain.Books, error) {
	return nil, nil
}

func (b *BookRepository) Delete() error {
	return nil
}

func (b *BookRepository) Store() error {
	return nil
}
