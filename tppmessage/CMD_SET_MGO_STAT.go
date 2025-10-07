package tppmessage

type CmdSetMgoStatRequest struct {
	Msgid string  `json:"msgid"`
	Rqid  int     `json:"rqid"`
	Stat  MgoStat `json:"stat"`
}

type CmdSetMgoStatResponse struct {
	CryptoType string      `json:"crypto_type"`
	Flowid     interface{} `json:"flowid"`
	Msgid      string      `json:"msgid"`
	Result     string      `json:"result"`
	Rqid       int         `json:"rqid"`
	Xuid       interface{} `json:"xuid"`
}
