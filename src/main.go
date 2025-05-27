package main

import (
	"net/http"
	"sysken-pay/models"
	"sysken-pay/repositories"
	"sysken-pay/utils"

	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	UserName string `json:"user_name"`
}

func main() {

	db, err := utils.NewDBConnection()
	if err != nil {
		panic("データベースの接続できませんでした")
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic("Userのマイグレーションに失敗しました。")
	}

	// Ginのルーティング
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Database",
		})
	})
	r.POST("/user", func(c *gin.Context) {
		var user models.User
		var req CreateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "リクエストの形式が正しくありません",
			})
			return
		}
		println(req.UserName)
		user.Create(req.UserName)
		err = repositories.CreateUser(db, &user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "error",
				"message": err,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"status":     "success",
			"user_id":    user.ID,
			"user_name":  user.Name,
			"created_at": user.CreatedAt,
		})
	})

	// サーバー起動
	r.Run(":8080") // localhost:8080で待機
}
