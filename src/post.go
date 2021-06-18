package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()

	router.POST("/article", addArticle)

	router.Run()
}

func addArticle(context *gin.Context) {

	title := context.PostForm("title")
	desc := context.PostForm("desc")

	context.JSON(200, gin.H{
		"code":  "200",
		"title": title,
		"desc":  desc,
	})
}
