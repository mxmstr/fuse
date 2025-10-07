package tppmessage

type CmdGetMgoMissionInfoRequest struct {
	Msgid          string `json:"msgid"`
	Rqid           int    `json:"rqid"`
	MatchType      int    `json:"match_type"`
	RuleType       int    `json:"rule_type"`
	SurvivalParams struct {
		CurrentSurvivalWins int    `json:"current_survival_wins"`
		EarnedSurvivalGp    int    `json:"earned_survival_gp"`
		RewardCategory      string `json:"reward_category"`
		RewardIdA           int    `json:"reward_id_a"`
		RewardIdB           int    `json:"reward_id_b"`
		RewardIdC           int    `json:"reward_id_c"`
		SurvivalUpdateKey   int    `json:"survival_update_key"`
	} `json:"survival_params"`
}

type CmdGetMgoMissionInfoResponse struct {
	CryptoType    string      `json:"crypto_type"`
	Flowid        interface{} `json:"flowid"`
	GpBoostMag    int         `json:"gp_boost_mag"`
	Msgid         string      `json:"msgid"`
	RankParam     struct {
		CurrentRankXp int   `json:"current_rank_xp"`
		EarnedRankXp  int   `json:"earned_rank_xp"`
		RankXpList    []int `json:"rank_xp_list"`
	} `json:"rank_param"`
	Result        string `json:"result"`
	Rqid          int    `json:"rqid"`
	SurvivalParams struct {
		CurrentSurvivalWins int    `json:"current_survival_wins"`
		EarnedSurvivalGp    int    `json:"earned_survival_gp"`
		RewardCategory      string `json:"reward_category"`
		RewardIdA           int    `json:"reward_id_a"`
		RewardIdB           int    `json:"reward_id_b"`
		RewardIdC           int    `json:"reward_id_c"`
		SurvivalUpdateKey   int    `json:"survival_update_key"`
	} `json:"survival_params"`
	XpBoostMag int         `json:"xp_boost_mag"`
	Xuid       interface{} `json:"xuid"`
}
