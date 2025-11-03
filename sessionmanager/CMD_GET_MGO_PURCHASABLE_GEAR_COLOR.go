package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
)

func HandleCmdGetMgoPurchasableGearColorRequest(ctx context.Context, msg *message.Message, m *SessionManager) error {
	var err error
	t := tppmessage.CmdGetMgoPurchasableGearColorRequest{}
	err = json.Unmarshal(msg.MData, &t)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	resp := tppmessage.CmdGetMgoPurchasableGearColorResponse{
		Msgid:  tppmessage.CMD_GET_MGO_PURCHASABLE_GEAR_COLOR.String(),
		Result: "NOERR",
		PurchasableGearColor: tppmessage.PurchasableGearColor{
			AlreadyReleased: 1,
			GearID:          t.GearID,
			PurchasableColorList: []tppmessage.PurchasableColor{
				{AlreadyPurchased: 1, Color: 378514155, Level: 0, Point: 50, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 373432697, Level: 0, Point: 50, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 3020208773, Level: 0, Point: 50, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 76167981, Level: 0, Point: 50, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 12356394, Level: 0, Point: 50, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 3652746413, Level: 0, Point: 50, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 3642541877, Level: 0, Point: 50, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 42204107, Level: 0, Point: 50, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 288703677, Level: 0, Point: 0, Prestige: 0, PurchaseType: 0},
				{AlreadyPurchased: 1, Color: 3753018657, Level: 0, Point: 50, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 3202192720, Level: 0, Point: 50, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 596127025, Level: 0, Point: 50, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 3075741992, Level: 0, Point: 50, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 2281905530, Level: 0, Point: 50, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 2358646918, Level: 0, Point: 50, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 1320004083, Level: 0, Point: 50, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 441935374, Level: 0, Point: 50, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 3102859567, Level: 0, Point: 50, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 1334550402, Level: 0, Point: 50, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 877211317, Level: 0, Point: 50, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 75594939, Level: 0, Point: 50, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 2510955615, Level: 0, Point: 50, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 1563954279, Level: 0, Point: 50, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 3609649821, Level: 0, Point: 50, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 2421028971, Level: 0, Point: 50, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 2377407379, Level: 0, Point: 5000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 274590504, Level: 0, Point: 300, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 3874100232, Level: 0, Point: 300, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 844354489, Level: 0, Point: 300, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 603682560, Level: 0, Point: 1000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 3358777140, Level: 0, Point: 300, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 1319182002, Level: 0, Point: 300, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 82930585, Level: 0, Point: 300, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 1253000828, Level: 0, Point: 300, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 1517145530, Level: 0, Point: 300, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 3308198696, Level: 0, Point: 500, Prestige: 0, PurchaseType: 3},
				{AlreadyPurchased: 1, Color: 1602343766, Level: 0, Point: 300, Prestige: 0, PurchaseType: 3},
				{AlreadyPurchased: 1, Color: 211193641, Level: 0, Point: 0, Prestige: 0, PurchaseType: 4},
				{AlreadyPurchased: 1, Color: 2981525399, Level: 0, Point: 0, Prestige: 0, PurchaseType: 4},
				{AlreadyPurchased: 1, Color: 956441261, Level: 0, Point: 0, Prestige: 0, PurchaseType: 4},
				{AlreadyPurchased: 1, Color: 1223860679, Level: 0, Point: 0, Prestige: 0, PurchaseType: 4},
				{AlreadyPurchased: 1, Color: 2072697952, Level: 0, Point: 0, Prestige: 0, PurchaseType: 4},
				{AlreadyPurchased: 1, Color: 1538492005, Level: 0, Point: 0, Prestige: 0, PurchaseType: 4},
				{AlreadyPurchased: 1, Color: 75500924, Level: 0, Point: 0, Prestige: 0, PurchaseType: 4},
				{AlreadyPurchased: 1, Color: 2775162649, Level: 0, Point: 0, Prestige: 0, PurchaseType: 4},
				{AlreadyPurchased: 1, Color: 1129643036, Level: 0, Point: 0, Prestige: 0, PurchaseType: 4},
				{AlreadyPurchased: 1, Color: 403463967, Level: 0, Point: 0, Prestige: 0, PurchaseType: 4},
				{AlreadyPurchased: 1, Color: 743569537, Level: 0, Point: 0, Prestige: 0, PurchaseType: 4},
				{AlreadyPurchased: 1, Color: 1483600195, Level: 0, Point: 0, Prestige: 0, PurchaseType: 4},
			},
			ReleaseDate: 0,
		},
	}

	msg.MData, err = json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("could not marshal get mgo purchasable gear color response: %w", err)
	}

	return nil
}
