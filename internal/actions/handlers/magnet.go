package handlers

import (
	"github.com/gongt/remote-shell/internal/actions/action-base"
	"os/exec"
	"log"
	"net/url"
)

type MagnetAction struct {
	action_base.MessageBase
	Magnet string
}

func NewMagnetAction(magnet string) action_base.Message {
	return &MagnetAction{
		action_base.MessageBase{
			TypeId: action_base.TypeMagnet,
			Id:     action_base.InvalidMessage,
		},
		magnet,
	}
}

func (act *MagnetAction) Handle() (err error) {
	p := act.Magnet
	p, err = url.PathUnescape(p)
	if err != nil {
		return err
	}

	cmd := exec.Command("qbittorrent", p)

	log.Println("run command: ", cmd.Args)

	if err := cmd.Run(); err != nil {
		log.Println("Error:", err)
	}
	return nil
}
