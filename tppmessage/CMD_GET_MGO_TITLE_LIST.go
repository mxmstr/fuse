package tppmessage

type CmdGetMgoTitleListRequest struct {
	Msgid string `json:"msgid"`
	Rqid  int    `json:"rqid"`
}

type MGOTitle struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CmdGetMgoTitleListResponse struct {
	CryptoType string      `json:"crypto_type"`
	Flowid     interface{} `json:"flowid"`
	Msgid      string      `json:"msgid"`
	Result     string      `json:"result"`
	Rqid       int         `json:"rqid"`
	TitleList  []struct {
		Flag int `json:"flag"`
		Gp   int `json:"gp"`
		ID   int `json:"id"`
	} `json:"title_list"`
	Xuid interface{} `json:"xuid"`
}
