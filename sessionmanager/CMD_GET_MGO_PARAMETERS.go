package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
)

func HandleCmdGetMgoParametersRequest(ctx context.Context, msg *message.Message, m *SessionManager) error {
	var err error
	t := tppmessage.CmdGetMgoParametersRequest{}
	err = json.Unmarshal(msg.MData, &t)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	resp := tppmessage.CmdGetMgoParametersResponse{
		Msgid: tppmessage.CMD_GET_MGO_PARAMETERS.String(),
		MgoParameter: []tppmessage.MgoParameter{
			{ID: 1328199047, Value: 1},
			{ID: 848079236, Value: 1000},
			{ID: 3779705668, Value: 5000},
			{ID: 1163418251, Value: 5000},
			{ID: 2465414369, Value: 64},
			{ID: 3329735861, Value: 160},
			{ID: 1316804683, Value: 900},
			{ID: 2973312482, Value: 450},
			{ID: 1840770005, Value: 1500},
			{ID: 810408023, Value: 1},
			{ID: 4012801301, Value: 1200},
			{ID: 1918707871, Value: 18000},
			{ID: 693879535, Value: 500},
			{ID: 1297488958, Value: 10000},
			{ID: 2599574229, Value: 1},
			{ID: 687947340, Value: 1000},
			{ID: 2341384761, Value: 200},
			{ID: 2029346817, Value: 5000},
			{ID: 3368912490, Value: 2000},
			{ID: 3376035593, Value: 8},
			{ID: 3342468677, Value: 4},
			{ID: 2261465657, Value: 70},
			{ID: 2008882968, Value: 120},
			{ID: 693362684, Value: 6},
		},
		Result:     "NOERR",
		CryptoType: "COMPOUND",
	}

	msg.MData, err = json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil

}
