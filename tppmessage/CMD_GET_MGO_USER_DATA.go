package tppmessage

type CmdGetMgoUserDataRequest struct {
	Msgid string `json:"msgid"`
	Rqid  int    `json:"rqid"`
}

type CmdGetMgoUserDataResponse struct {
	CryptoType            string      `json:"crypto_type"`
	Flowid                interface{} `json:"flowid"`
	Gp                    int         `json:"gp"`
	GpBoostMag            int         `json:"gp_boost_mag"`
	GpExpire              string      `json:"gp_expire"`
	GpExpireUnixTimestamp int         `json:"gp_expire_unix_timestamp"`
	Msgid                 string      `json:"msgid"`
	RankXp                int         `json:"rank_xp"`
	Result                string      `json:"result"`
	Reward                MGOReward   `json:"reward"`
	Rqid                  int         `json:"rqid"`
	SurvivalTicketRemain  int         `json:"survival_ticket_remain"`
	XpBoostMag            int         `json:"xp_boost_mag"`
	XpExpire              string      `json:"xp_expire"`
	XpExpireUnixTimestamp int         `json:"xp_expire_unix_timestamp"`
	Xuid                  interface{} `json:"xuid"`
}
