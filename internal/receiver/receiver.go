package receiver

import (
	"net"
	"github.com/dmichael/go-multicast/multicast"
	"github.com/gongt/remote-shell/internal/actions"
	"github.com/gongt/remote-shell/internal/constants"
	"log"
	"github.com/gongt/remote-shell/internal/broadcaster"
	"github.com/gongt/remote-shell/internal/actions/handlers"
	"bytes"
)

func StartListener() {
	address := constants.MulticastAddress
	log.Printf("Listening on %s\n", address)

	multicast.Listen(address, msgHandler)
}

func StartCallbackListener() {
	address := constants.MulticastAddress
	log.Printf("Listening (only callback) on %s\n", address)

	multicast.Listen(address, cbHandler)
}

func msgHandler(src *net.UDPAddr, n int, b []byte) {
	b = bytes.Trim(b, "\x00")
	log.Println("receive: ", string(b))

	message, err := actions.Unmarshal(b)
	if err != nil {
		log.Println("failed to unmarshal:", err)
		return
	}

	err = message.Handle()
	if err != nil {
		log.Println("run action handler failed:", err)
		return
	}
	log.Println("run action handler complete. will response.")

	err = broadcaster.Action(handlers.NewCallbackAction(message.GetId()))
	if err != nil {
		log.Println("send callback failed:", err)
		return
	}
}

func cbHandler(src *net.UDPAddr, n int, b []byte) {
	b = bytes.Trim(b, "\x00")

	message, err := actions.Unmarshal(b)
	if err != nil {
		log.Println("failed to unmarshal:", err)
	}

	switch message.(type) {
	case *handlers.CallbackAction:
	default:
		return
	}

	log.Println("receive: ", string(b))

	err = message.Handle()
	if err != nil {
		log.Println("run action handler failed:", err)
	}
}
