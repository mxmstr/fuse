package tppmessage

type CmdGetMgoProgressionRequest struct {
	Msgid    string `json:"msgid"`
	PlayerID int    `json:"player_id"`
	Rqid     int    `json:"rqid"`
}

type CmdGetMgoProgressionResponse struct {
	CryptoType string      `json:"crypto_type"`
	Flowid     interface{} `json:"flowid"`
	Msgid      string      `json:"msgid"`
	Progression struct {
		CharacterList []struct {
			Legendary int `json:"legendary"`
			Prestige  int `json:"prestige"`
			Xp        int `json:"xp"`
		} `json:"character_list"`
		PermanentUnlockList []uint32 `json:"permanent_unlock_list"`
		Version             int64    `json:"version"`
	} `json:"progression"`
	Result string      `json:"result"`
	Rqid   int         `json:"rqid"`
	Xuid   interface{} `json:"xuid"`
}
