package main

import (
	"net/http"
	"sysken-pay/models"
	"sysken-pay/repositories"
	"sysken-pay/types"
	"sysken-pay/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	//データベースとの接続
	db, err := utils.NewDBConnection()
	if err != nil {
		println("データベースの接続できませんでした")
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		println("Userのマイグレーションに失敗しました。")
	}

	err = db.AutoMigrate(&models.Item{})
	if err != nil {
		println("itemのマイグレーションに失敗しました。")
	}

	err = db.AutoMigrate(&models.Charge{})
	if err != nil {
		println("chargeのマイグレーションに失敗しました。")
	}

	err = db.AutoMigrate(&models.Purchase{})
	if err != nil {
		println("purchaseのマイグレーションに失敗しました。")
	}

	// Ginのルーティング
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Database",
		})
	})

	r.POST("/item", func(c *gin.Context) {
		var item models.Item
		var req types.CreateItemRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "リクエストの形式が正しくありません",
			})
			return
		}
		println("JANコード", req.JanCode, "名前", req.ItemName, "値段", req.Price) //デバック用
		item.Create(req.JanCode, req.ItemName, req.Price)
		err = repositories.CreateItem(db, &item)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "error",
				"message": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":     "success",
			"id":         item.ID,
			"jan_code":   item.JanCode,
			"item_name":  item.Name,
			"price":      item.Price,
			"created_at": item.CreatedAt,
			"updated_at": item.UpdatedAt,
		})
	})

	r.POST("/user", func(c *gin.Context) {
		var user models.User
		var req types.CreateUserRequest
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
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":     "success",
			"user_id":    user.ID,
			"user_name":  user.Name,
			"created_at": user.CreatedAt,
		})
	})

	r.POST("/user/:user_id/charge", func(c *gin.Context) {
		userId := c.Param("user_id")
		println("userId: ", userId)
		var charge models.Charge

		var req types.CreateChargeRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "リクエストの形式が正しくありません",
			})
			return
		}
		println("amount: ", req.Amount)
		// uuid.Parse() を使って文字列からUUID型に変換
		parsedUUID, err := uuid.Parse(userId)
		if err != nil {
			// パースに失敗した場合のエラー処理
			println("UUIDのパースに失敗しました: %v", err)
		}

		charge.Create(parsedUUID, req.Amount)

		err = repositories.CreateCharge(db, &charge)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "error",
				"message": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":        "success",
			"charge_id":     charge.ID,
			"charge_amount": charge.Amount,
			"user_id":       charge.UserId,
			"balance":       1000,
			"created_at":    charge.CreatedAt,
		})
	})

	r.GET("/item/:jan_code", func(c *gin.Context) {
		janCode := c.Param("jan_code")
		var item models.Item

		err = repositories.ReadItemByJanCode(db, janCode, &item)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "error",
				"message": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":    "success",
			"item_id":   item.ID,
			"item_name": item.Name,
			"price":     item.Price,
			"jan_code":  item.JanCode,
		})
	})

	r.POST("/purchase", func(c *gin.Context) {
		var purchase models.Purchase
		var req types.CreatePurchaseRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "リクエストの形式が正しくありません",
			})
			return
		}
		if err != nil {
			// パースに失敗した場合のエラー処理
			println("UUIDのパースに失敗しました: %v", err)
		}
		for i := range req.Items {
			println("Quantity", req.Items[i].Quantity)
			for range req.Items[i].Quantity {
				purchase.Create(req.UserID, req.Items[i].ItemID)
				println("itemID", req.Items[i].ItemID)
				err = repositories.CreatePurchase(db, &purchase)
				if err != nil {
					c.JSON(http.StatusOK, gin.H{
						"status":  "error",
						"message": err,
					})
					return
				}
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"status":      "success",
			"purchase_id": purchase.ID,
		})
	})

	// サーバー起動
	r.Run(":8080") // localhost:8080で待機
}
