package tppmessage

type CmdGetMgoLoadoutRequest struct {
	Msgid string `json:"msgid"`
	Rqid  int    `json:"rqid"`
}

type CmdGetMgoLoadoutResponse struct {
	Loadout    MGOLoadoutData `json:"loadout"`
	CryptoType string         `json:"crypto_type"`
	Flowid     interface{} `json:"flowid"`
	Msgid      string      `json:"msgid"`
	Result     string      `json:"result"`
	Rqid       int         `json:"rqid"`
	Xuid       interface{} `json:"xuid"`
}
