package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
)

func HandleCmdSetMgoCharacter2Request(ctx context.Context, msg *message.Message, m *SessionManager) error {
	var req tppmessage.CmdSetMgoCharacter2Request
	err := json.Unmarshal(msg.MData, &req)
	if err != nil {
		return fmt.Errorf("could not unmarshal set mgo character request: %w", err)
	}

	// for each character in character list, upsert into database
	// for _, character := range req.Character.CharacterList {
	// 	err = m.MGOCharacterRepo.Upsert(ctx, msg.PlayerID, character)
	// 	if err != nil {
	// 		return fmt.Errorf("could not upsert mgo character: %w", err)
	// 	}
	// }

	resp := tppmessage.CmdSetMgoCharacter2Response{
		Msgid:  tppmessage.CMD_SET_MGO_CHARACTER2.String(),
		Result: "NOERR",
	}

	msg.MData, err = json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("could not marshal set mgo character and loadout response: %w", err)
	}

	return nil
}
