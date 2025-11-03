package tppmessage

type CmdGetMgoTitleUsrRequest struct {
	Msgid  string `json:"msgid"`
	Rqid   int    `json:"rqid"`
	Target struct {
		NpID struct {
			Handler struct {
				Data string `json:"data"`
				Term int    `json:"term"`
			} `json:"handler"`
		} `json:"np_id"`
		PlayerID int `json:"player_id"`
		SteamID  int `json:"steam_id"`
		Xuid     int `json:"xuid"`
	} `json:"target"`
}

type MgoTitleList struct {
	Flag int `json:"flag"`
	Gp   int `json:"gp"`
	ID   int `json:"id"`
}

type CmdGetMgoTitleUsrResponse struct {
	CryptoType string         `json:"crypto_type"`
	Flowid     interface{}    `json:"flowid"`
	Msgid      string         `json:"msgid"`
	Result     string         `json:"result"`
	Rqid       int            `json:"rqid"`
	TitleList  []MgoTitleList `json:"title_list"`
	Xuid       interface{}    `json:"xuid"`
}
