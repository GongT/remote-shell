package main

import (
	"fmt"
	"log"
	"os"

	"github.com/getlantern/systray"
	"github.com/gongt/remote-shell/internal/icon"
	"github.com/gongt/remote-shell/internal/networking"
	"github.com/gongt/remote-shell/internal/receiver"
)

func main() {
	log.Println("hello, server.")

	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon.IconData)
	systray.SetTitle("Remote Shell")
	mQuit := systray.AddMenuItem("Quit", "Quit the app")
	startup()
	<-mQuit.ClickedCh
	onExit()
	os.Exit(0)
}

func onExit() {
	systray.Quit()
}

func startup() {
	s, err := networking.CreateServer(receiver.MessageHandler)

	if err != nil {
		panic(fmt.Errorf("Can not create server: %w", err))
	}

	go s.Start()
}
