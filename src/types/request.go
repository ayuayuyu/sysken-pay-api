package types

import "github.com/google/uuid"

// userエンドポイントのリクエスト
type CreateUserRequest struct {
	UserName string `json:"user_name"`
}

// itemエンドポイントのリクエスト
type CreateItemRequest struct {
	JanCode  string `json:"jan_code"`
	ItemName string `json:"item_name"`
	Price    int    `json:"price"`
}

type CreateChargeRequest struct {
	Amount int `json:"amount"`
}

// 購入する商品情報を表す構造体
type PurchaseItem struct {
	ItemID   int `json:"item_id"` // 商品IDを文字列型と仮定（必要に応じてintなどに変更）
	Quantity int `json:"quantity"`
}

// リクエストボディ全体の構造体
type CreatePurchaseRequest struct {
	UserID uuid.UUID      `json:"user_id"`
	Items  []PurchaseItem `json:"items"`
}
