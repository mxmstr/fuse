package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
)

func HandleCmdMgoMissionResultRequest(ctx context.Context, msg *message.Message, m *SessionManager) error {
	var err error
	t := tppmessage.CmdMgoMissionResultRequest{}
	err = json.Unmarshal(msg.MData, &t)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	resp := tppmessage.CmdMgoMissionResultResponse{
		Msgid:          tppmessage.CMD_MGO_MISSION_RESULT.String(),
		Result:         "NOERR",
		CryptoType:     "COMPOUND",
		Flowid:         nil,
		Xuid:           nil,
		CharacterIndex: 0,
		Code:           0,
		CurrentGp:      999999,
		CurrentXp:      999999,
		EarnedGp:       0,
		EarnedXp:       0,
		Rqid:           0,
		RankParam: tppmessage.RankParam{
			CurrentRankXp: 999999,
			EarnedRankXp:  0,
			RankXpList:    []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		SurvivalParams: tppmessage.SurvivalParams{
			CurrentSurvivalWins: 0,
			EarnedSurvivalGp:    0,
			RewardCategory:      "NotImplement",
			RewardIdA:           0,
			RewardIdB:           0,
			RewardIdC:           0,
			SurvivalUpdateKey:   0,
		},
		Ucd: "595034-0",
	}

	msg.MData, err = json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
