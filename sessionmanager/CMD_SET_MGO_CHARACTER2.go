package sessionmanager

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/tppmessage"
)

func HandleCmdSetMgoCharacter2Request(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	t := tppmessage.CmdSetMgoCharacter2Request{}
	err = json.Unmarshal(message.MData, &t)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
