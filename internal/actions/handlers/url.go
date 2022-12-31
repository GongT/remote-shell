package handlers

import (
	"log"
	"net/url"
	"os/exec"
	"strings"

	action_base "github.com/gongt/remote-shell/internal/actions/action-base"
)

type UrlAction struct {
	action_base.MessageBase
	Path string
}

func NewUrlAction(file string) action_base.Message {
	return &UrlAction{
		action_base.MessageBase{
			TypeId: action_base.TypeUrl,
			Id:     action_base.InvalidMessage,
		},
		file,
	}
}

func (act *UrlAction) Handle() (reply bool, err error) {
	p := strings.Replace(act.Path, "/", "\\", -1)
	p, err = url.PathUnescape(p)
	if err != nil {
		return true, err
	}

	cmd := exec.Command("powershell", "-Command", "Start-Process", p)

	log.Println("run command: ", cmd.Args)

	if err := cmd.Run(); err != nil {
		log.Println("Error:", err)
	}
	return true, nil
}
