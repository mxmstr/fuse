package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
)

func HandleCmdGetMgoProgressionRequest(ctx context.Context, msg *message.Message, m *SessionManager) error {
	var err error
	t := tppmessage.CmdGetMgoProgressionRequest{}
	err = json.Unmarshal(msg.MData, &t)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	resp := tppmessage.CmdGetMgoProgressionResponse{
		Msgid:      tppmessage.CMD_GET_MGO_PROGRESSION.String(),
		Result:     "NOERR",
		Rqid:       0,
		CryptoType: tppmessage.CRYPTO_TYPE_COMPOUND,
		Progression: tppmessage.MgoProgression{
			Version: 143737279559449,
			CharacterList: []tppmessage.MgoCharacterProgression{
				{Legendary: 1, Prestige: 2, Xp: 999999},
				{Legendary: 1, Prestige: 2, Xp: 999999},
				{Legendary: 1, Prestige: 2, Xp: 999999},
				{Legendary: 1, Prestige: 2, Xp: 999999},
			},
			PermanentUnlockList: []uint32{1146596596, 1689202044},
		},
	}

	msg.MData, err = json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
