package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"fuse/message"
	"fuse/tppmessage"
)

func HandleCmdGetPurchasableAreaListRequest(ctx context.Context, message *message.Message, manager *SessionManager) error {
	var err error
	t := tppmessage.CmdGetPurchasableAreaListRequest{}
	err = json.Unmarshal(message.MData, &t)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	d := GetCmdGetPurchasableAreaListResponse(ctx, manager, message)

	message.MData, err = json.Marshal(d)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
