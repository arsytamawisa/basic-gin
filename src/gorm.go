package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {

	/* Config database */
	db, err := gorm.Open("mysql", "root:root@/sekolahkoding_gin")

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	/* Router */
	router := gin.Default()
	router.GET("items", items)
	router.Run()

}

func items(context *gin.Context) {
	context.JSON(200, gin.H{
		"code":    "200",
		"message": "items",
	})
}
