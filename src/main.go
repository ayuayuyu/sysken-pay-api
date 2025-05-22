package main

import (
	"log"
	"net/http"
	"sysken-pay/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func main() {

	type Test struct {
		gorm.Model
		Code  string
		Price uint
	}

	type Test1 struct {
		gorm.Model
		ID    uuid.UUID
		Code  string
		Price uint
	}

	db, err := utils.NewDBConnection()

	if err != nil {
		log.Fatal("error: ", err)
	}

	// Migrate the schema
	db.AutoMigrate(&Test{})
	db.AutoMigrate(&Test1{})

	// Ginのルーティング
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Database",
		})
	})
	r.GET("/test", func(c *gin.Context) {
		uid := uuid.Must(uuid.NewV7())
		// Create
		db.Create(&Test{Code: "D44", Price: 100})
		db.Create(&Test1{ID: uid, Code: "D45", Price: 100})
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	})

	// サーバー起動
	r.Run(":8080") // localhost:8080で待機
}
