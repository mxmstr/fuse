package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
)

/*
data = {
	'crypto_type': 'COMPOUND',
	'flowid': None,
	'gp_boost_mag': 0,
	'msgid': 'CMD_GET_MGO_MISSION_INFO',
	'rank_param': {
		'current_rank_xp': 18,
		'earned_rank_xp': 0,
		'rank_xp_list': [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
	},
	'result': 'NOERR',
	'rqid': 0,
	'survival_params': {
		'current_survival_wins': 0,
		'earned_survival_gp': 0,
		'reward_category': 'NotImplement',
		'reward_id_a': 0,
		'reward_id_b': 0,
		'reward_id_c': 0,
		'survival_update_key': 0
	},
	'xp_boost_mag': 0,
	'xuid': None
}
*/

func HandleCmdGetMgoMissionInfoRequest(ctx context.Context, msg *message.Message, m *SessionManager) error {
	var err error
	t := tppmessage.CmdGetMgoMissionInfoRequest{}
	err = json.Unmarshal(msg.MData, &t)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	resp := tppmessage.CmdGetMgoMissionInfoResponse{
		Msgid:      tppmessage.CMD_GET_MGO_MISSION_INFO.String(),
		Result:     "NOERR",
		CryptoType: "COMPOUND",
		Flowid:     nil,
		Rqid:       0,
		Xuid:       nil,
		XpBoostMag: 0,
		GpBoostMag: 0,
	}
	resp.RankParam.CurrentRankXp = 18
	resp.RankParam.EarnedRankXp = 0
	resp.RankParam.RankXpList = make([]int, 16)
	resp.SurvivalParams.CurrentSurvivalWins = 0
	resp.SurvivalParams.EarnedSurvivalGp = 0
	resp.SurvivalParams.RewardCategory = "NotImplement"
	resp.SurvivalParams.RewardIdA = 0
	resp.SurvivalParams.RewardIdB = 0
	resp.SurvivalParams.RewardIdC = 0
	resp.SurvivalParams.SurvivalUpdateKey = 0

	msg.MData, err = json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
