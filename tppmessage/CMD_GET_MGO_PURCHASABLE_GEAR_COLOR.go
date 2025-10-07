package tppmessage

type CmdGetMgoPurchasableGearColorRequest struct {
	Msgid  string `json:"msgid"`
	Rqid   int    `json:"rqid"`
	GearID uint32 `json:"gear_id"`
}

type CmdGetMgoPurchasableGearColorResponse struct {
	CryptoType           string      `json:"crypto_type"`
	Flowid               interface{} `json:"flowid"`
	Msgid                string      `json:"msgid"`
	PurchasableGearColor struct {
		AlreadyReleased      int    `json:"already_released"`
		GearID               uint32 `json:"gear_id"`
		PurchasableColorList []struct {
			AlreadyPurchased int    `json:"already_purchased"`
			Color            uint32 `json:"color"`
			Level            int    `json:"level"`
			Point            int    `json:"point"`
			Prestige         int    `json:"prestige"`
			PurchaseType     int    `json:"purchase_type"`
		} `json:"purchasable_color_list"`
		ReleaseDate int `json:"release_date"`
	} `json:"purchasable_gear_color"`
	Result string      `json:"result"`
	Rqid   int         `json:"rqid"`
	Xuid   interface{} `json:"xuid"`
}
