package tppmessage

type CmdGetNextMaintenanceRequest struct {
	Msgid string `json:"msgid"`
	Rqid  int    `json:"rqid"`
}

type CmdGetNextMaintenanceResponse struct {
	Msgid           string      `json:"msgid"`
	Result          string      `json:"result"`
	CryptoType      string      `json:"crypto_type"`
	Flowid          interface{} `json:"flowid"`
	NextMaintenance int         `json:"next_maintenance"`
	MaintenanceType int         `json:"maintenance_type"`
	MessageType     int         `json:"message_type"`
	Rqid            int         `json:"rqid"`
	Xuid            interface{} `json:"xuid"`
}
