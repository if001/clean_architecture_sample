package usecases

import "clean_architecture_sample/api/domain"

type bookUseCase struct {
	BookRepo BookRepository
}
type BookUseCase interface {
	GetAllBooks() (*domain.Books, error)
}

func NewBookUseCase(repo BookRepository) BookUseCase {
	return &bookUseCase{BookRepo: repo}
}

func (b *bookUseCase) GetAllBooks() (*domain.Books, error) {
	books, err := b.BookRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return books, nil
}
