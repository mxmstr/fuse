package tppmessage

type CmdGetMgoPurchasableGearRequest struct {
	Msgid      string `json:"msgid"`
	Rqid       int    `json:"rqid"`
	GearIDList struct {
		GearIDList []uint32 `json:"gear_id_list"`
	} `json:"gear_id_list"`
}

type CmdGetMgoPurchasableGearResponse struct {
	CryptoType          string      `json:"crypto_type"`
	Flowid              interface{} `json:"flowid"`
	Msgid               string      `json:"msgid"`
	PurchasableGearList struct {
		PurchasableGearList []struct {
			AlreadyPurchased int    `json:"already_purchased"`
			AlreadyReleased  int    `json:"already_released"`
			DefaultColor     uint32 `json:"default_color"`
			GearID           uint32 `json:"gear_id"`
			Point            int    `json:"point"`
			Prestige         int    `json:"prestige"`
			PurchaseType     int    `json:"purchase_type"`
		} `json:"purchasable_gear_list"`
	} `json:"purchasable_gear_list"`
	Result string      `json:"result"`
	Rqid   int         `json:"rqid"`
	Xuid   interface{} `json:"xuid"`
}
