package sessionmanager

import (
	"context"
	"fmt"
	"fuse/message"
	"fuse/tppmessage"
	"log/slog"
)

// nuke only
func GetCmdGetServerItemResponse(ctx context.Context, msg *message.Message, manager *SessionManager) (tppmessage.CmdGetServerItemResponse, error) {
	t := tppmessage.CmdGetServerItemResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_GET_RESOURCE_PARAM.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	// there is no server product id for nuke
	// server responds with your latest item development id and nuke development time
	// TODO separate record for nuke
	item, err := manager.ServerItemRepo.GetNukeTime(ctx, msg.PlayerID)
	if err != nil {
		t.Item = tppmessage.ServerItem{}

		t.Result = tppmessage.RESULT_NOERR
		return t, fmt.Errorf("using dummy nuke info, error: %w", err)
	}

	t.Item = tppmessage.ServerItem{
		CreateDate: item.CreateDate,
		Develop:    item.Develop,
		Gmp:        item.Gmp,
		Id:         item.ProductID,
		MaxSecond:  item.MaxSecond,
		MbCoin:     item.MbCoin,
		Open:       item.Open,
	}

	return t, nil
}

func HandleCmdGetServerItemResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	//var err error
	//t := GetCmdGetServerItemResponse()

	//message.MData, err = json.Marshal(t)
	//if err != nil {
	//	return fmt.Errorf("cannot marshal: %w", err)
	//}

	return nil
}
