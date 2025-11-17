package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type StudentCreateDto struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Group struct {
	gorm.Model
	Name string `json:"name"`
}

func main() {

	var dsn = "host=localhost user=user password=Zz004005 dbname=gorm-gin-intro port=5432 sslmode=disable"
	var db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed")
	}

	db.AutoMigrate(&Student{}, &Group{})

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"massage": "PONG",
		})
	})

	router.POST("/students", func(ctx *gin.Context) {
		var student StudentCreateDto

		if err := ctx.ShouldBindJSON(&student); err != nil {
			ctx.JSON(400, gin.H{
				"err": "Ошибка",
			})
			return
		}

		// student = {Name: "Yusup", Age: 21}

		result := db.Create(&student)

		if result.Error != nil {
			ctx.JSON(500, gin.H{
				"err": "Ошибка сервера",
			})
			return
		}

		ctx.JSON(201, &student)

	})

	router.GET("/student/:id", func(ctx *gin.Context) {
		student := Student{}
		id := ctx.Param("id")

		if result := db.First(&student, id); result.Error != nil {
			ctx.JSON(400, gin.H{
				"err": "Такого студента нет",
			})
			return
		}

		ctx.JSON(200, student)

	})

	router.PATCH("/student/:id", func(ctx *gin.Context) {
		var student Student

		id := ctx.Param("id")

		if result := db.First(&student, id); result.Error != nil {
			ctx.JSON(400, gin.H{
				"error": "Такого студента нет1",
			})
			return
		}

		var newStudent StudentCreateDto

		if err := ctx.ShouldBindJSON(&newStudent); err != nil {
			ctx.JSON(400, gin.H{
				"err": "Ошибка",
			})
			return
		}

		if result := db.Model(&Student{}).Where("id = ?", id).Updates(newStudent); result.Error != nil {
			ctx.JSON(400, gin.H{
				"err": "Ошибка",
			})
		}

		ctx.JSON(200, student)
	})

	router.DELETE("/student/:id", func(ctx *gin.Context) {
		var student Student

		id := ctx.Param("id")

		if result := db.First(&student, id); result.Error != nil {
			ctx.JSON(400, gin.H{
				"error": "Такого студента нет1",
			})
			return
		}

		db.Delete(&student)
		ctx.JSON(200, student)

	})

	router.POST("/group", func(ctx *gin.Context) {
		var group Group

		if err := ctx.ShouldBindJSON(&group); err != nil {
			ctx.JSON(400, gin.H{
				"err": "Ошибка",
			})
			return
		}
		result := db.Create(&group)

		if result.Error != nil {
			ctx.JSON(500, gin.H{
				"err": "Ошибка сервера",
			})
		}

		ctx.JSON(200, group)
	})

	router.GET("/group/:id", func(ctx *gin.Context) {
		var group Group
		id := ctx.Param("id")

		if result := db.First(&group, id); result.Error != nil {
			ctx.JSON(400, gin.H{
				"err": "такой группы нету",
			})
			return
		}
		ctx.JSON(200, group)

	})

	router.PATCH("/group/:id", func(ctx *gin.Context) {
		var group Group
		id := ctx.Param("id")
		var newGroup Group

		if err := ctx.ShouldBindJSON(&newGroup); err != nil {
			ctx.JSON(400, gin.H{
				"err": "Ошибка",
			})
			return
		}

		if result := db.First(&group, id); result.Error != nil {
			ctx.JSON(400, gin.H{
				"err": "Ошибка",
			})
			return
		}

		if result := db.Model(&Group{}).Where("id = ?", id).Updates(&newGroup); result.Error != nil {
			ctx.JSON(400, gin.H{
				"err": "Ошибка",
			})
			return
		}

		ctx.JSON(200, group)

	})

	router.DELETE("group/:id", func(ctx *gin.Context) {
		var group Group
		id := ctx.Param("id")

		if result := db.First(&group, id); result.Error != nil {
			ctx.JSON(400, gin.H{
				"err": "Ошибка",
			})
			return
		}

		db.Delete(&group)
		ctx.JSON(200, group)
	})

	router.GET("/group", func(ctx *gin.Context) {
		var group []Group
		db.Find(&group)
		ctx.JSON(200, &group)
	})

	router.Run()

}
