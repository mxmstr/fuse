package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
)

func HandleCmdGetMgoGpRequest(ctx context.Context, msg *message.Message, m *SessionManager) error {
	err := error(nil) // playerStatus, err := m.PlayerStatusRepo.Get(ctx, msg.PlayerID)
	// if err != nil {
	// 	return fmt.Errorf("could not get player status: %w", err)
	// }

	resp := tppmessage.CmdGetMgoGpResponse{
		Msgid:  tppmessage.CMD_GET_MGO_GP.String(),
		Result: "NOERR",
		Gp:     999999,
	}

	msg.MData, err = json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("could not marshal get mgo gp response: %w", err)
	}

	return nil
}
