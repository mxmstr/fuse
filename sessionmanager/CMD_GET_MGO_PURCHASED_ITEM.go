package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
)

func HandleCmdGetMgoPurchasedItemRequest(ctx context.Context, msg *message.Message, m *SessionManager) error {
	var err error
	t := tppmessage.CmdGetMgoPurchasedItemRequest{}
	err = json.Unmarshal(msg.MData, &t)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	resp := tppmessage.CmdGetMgoPurchasedItemResponse{
		Msgid:  tppmessage.CMD_GET_MGO_PURCHASED_ITEM.String(),
		Result: "NOERR",
		PurchasableItemList: tppmessage.MGOPurchasedItemData{
			PurchasedItemList: []tppmessage.MGOPurchasedItem{},
		},
	}

	msg.MData, err = json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
