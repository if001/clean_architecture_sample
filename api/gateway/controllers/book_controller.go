package controllers

import "clean_architecture_sample/api/usecases"

type BookController struct {
	bookUseCase usecases.BookUseCase
}

