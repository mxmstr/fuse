package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
)

func HandleCmdGetMgoTitleUsrRequest(ctx context.Context, msg *message.Message, m *SessionManager) error {

	var err error
	t := tppmessage.CmdGetMgoTitleUsrRequest{}
	err = json.Unmarshal(msg.MData, &t)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	resp := tppmessage.CmdGetMgoTitleUsrResponse{
		Msgid: tppmessage.CMD_GET_MGO_TITLE_USR.String(),
		TitleList: []tppmessage.MgoTitleList{
			{Flag: 0, Gp: 0, ID: 90900},
			{Flag: 0, Gp: 0, ID: 11160},
			{Flag: 0, Gp: 0, ID: 11180},
		},
		Result: "NOERR",
	}

	msg.MData, err = json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
