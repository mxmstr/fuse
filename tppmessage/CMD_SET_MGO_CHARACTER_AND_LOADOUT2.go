package tppmessage

type CmdSetMgoCharacterAndLoadout2Request struct {
	Msgid string `json:"msgid"`
	Rqid  int    `json:"rqid"`
	// Character's structure is defined in an external JSON file (default_character.json)
	// which is not available in the repository. Using interface{} to accommodate any valid JSON.
	Character interface{} `json:"character"`
	// Loadout's structure is defined in an external JSON file (default_loadout.json)
	// which is not available in the repository. Using interface{} to accommodate any valid JSON.
	Loadout interface{} `json:"loadout"`
}

type CmdSetMgoCharacterAndLoadout2Response struct {
	CryptoType string      `json:"crypto_type"`
	Flowid     interface{} `json:"flowid"`
	Msgid      string      `json:"msgid"`
	Result     string      `json:"result"`
	Rqid       int         `json:"rqid"`
	Xuid       interface{} `json:"xuid"`
}
