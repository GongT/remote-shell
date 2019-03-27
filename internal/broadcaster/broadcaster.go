package broadcaster

import (
	"net"
	"github.com/gongt/remote-shell/internal/constants"
	"github.com/dmichael/go-multicast/multicast"
	"encoding/json"
	"log"
	"github.com/gongt/remote-shell/internal/timeout-controller"
	"github.com/gongt/remote-shell/internal/actions/action-base"
	"github.com/gongt/remote-shell/internal/actions/handlers"
)

type Broadcaster struct {
	conn *net.UDPConn
}

var bc *Broadcaster

func Init() {
	bc = createBroadcaster()
}

func createBroadcaster() (ret *Broadcaster) {
	address := constants.MulticastAddress
	log.Printf("Creating Broadcaster on %s\n", address)

	broadcaster, err := multicast.NewBroadcaster(address)
	if err != nil {
		log.Printf("broadcaster create failed %s\n", err)
	}

	ret = &Broadcaster{
		broadcaster,
	}

	return
}

func (bc *Broadcaster) Action(msg action_base.Message) (err error) {
	guid, timeout := timeout_controller.Wait(constants.Timeout)

	msg.SetId(guid)
	bytes, err := json.Marshal(msg)
	if err != nil {
		log.Println("failed to encode json:", err)
		return
	}

	log.Println("broadcasting message: ", string(bytes))
	_, err = bc.conn.Write(bytes)
	if err != nil {
		log.Println("failed to broadcast:", err)
		return
	}

	switch msg.(type) {
	case *handlers.CallbackAction:
		return
	}

	err = <-timeout
	if err != nil {
		log.Println("not get any response:", err)
		return
	}

	return
}

func Action(msg action_base.Message) error {
	return bc.Action(msg)
}
