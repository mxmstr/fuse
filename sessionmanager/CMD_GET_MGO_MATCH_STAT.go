package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
)

func HandleCmdGetMgoMatchStatRequest(ctx context.Context, msg *message.Message, m *SessionManager) error {
	var err error
	t := tppmessage.CmdGetMgoMatchStatRequest{}
	err = json.Unmarshal(msg.MData, &t)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	resp := tppmessage.CmdGetMgoMatchStatResponse{
		Msgid:      tppmessage.CMD_GET_MGO_MATCH_STAT.String(),
		Result:     "NOERR",
		CryptoType: "COMPOUND",
		Flowid:     nil,
		Rqid:       0,
		Xuid:       nil,
		Abandon:    0,
		Played:     0,
		Started:    0,
	}

	msg.MData, err = json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
