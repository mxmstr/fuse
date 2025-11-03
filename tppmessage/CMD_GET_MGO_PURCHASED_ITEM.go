package tppmessage

type CmdGetMgoPurchasedItemRequest struct {
	Msgid      string `json:"msgid"`
	Rqid       int    `json:"rqid"`
	Category   int    `json:"category"`
	PurchaseID int    `json:"purchase_id"`
}

type MGOPurchasedItem struct {
	Category     int `json:"category"`
	Price        int `json:"price"`
	PurchaseID   int `json:"purchase_id"`
	PurchaseType int `json:"purchase_type"`
}

type MGOPurchasedItemData struct {
	PurchasedItemList []MGOPurchasedItem `json:"purchased_item_list"`
}

type CmdGetMgoPurchasedItemResponse struct {
	CryptoType          string               `json:"crypto_type"`
	Flowid              interface{}          `json:"flowid"`
	Msgid               string               `json:"msgid"`
	PurchasableItemList MGOPurchasedItemData `json:"purchasable_item_list"`
	Result              string               `json:"result"`
	Rqid                int                  `json:"rqid"`
	Xuid                interface{}          `json:"xuid"`
}
