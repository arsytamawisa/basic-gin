package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

/* Set Global Variable */
var DB *gorm.DB
var err error

func main() {

	/* Config database */
	DB, err = gorm.Open("mysql", "root:root@/learn_sekolahkoding_gin")

	if err != nil {
		panic("failed to connect database")
	}

	defer DB.Close()

	/* Migrate */
	DB.AutoMigrate(&Article{})

	/* Router */
	router := gin.Default()
	router.GET("items", getAll)
	router.GET("items/:slug", getOne)
	router.POST("items", addItem)
	router.Run()
}

type Article struct {
	gorm.Model
	Title string
	Slug  string `gorm:"unique_index"`
	Desc  string `sql:"type:text"`
}

func addItem(context *gin.Context) {

	item := Article{
		Title: context.PostForm("title"),
		Slug:  context.PostForm("slug"),
		Desc:  context.PostForm("desc"),
	}

	/* Insert data to table article */
	DB.Create(&item)

	/* Return json */
	context.JSON(200, gin.H{
		"code": 200,
		"data": item,
	})
}

func getAll(context *gin.Context) {

	/* Create empty variable from struct article */
	items := []Article{}

	/* Retrive all data from table article */
	DB.Find(&items)

	/* Convert data to JSON */
	context.JSON(200, gin.H{
		"code": "200",
		"data": items,
	})
}

func getOne(context *gin.Context) {

	slug := context.Param("slug")
	var item Article

	/* Validata data if not exists */
	if DB.First(&item, "slug=?", slug).RecordNotFound() {
		context.JSON(404, gin.H{
			"code":    "404",
			"message": "Data not found",
		})

		/* Abort request if data not found */
		context.Abort()

		/* Return json */
		return
	}

	context.JSON(200, gin.H{
		"code": "200",
		"data": item,
	})
}
