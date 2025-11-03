package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
)

func HandleCmdSetMgoCharacterAndLoadout2Request(ctx context.Context, msg *message.Message, m *SessionManager) error {
	var req tppmessage.CmdSetMgoCharacterAndLoadout2Request
	err := json.Unmarshal(msg.MData, &req)
	if err != nil {
		return fmt.Errorf("could not unmarshal set mgo character and loadout request: %w", err)
	}

	err = m.MGOCharacterRepo.Upsert(ctx, msg.PlayerID, req.Character)
	if err != nil {
		return fmt.Errorf("could not upsert mgo characters: %w", err)
	}

	err = m.MGOLoadoutRepo.Upsert(ctx, msg.PlayerID, req.Loadout)
	if err != nil {
		return fmt.Errorf("could not upsert mgo loadouts: %w", err)
	}

	resp := tppmessage.CmdSetMgoCharacterAndLoadout2Response{
		Msgid:  tppmessage.CMD_SET_MGO_CHARACTER_AND_LOADOUT2.String(),
		Result: "NOERR",
	}

	msg.MData, err = json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("could not marshal set mgo character and loadout response: %w", err)
	}

	return nil
}
