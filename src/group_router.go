package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()

	v1 := router.Group("api/v1/")
	{
		articles := v1.Group("/article")
		{
			articles.GET("/", article)
		}

		v1.GET("/users", users)
	}

	router.Run()
}

func article(context *gin.Context) {
	context.JSON(200, gin.H{
		"code":    "200",
		"message": "article",
	})
}

func users(context *gin.Context) {
	context.JSON(200, gin.H{
		"code":    "200",
		"message": "users",
	})
}
