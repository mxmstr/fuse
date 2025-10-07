package tppmessage

type CmdGetMgoMatchStatRequest struct {
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

type CmdGetMgoMatchStatResponse struct {
	Abandon    int         `json:"abandon"`
	CryptoType string      `json:"crypto_type"`
	Flowid     interface{} `json:"flowid"`
	Msgid      string      `json:"msgid"`
	Played     int         `json:"played"`
	Result     string      `json:"result"`
	Rqid       int         `json:"rqid"`
	Started    int         `json:"started"`
	Xuid       interface{} `json:"xuid"`
}
