package tppmessage

type CmdGetMgoStatRequest struct {
	Msgid  string `json:"msgid"`
	Rqid   int    `json:"rqid"`
	Target struct {
		NpID struct {
			Handler struct {
				Data string `json:"data"`
				Term int    `json:"term"`
			} `json:"handler"`
		} `json:"np_id"`
		PlayerID int `json:"player_id"`
		SteamID  int `json:"steam_id"`
		Xuid     int `json:"xuid"`
	} `json:"target"`
}

type MgoStat struct {
	RuleStatList []struct {
		ID       uint32 `json:"id"`
		RuleCode uint32 `json:"rule_code"`
		Value    int    `json:"value"`
	} `json:"rule_stat_list"`
	StatList []struct {
		ID    uint32 `json:"id"`
		Value int    `json:"value"`
	} `json:"stat_list"`
}

type CmdGetMgoStatResponse struct {
	CryptoType string      `json:"crypto_type"`
	Flowid     interface{} `json:"flowid"`
	Msgid      string      `json:"msgid"`
	Result     string      `json:"result"`
	Rqid       int         `json:"rqid"`
	Stat       MgoStat     `json:"stat"`
	Xuid       interface{} `json:"xuid"`
}
