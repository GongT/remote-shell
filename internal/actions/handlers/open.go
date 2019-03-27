package handlers

import (
	"github.com/gongt/remote-shell/internal/actions/action-base"
	"strings"
	"os/exec"
	"log"
	"net/url"
	"os"
	"fmt"
)

type OpenAction struct {
	action_base.MessageBase
	Path string
}

func NewOpenAction(file string) action_base.Message {
	return &OpenAction{
		action_base.MessageBase{
			TypeId: action_base.TypeOpen,
			Id:     action_base.InvalidMessage,
		},
		file,
	}
}

func (act *OpenAction) Handle() (err error) {
	lpPath := strings.Replace(act.Path, "/", "\\", -1)
	lpPath, err = url.PathUnescape(lpPath)
	if err != nil {
		return err
	}

	cmd := exec.Command("powershell", "-Command", "Start-Process", fmt.Sprintf("\"%s\"", lpPath))

	log.Println("run command: ", cmd.Args)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Println("Error:", err)
	}
	return nil
}
