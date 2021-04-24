package main

import (
	"log"
	"github.com/getlantern/systray"
	"os"
	"github.com/gongt/remote-shell/internal/icon"
	"github.com/gongt/remote-shell/internal/broadcaster"
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
	go startup()
	<-mQuit.ClickedCh
	onExit()
	os.Exit(0)
}

func onExit() {
	systray.Quit()
}

func startup() {
	broadcaster.Init()
	receiver.StartListener()
}
