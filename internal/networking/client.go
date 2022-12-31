package networking

import (
	"encoding/json"
	"log"
	"net"

	action_base "github.com/gongt/remote-shell/internal/actions/action-base"
	"github.com/gongt/remote-shell/internal/actions/handlers"
	"github.com/gongt/remote-shell/internal/constants"
	timeout_controller "github.com/gongt/remote-shell/internal/timeout-controller"
)

type client struct {
	conn net.Conn
}

var cli *client

func CreateClient() {
	c, err := net.Dial("tcp", constants.Server)
	if err != nil {
		log.Printf("client create failed %s\n", err)
	}

	cli = &client{
		conn: c,
	}
}

func Action(msg action_base.Message) error {
	return cli.Action(msg)
}

func (cli *client) Action(msg action_base.Message) (err error) {
	guid, timeout := timeout_controller.Wait(constants.Timeout)

	msg.SetId(guid)
	bytes, err := json.Marshal(msg)
	if err != nil {
		log.Println("failed to encode json:", err)
		return
	}

	log.Println("sending message: ", string(bytes))
	_, err = cli.conn.Write(bytes)
	if err != nil {
		log.Println("failed to sending:", err)
		return
	}

	switch msg.(type) {
	case *handlers.CallbackAction:
		return
	}

	err = <-timeout
	if err != nil {
		log.Println("not get any response:", err)
		// return
	}

	return
}
