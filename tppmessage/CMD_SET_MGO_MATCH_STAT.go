package tppmessage

type CmdSetMgoMatchStatRequest struct {
	Msgid   string `json:"msgid"`
	Rqid    int    `json:"rqid"`
	Abandon int    `json:"abandon"`
	Played  int    `json:"played"`
	Started int    `json:"started"`
}

type CmdSetMgoMatchStatResponse struct {
	CryptoType string      `json:"crypto_type"`
	Flowid     interface{} `json:"flowid"`
	Msgid      string      `json:"msgid"`
	Result     string      `json:"result"`
	Rqid       int         `json:"rqid"`
	Xuid       interface{} `json:"xuid"`
}
