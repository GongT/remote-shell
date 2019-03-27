package actions

import (
	"encoding/json"
	"github.com/gongt/remote-shell/internal/actions/handlers"
	"errors"
	"fmt"
	"github.com/gongt/remote-shell/internal/actions/action-base"
)

func Unmarshal(b []byte) (ret action_base.Message, err error) {
	var base action_base.MessageBase
	err = json.Unmarshal(b, &base)
	if err != nil {
		return
	}

	switch base.TypeId {
	case action_base.TypeCallback:
		ret = &handlers.CallbackAction{}
	case action_base.TypeOpen:
		ret = &handlers.OpenAction{}
	case action_base.TypeUrl:
		ret = &handlers.UrlAction{}
	case action_base.TypeMagnet:
		ret = &handlers.MagnetAction{}
	default:
		err = errors.New(fmt.Sprintf("message type is unknown: %d", base.TypeId))
	}

	json.Unmarshal(b, ret)

	return
}
