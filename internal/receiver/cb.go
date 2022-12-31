package receiver

import (
	"encoding/json"
	"log"

	action_base "github.com/gongt/remote-shell/internal/actions/action-base"
	"github.com/gongt/remote-shell/internal/actions/handlers"
)

func makeCallback(msg action_base.Message) []byte {
	m := handlers.NewCallbackAction(msg.GetId())

	bytes, err := json.Marshal(m)
	if err != nil {
		log.Println("failed to encode json:", err)
		return nil
	}

	return bytes
}
