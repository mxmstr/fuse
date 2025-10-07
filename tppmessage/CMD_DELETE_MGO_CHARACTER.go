package tppmessage

type CmdDeleteMgoCharacterRequest struct {
	Msgid          string `json:"msgid"`
	Rqid           int    `json:"rqid"`
	CharacterIndex int    `json:"characterIndex"`
}

type CmdDeleteMgoCharacterResponse struct {
	CryptoType string      `json:"crypto_type"`
	Flowid     interface{} `json:"flowid"`
	Msgid      string      `json:"msgid"`
	Result     string      `json:"result"`
	Rqid       int         `json:"rqid"`
	Xuid       interface{} `json:"xuid"`
}
