package tppmessage

type CmdGetMgoPurchasableItemListRequest struct {
	Msgid      string `json:"msgid"`
	Rqid       int    `json:"rqid"`
	PurchaseID int    `json:"purchase_id"`
}

type CmdGetMgoPurchasableItemListResponse struct {
	CryptoType          string      `json:"crypto_type"`
	Flowid              interface{} `json:"flowid"`
	Msgid               string      `json:"msgid"`
	PurchasableItemList struct {
		PurchasableItemList []struct {
			Category     int `json:"category"`
			Price        int `json:"price"`
			PurchaseID   int `json:"purchase_id"`
			PurchaseType int `json:"purchase_type"`
		} `json:"purchasable_item_list"`
	} `json:"purchasable_item_list"`
	Result string      `json:"result"`
	Rqid   int         `json:"rqid"`
	Xuid   interface{} `json:"xuid"`
}
