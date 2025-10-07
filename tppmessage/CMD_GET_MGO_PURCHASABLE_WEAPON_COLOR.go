package tppmessage

type CmdGetMgoPurchasableWeaponColorRequest struct {
	Msgid    string `json:"msgid"`
	Rqid     int    `json:"rqid"`
	WeaponID int    `json:"weapon_id"`
}

type CmdGetMgoPurchasableWeaponColorResponse struct {
	CryptoType             string      `json:"crypto_type"`
	Flowid                 interface{} `json:"flowid"`
	Msgid                  string      `json:"msgid"`
	PurchasableWeaponColor struct {
		AlreadyReleased      int    `json:"already_released"`
		PurchasableColorList []struct {
			AlreadyPurchased int    `json:"already_purchased"`
			Color            uint32 `json:"color"`
			Level            int    `json:"level"`
			Point            int    `json:"point"`
			Prestige         int    `json:"prestige"`
			PurchaseType     int    `json:"purchase_type"`
		} `json:"purchasable_color_list"`
		ReleaseDate int `json:"release_date"`
		WeaponID    int `json:"weapon_id"`
	} `json:"purchasable_weapon_color"`
	Result string      `json:"result"`
	Rqid   int         `json:"rqid"`
	Xuid   interface{} `json:"xuid"`
}
