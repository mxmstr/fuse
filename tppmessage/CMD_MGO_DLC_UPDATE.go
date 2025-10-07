package tppmessage

type CmdMgoDlcUpdateRequest struct {
	Msgid    string `json:"msgid"`
	Rqid     int    `json:"rqid"`
	PlayerID int    `json:"player_id"`
	DlcFlags int    `json:"dlc_flags"`
}

type CmdMgoDlcUpdateResponse struct {
	CryptoType  string      `json:"crypto_type"`
	Flowid      interface{} `json:"flowid"`
	Msgid       string      `json:"msgid"`
	NowDlcFlags int         `json:"now_dlc_flags"`
	OldDlcFlags int         `json:"old_dlc_flags"`
	Result      string      `json:"result"`
	Rqid        int         `json:"rqid"`
	Xuid        interface{} `json:"xuid"`
}
