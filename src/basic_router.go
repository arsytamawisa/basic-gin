package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()

	router.GET("/", getHome)

	router.Run()
}

func getHome(context *gin.Context) {
	context.JSON(200, gin.H{
		"code":    "200",
		"message": "welcome",
	})
}
