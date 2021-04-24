package handlers

import (
	"log"
	"net/url"
	"os"
	"os/exec"

	action_base "github.com/gongt/remote-shell/internal/actions/action-base"
	"github.com/gongt/remote-shell/internal/helpers"
)

type OpenAction struct {
	action_base.MessageBase
	Root string
	Path string
}

func NewOpenAction(root, file string) action_base.Message {
	return &OpenAction{
		action_base.MessageBase{
			TypeId: action_base.TypeOpen,
			Id:     action_base.InvalidMessage,
		},
		root,
		file,
	}
}

func (act *OpenAction) Handle() (err error) {
	letter := helpers.FindDriveById(act.Root)
	log.Printf("get samba drive: %s:", letter)
	if letter == "" {
		log.Println("ignore not exists:", act.Root)
		return
	}
	lpPath, err := url.PathUnescape(act.Path)
	if err != nil {
		return err
	}
	lpPath = letter + ":/" + lpPath

	cmd := exec.Command("rundll32.exe", "url.dll,FileProtocolHandler", lpPath)
	// cmd := exec.Command("powershell", "-windowstyle", "hidden", "-Command", "Start-Process", fmt.Sprintf("\"%s\"", lpPath))

	log.Println("run command: ", cmd.Args)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Println("Error:", err)
	}
	return nil
}
