package tppmessage

type CmdGetMgoCharacter2Request struct {
	Msgid string `json:"msgid"`
	Rqid  int    `json:"rqid"`
}

type CmdGetMgoCharacter2Response struct {
	// Character's structure is defined in an external JSON file (default_character.json)
	// which is not available in the repository. Using interface{} to accommodate any valid JSON.
	Character   interface{} `json:"character"`
	CryptoType  string      `json:"crypto_type"`
	Flowid      interface{} `json:"flowid"`
	Msgid       string      `json:"msgid"`
	Result      string      `json:"result"`
	Rqid        int         `json:"rqid"`
	Xuid        interface{} `json:"xuid"`
}
