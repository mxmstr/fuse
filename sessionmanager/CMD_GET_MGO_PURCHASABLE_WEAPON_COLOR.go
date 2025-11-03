package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
)

func HandleCmdGetMgoPurchasableWeaponColorRequest(ctx context.Context, msg *message.Message, m *SessionManager) error {
	var err error
	t := tppmessage.CmdGetMgoPurchasableWeaponColorRequest{}
	err = json.Unmarshal(msg.MData, &t)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	resp := tppmessage.CmdGetMgoPurchasableWeaponColorResponse{
		Msgid:      tppmessage.CMD_GET_MGO_PURCHASABLE_WEAPON_COLOR.String(),
		Result:     "NOERR",
		CryptoType: "COMPOUND",
		Flowid:     nil,
		Rqid:       0,
		Xuid:       nil,
		PurchasableWeaponColor: tppmessage.PurchasableWeaponColorData{
			AlreadyReleased: 1,
			PurchasableColorList: []tppmessage.PurchasableWeaponColor{
				{AlreadyPurchased: 1, Color: 2256514228, Level: 0, Point: 0, Prestige: 0, PurchaseType: 0},
				{AlreadyPurchased: 1, Color: 899472314, Level: 4, Point: 0, Prestige: 0, PurchaseType: 0},
				{AlreadyPurchased: 1, Color: 144194567, Level: 8, Point: 0, Prestige: 0, PurchaseType: 0},
				{AlreadyPurchased: 1, Color: 2715260476, Level: 12, Point: 0, Prestige: 0, PurchaseType: 0},
				{AlreadyPurchased: 1, Color: 1239413227, Level: 16, Point: 0, Prestige: 0, PurchaseType: 0},
				{AlreadyPurchased: 1, Color: 2727398566, Level: 20, Point: 0, Prestige: 0, PurchaseType: 0},
				{AlreadyPurchased: 1, Color: 3480584313, Level: 24, Point: 0, Prestige: 0, PurchaseType: 0},
				{AlreadyPurchased: 1, Color: 1918243418, Level: 28, Point: 0, Prestige: 0, PurchaseType: 0},
				{AlreadyPurchased: 1, Color: 3969456100, Level: 32, Point: 0, Prestige: 0, PurchaseType: 0},
				{AlreadyPurchased: 1, Color: 1715554328, Level: 36, Point: 0, Prestige: 0, PurchaseType: 0},
				{AlreadyPurchased: 1, Color: 593884744, Level: 0, Point: 300, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 255878003, Level: 0, Point: 300, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 1406043414, Level: 0, Point: 1000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 4056491133, Level: 0, Point: 1000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 3931447461, Level: 0, Point: 3000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 3694397221, Level: 0, Point: 300, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 2214881393, Level: 0, Point: 300, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 2479596555, Level: 0, Point: 1000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 2766761679, Level: 0, Point: 1000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 3770884251, Level: 0, Point: 3000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 3003022551, Level: 0, Point: 300, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 1791537266, Level: 0, Point: 300, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 2542877698, Level: 0, Point: 1000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 849237093, Level: 0, Point: 1000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 3312109921, Level: 0, Point: 3000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 1744411242, Level: 0, Point: 300, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 799374112, Level: 0, Point: 300, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 3134266807, Level: 0, Point: 1000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 554883529, Level: 0, Point: 1000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 1677907062, Level: 0, Point: 3000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 2728585102, Level: 0, Point: 100, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 2525769084, Level: 0, Point: 100, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 2313332439, Level: 0, Point: 100, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 3062087167, Level: 0, Point: 100, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 1415697796, Level: 0, Point: 100, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 1466343249, Level: 0, Point: 100, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 3946158833, Level: 0, Point: 100, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 1708178415, Level: 0, Point: 100, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 903507899, Level: 0, Point: 100, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 2662212364, Level: 0, Point: 100, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 3019361675, Level: 0, Point: 100, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 1190052616, Level: 0, Point: 100, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 1148551004, Level: 0, Point: 100, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 3414362956, Level: 0, Point: 100, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 3824204164, Level: 0, Point: 100, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 3191695622, Level: 0, Point: 10000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 907055393, Level: 0, Point: 2000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 4224686495, Level: 0, Point: 500, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 2176713901, Level: 0, Point: 500, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 2928993829, Level: 0, Point: 500, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 1811238617, Level: 0, Point: 500, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 2978906052, Level: 0, Point: 500, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 592797441, Level: 0, Point: 500, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 683280502, Level: 0, Point: 500, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 2988399529, Level: 0, Point: 500, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 384867490, Level: 0, Point: 1000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 4253086746, Level: 0, Point: 1000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 3230624264, Level: 0, Point: 1000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 3151572568, Level: 0, Point: 1000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 1848135364, Level: 0, Point: 1000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 1501478995, Level: 0, Point: 1000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 2465590794, Level: 0, Point: 1000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 4270413923, Level: 0, Point: 1000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 3934771567, Level: 0, Point: 1000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 3256609863, Level: 0, Point: 1000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 3978420825, Level: 0, Point: 1000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 663531477, Level: 0, Point: 1000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 970679455, Level: 0, Point: 1000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 110309555, Level: 0, Point: 1000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 1823855623, Level: 0, Point: 1000, Prestige: 0, PurchaseType: 2},
				{AlreadyPurchased: 1, Color: 2458510571, Level: 0, Point: 100, Prestige: 0, PurchaseType: 3},
				{AlreadyPurchased: 1, Color: 3698382022, Level: 0, Point: 100, Prestige: 0, PurchaseType: 3},
				{AlreadyPurchased: 1, Color: 2941984636, Level: 0, Point: 100, Prestige: 0, PurchaseType: 3},
				{AlreadyPurchased: 1, Color: 258590284, Level: 0, Point: 100, Prestige: 0, PurchaseType: 3},
				{AlreadyPurchased: 1, Color: 3416642560, Level: 0, Point: 100, Prestige: 0, PurchaseType: 3},
				{AlreadyPurchased: 1, Color: 3628633721, Level: 0, Point: 100, Prestige: 0, PurchaseType: 3},
				{AlreadyPurchased: 1, Color: 1104523036, Level: 0, Point: 100, Prestige: 0, PurchaseType: 3},
				{AlreadyPurchased: 1, Color: 1609172183, Level: 0, Point: 100, Prestige: 0, PurchaseType: 3},
				{AlreadyPurchased: 1, Color: 2451385763, Level: 0, Point: 100, Prestige: 0, PurchaseType: 3},
				{AlreadyPurchased: 1, Color: 2762514868, Level: 0, Point: 100, Prestige: 0, PurchaseType: 3},
				{AlreadyPurchased: 1, Color: 2349223133, Level: 0, Point: 0, Prestige: 0, PurchaseType: 1},
				{AlreadyPurchased: 1, Color: 3534932600, Level: 0, Point: 0, Prestige: 0, PurchaseType: 4},
				{AlreadyPurchased: 1, Color: 4056653872, Level: 0, Point: 0, Prestige: 0, PurchaseType: 4},
				{AlreadyPurchased: 1, Color: 819053298, Level: 0, Point: 0, Prestige: 0, PurchaseType: 4},
				{AlreadyPurchased: 1, Color: 3957873089, Level: 0, Point: 0, Prestige: 0, PurchaseType: 4},
				{AlreadyPurchased: 1, Color: 907591280, Level: 0, Point: 0, Prestige: 0, PurchaseType: 4},
				{AlreadyPurchased: 1, Color: 795656783, Level: 0, Point: 0, Prestige: 0, PurchaseType: 4},
				{AlreadyPurchased: 1, Color: 911714693, Level: 0, Point: 0, Prestige: 0, PurchaseType: 4},
				{AlreadyPurchased: 1, Color: 2693407561, Level: 0, Point: 0, Prestige: 0, PurchaseType: 4},
				{AlreadyPurchased: 1, Color: 1820872570, Level: 0, Point: 0, Prestige: 0, PurchaseType: 4},
				{AlreadyPurchased: 1, Color: 1669255690, Level: 0, Point: 0, Prestige: 0, PurchaseType: 4},
				{AlreadyPurchased: 1, Color: 1989783400, Level: 0, Point: 0, Prestige: 0, PurchaseType: 4},
				{AlreadyPurchased: 1, Color: 4156693392, Level: 0, Point: 0, Prestige: 0, PurchaseType: 4},
				{AlreadyPurchased: 1, Color: 3221919395, Level: 0, Point: 0, Prestige: 0, PurchaseType: 4},
			},
			ReleaseDate: 0,
			WeaponID:    582,
		},
	}

	msg.MData, err = json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
