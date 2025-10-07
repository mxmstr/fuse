package tppmessage

type CmdGetMgoBoostRequest struct {
	Msgid     string `json:"msgid"`
	Rqid      int    `json:"rqid"`
	BoostType int    `json:"boost_type"`
}

type CmdGetMgoBoostResponse struct {
	CryptoType string      `json:"crypto_type"`
	Flowid     interface{} `json:"flowid"`
	Msgid      string      `json:"msgid"`
	Result     string      `json:"result"`
	Rqid       int         `json:"rqid"`
	Xuid       interface{} `json:"xuid"`
}
