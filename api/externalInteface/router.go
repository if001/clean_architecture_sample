package externalInteface

import (
	"github.com/gin-gonic/gin"
	"clean_architecture_sample/api/externalInteface/database"
	"clean_architecture_sample/api/gateway/controllers"
)

func Router() *gin.Engine {
	router := gin.Default()

	conn := database.NewSqlConnection()

	c := controllers.NewBookController(conn)

	router.GET("/books", c.GetAllBooks )

	return router
}