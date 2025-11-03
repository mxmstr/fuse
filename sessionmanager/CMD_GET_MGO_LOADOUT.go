package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
)

func HandleCmdGetMgoLoadoutRequest(ctx context.Context, msg *message.Message, m *SessionManager) error {
	var req tppmessage.CmdGetMgoLoadoutRequest
	err := json.Unmarshal(msg.MData, &req)
	if err != nil {
		return fmt.Errorf("could not unmarshal get mgo loadout request: %w", err)
	}

	loadout, err := m.MGOLoadoutRepo.FindAllByPlayer(ctx, msg.PlayerID)
	if err != nil {
		return fmt.Errorf("could not get mgo loadouts: %w", err)
	}

	resp := tppmessage.CmdGetMgoLoadoutResponse{
		Msgid:   tppmessage.CMD_GET_MGO_LOADOUT.String(),
		Loadout: loadout,
		Result:  "NOERR",
	}

	msg.MData, err = json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("could not marshal get mgo loadout response: %w", err)
	}

	return nil
}
