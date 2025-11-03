package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
)

func HandleCmdDeleteMgoCharacterRequest(ctx context.Context, msg *message.Message, m *SessionManager) error {
	var req tppmessage.CmdDeleteMgoCharacterRequest
	err := json.Unmarshal(msg.MData, &req)
	if err != nil {
		return fmt.Errorf("could not unmarshal delete mgo character request: %w", err)
	}

	err = m.MGOCharacterRepo.Delete(ctx, msg.PlayerID, req.CharacterIndex)
	if err != nil {
		return fmt.Errorf("could not delete mgo character: %w", err)
	}

	resp := tppmessage.CmdDeleteMgoCharacterResponse{
		Msgid:  tppmessage.CMD_DELETE_MGO_CHARACTER.String(),
		Result: "NOERR",
	}

	msg.MData, err = json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("could not marshal delete mgo character response: %w", err)
	}

	return nil
}
