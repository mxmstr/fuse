package tppmessage

type CmdGetMgoCharacter2Request struct {
	Msgid string `json:"msgid"`
	Rqid  int    `json:"rqid"`
}

type CmdGetMgoCharacter2Response struct {
	Character   MGOCharacterData `json:"character"`
	CryptoType  string           `json:"crypto_type"`
	Flowid      interface{} `json:"flowid"`
	Msgid       string      `json:"msgid"`
	Result      string      `json:"result"`
	Rqid        int         `json:"rqid"`
	Xuid        interface{} `json:"xuid"`
}
