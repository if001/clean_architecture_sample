package usecases

import "clean_architecture_sample/api/domain"

type BookRepository interface {
	FindAll() (*domain.Books, error)
	Delete() error
	Store() error
}
