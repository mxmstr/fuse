package tppmessage

type CmdGetMgoParametersRequest struct {
	Msgid string `json:"msgid"`
	Rqid  int    `json:"rqid"`
}

type MgoParameter struct {
	ID    uint32 `json:"id"`
	Value int    `json:"value"`
}

type CmdGetMgoParametersResponse struct {
	MgoParameter []MgoParameter `json:"MgoParameter"`
	CryptoType   string         `json:"crypto_type"`
	Flowid       interface{}    `json:"flowid"`
	Msgid        string         `json:"msgid"`
	Result       string         `json:"result"`
	Rqid         int            `json:"rqid"`
	Xuid         interface{}    `json:"xuid"`
}
