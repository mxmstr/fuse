package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
)

func HandleCmdGetNextMaintenanceRequest(ctx context.Context, msg *message.Message, m *SessionManager) error {
	var err error
	t := tppmessage.CmdGetNextMaintenanceRequest{}
	err = json.Unmarshal(msg.MData, &t)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	resp := tppmessage.CmdGetNextMaintenanceResponse{
		Msgid:           tppmessage.CMD_GET_NEXT_MAINTENANCE.String(),
		Result:          "NOERR",
		CryptoType:      "COMPOUND",
		Flowid:          nil,
		NextMaintenance: 1110,
		MaintenanceType: 0,
		MessageType:     0,
		Rqid:            0,
		Xuid:            nil,
	}

	msg.MData, err = json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("could not marshal get next maintenance response: %w", err)
	}

	return nil
}
