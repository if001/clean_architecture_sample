package usecases

import "clean_architecture_sample/api/domain"

type BookUseCase struct {
	bookRepository BookRepository
}


func (u *BookUseCase) GetAll() (domain.Books, error) {
	return make(domain.Books,0), nil
}