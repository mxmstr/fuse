package tppmessage

type CmdMgoMissionResultRequest struct {
	Msgid          string `json:"msgid"`
	Rqid           int    `json:"rqid"`
	CharacterIndex int    `json:"character_index"`
	Code           int    `json:"code"`
	EarnedGp       int    `json:"earned_gp"`
	EarnedXp       int    `json:"earned_xp"`
	RankParam      struct {
		EarnedRankXp int `json:"earned_rank_xp"`
	} `json:"rank_param"`
	SurvivalParams struct {
		EarnedSurvivalGp int `json:"earned_survival_gp"`
	} `json:"survival_params"`
	Ucd string `json:"ucd"`
}

type RankParam struct {
	CurrentRankXp int   `json:"current_rank_xp"`
	EarnedRankXp  int   `json:"earned_rank_xp"`
	RankXpList    []int `json:"rank_xp_list"`
}

type SurvivalParams struct {
	CurrentSurvivalWins int    `json:"current_survival_wins"`
	EarnedSurvivalGp    int    `json:"earned_survival_gp"`
	RewardCategory      string `json:"reward_category"`
	RewardIdA           int    `json:"reward_id_a"`
	RewardIdB           int    `json:"reward_id_b"`
	RewardIdC           int    `json:"reward_id_c"`
	SurvivalUpdateKey   int    `json:"survival_update_key"`
}

type CmdMgoMissionResultResponse struct {
	CharacterIndex int            `json:"character_index"`
	Code           int            `json:"code"`
	CryptoType     string         `json:"crypto_type"`
	CurrentGp      int            `json:"current_gp"`
	CurrentXp      int            `json:"current_xp"`
	EarnedGp       int            `json:"earned_gp"`
	EarnedXp       int            `json:"earned_xp"`
	Flowid         interface{}    `json:"flowid"`
	Msgid          string         `json:"msgid"`
	RankParam      RankParam      `json:"rank_param"`
	Result         string         `json:"result"`
	Rqid           int            `json:"rqid"`
	SurvivalParams SurvivalParams `json:"survival_params"`
	Ucd            string         `json:"ucd"`
	Xuid           interface{}    `json:"xuid"`
}
