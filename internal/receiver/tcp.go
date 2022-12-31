package receiver

import (
	"bytes"
	"log"
	"net"

	"github.com/gongt/remote-shell/internal/actions"
)

func MessageHandler(conn net.Conn, b []byte) {

	b = bytes.Trim(b, "\x00")
	log.Println("receive: ", string(b))

	message, err := actions.Unmarshal(b)
	if err != nil {
		log.Println("failed to unmarshal:", err)
		return
	}

	reply, err := message.Handle()
	if err != nil {
		log.Println("run action handler failed:", err)
		return
	}

	if !reply {
		return
	}

	log.Println("run action handler complete. will response.")

	bs := makeCallback(message)
	_, err = conn.Write(bs)
	if err != nil {
		log.Println("send callback failed:", err)
		return
	}
}
