package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
)

func HandleCmdGetMgoUserDataRequest(ctx context.Context, msg *message.Message, m *SessionManager) error {
	var err error
	t := tppmessage.CmdGetMgoUserDataRequest{}
	err = json.Unmarshal(msg.MData, &t)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	resp := tppmessage.CmdGetMgoUserDataResponse{
		Msgid:                 tppmessage.CMD_GET_MGO_USER_DATA.String(),
		Result:                "NOERR",
		CryptoType:            "COMPOUND",
		Flowid:                nil,
		Gp:                    999999,
		GpBoostMag:            0,
		GpExpire:              "NotImplement",
		GpExpireUnixTimestamp: 0,
		RankXp:                999999,
		Reward: tppmessage.MGOReward{
			RewardCategory: "MGO_REWARD_CATEGORY_GEAR",
			RewardIdA:      2589171374,
			RewardIdB:      0,
			RewardIdC:      0,
		},
		Rqid:                  0,
		SurvivalTicketRemain:  10,
		XpBoostMag:            0,
		XpExpire:              "NotImplement",
		XpExpireUnixTimestamp: 0,
		Xuid:                  nil,
	}

	msg.MData, err = json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("could not marshal get mgo user data response: %w", err)
	}

	return nil
}
