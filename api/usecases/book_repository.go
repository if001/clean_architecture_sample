package usecases

import "clean_architecture_sample/api/domain"

type BookRepository interface {
	GetAll() (domain.Books, error)
	//Find() (domain.Book, error)
	//Create() (domain.Book, error)
	//Update() (domain.Book, error)
	//Delete() (domain.Book, error)
}
