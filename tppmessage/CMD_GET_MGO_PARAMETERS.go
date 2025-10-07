package tppmessage

type CmdGetMgoParametersRequest struct {
	Msgid string `json:"msgid"`
	Rqid  int    `json:"rqid"`
}

type CmdGetMgoParametersResponse struct {
	MgoParameter []struct {
		ID    uint32 `json:"id"`
		Value int    `json:"value"`
	} `json:"MgoParameter"`
	CryptoType string      `json:"crypto_type"`
	Flowid     interface{} `json:"flowid"`
	Msgid      string      `json:"msgid"`
	Result     string      `json:"result"`
	Rqid       int         `json:"rqid"`
	Xuid       interface{} `json:"xuid"`
}
