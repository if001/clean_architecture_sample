package controllers

import (
	"fmt"
	"net/http"

	"clean_architecture_sample/api/gateway/repositories"
	"clean_architecture_sample/api/usecases"

	"github.com/gin-gonic/gin"
)

type bookController struct {
	UseCase usecases.BookUseCase
}

type BookController interface {
	GetAllBooks(c *gin.Context)
}

func (b *bookController) GetAllBooks(c *gin.Context) {
	fmt.Println()
	books, err := b.UseCase.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func NewBookController(dbConnection repositories.DBConnection) BookController {
	repo := repositories.NewBookRepository(dbConnection)
	u := usecases.NewBookUseCase(repo)
	return &bookController{UseCase: u}
}
