package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
)

func HandleCmdMgoDlcUpdateRequest(ctx context.Context, msg *message.Message, m *SessionManager) error {
	var req tppmessage.CmdMgoDlcUpdateRequest
	err := json.Unmarshal(msg.MData, &req)
	if err != nil {
		return fmt.Errorf("could not unmarshal mgo dlc update request: %w", err)
	}

	slog.Info("received dlc update request", slog.Any("request", req))

	resp := tppmessage.CmdMgoDlcUpdateResponse{
		Msgid:       tppmessage.CMD_MGO_DLC_UPDATE.String(),
		Result:      "NOERR",
		NowDlcFlags: 15,
		OldDlcFlags: 15,
	}

	msg.MData, err = json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("could not marshal mgo dlc update response: %w", err)
	}

	return nil
}
