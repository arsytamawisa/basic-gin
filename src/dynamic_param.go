package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/article/:title", getArticle)

	router.Run()
}

func getArticle(context *gin.Context) {

	title := context.Param("title")

	context.JSON(200, gin.H{
		"code":    "200",
		"message": title,
	})
}
