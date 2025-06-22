package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"fuse/message"
	"fuse/tppmessage"
)

func HandleCmdGetMBCoinRemainderRequest(ctx context.Context, msg *message.Message, manager *SessionManager) error {
	var err error
	t := tppmessage.CmdGetMBCoinRemainderRequest{}
	err = json.Unmarshal(msg.MData, &t)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	d := GetCmdGetMBCoinRemainderResponse(ctx, manager, msg)

	msg.MData, err = json.Marshal(d)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
