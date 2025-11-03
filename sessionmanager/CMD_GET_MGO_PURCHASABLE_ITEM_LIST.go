package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
)

func HandleCmdGetMgoPurchasableItemListRequest(ctx context.Context, msg *message.Message, m *SessionManager) error {
	var err error
	t := tppmessage.CmdGetMgoPurchasableItemListRequest{}
	err = json.Unmarshal(msg.MData, &t)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	resp := tppmessage.CmdGetMgoPurchasableItemListResponse{
		Msgid:  tppmessage.CMD_GET_MGO_PURCHASABLE_ITEM_LIST.String(),
		Result: "NOERR",
		PurchasableItemList: struct {
			PurchasableItemList []tppmessage.MGOPurchasableItem `json:"purchasable_item_list"`
		}{
			PurchasableItemList: []tppmessage.MGOPurchasableItem{
				{Category: 0, Price: 800, PurchaseID: 3001001, PurchaseType: 0},
				{Category: 1, Price: 100, PurchaseID: 4115003, PurchaseType: 0},
			},
		},
		CryptoType: "COMPOUND",
		Flowid:     nil,
		Rqid:       0,
	}

	msg.MData, err = json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
