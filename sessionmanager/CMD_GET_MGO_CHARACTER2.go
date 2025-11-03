package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
)

func HandleCmdGetMgoCharacter2Request(ctx context.Context, msg *message.Message, m *SessionManager) error {
	var req tppmessage.CmdGetMgoCharacter2Request
	err := json.Unmarshal(msg.MData, &req)
	if err != nil {
		return fmt.Errorf("could not unmarshal get mgo character request: %w", err)
	}

	chars, err := m.MGOCharacterRepo.FindAllByPlayer(ctx, msg.PlayerID)
	if err != nil {
		return fmt.Errorf("could not get mgo characters: %w", err)
	}

	resp := tppmessage.CmdGetMgoCharacter2Response{
		Msgid:     tppmessage.CMD_GET_MGO_CHARACTER2.String(),
		Result:    "NOERR",
		Character: chars,
	}

	msg.MData, err = json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("could not marshal get mgo character response: %w", err)
	}

	return nil
}
