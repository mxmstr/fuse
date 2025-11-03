package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
)

func HandleCmdSetMgoStatRequest(ctx context.Context, msg *message.Message, m *SessionManager) error {
	var err error
	t := tppmessage.CmdSetMgoStatRequest{}
	err = json.Unmarshal(msg.MData, &t)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	resp := tppmessage.CmdSetMgoStatResponse{
		Msgid:  tppmessage.CMD_SET_MGO_STAT.String(),
		Result: "NOERR",
	}

	msg.MData, err = json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
