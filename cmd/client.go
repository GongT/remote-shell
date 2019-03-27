package main

import (
	"github.com/gongt/remote-shell/internal/broadcaster"
	"github.com/gongt/remote-shell/internal/actions/handlers"
	"os"
	"log"
	"github.com/gongt/remote-shell/internal/receiver"
	"strings"
	"github.com/gongt/remote-shell/internal/constants"
)

func main() {
	log.Println("hello, client.")

	if len(os.Args) == 1 {
		log.Println("Error: no arguments.")
		os.Exit(1)
	}

	broadcaster.Init()

	go receiver.StartCallbackListener()

	succ := false

	for i, f := range os.Args {
		if i == 0 {
			continue
		}
		if strings.HasPrefix(f, constants.FileOpenLocalPrefix) || strings.HasPrefix(f, "file://"+constants.FileOpenLocalPrefix) {
			f = strings.Replace(f, constants.FileOpenLocalPrefix, constants.FileOpenBase, 1)
			f = strings.Replace(f, "file://", "", 1)
			log.Println("open file:", f)
			succ = succ || open(f)
		} else if strings.HasPrefix(f, "http://") || strings.HasPrefix(f, "https://") {
			log.Println("open browser:", f)
			succ = succ || url(f)
		} else if strings.HasPrefix(f, "magnet:") {
			log.Println("open magnet:", f)
			succ = succ || magnet(f)
		} else {
			log.Println("Error: file not in valumes:", f)
		}
	}

	if !succ {
		os.Exit(1)
	}
}

func open(f string) bool {
	err := broadcaster.Action(handlers.NewOpenAction(f))
	if err != nil {
		log.Println("Error open file:", err)
		return false
	}
	return true
}

func url(f string) bool {
	err := broadcaster.Action(handlers.NewUrlAction(f))
	if err != nil {
		log.Println("Error open browser:", err)
		return false
	}
	return true
}

func magnet(f string) bool {
	err := broadcaster.Action(handlers.NewMagnetAction(f))
	if err != nil {
		log.Println("Error open magnet:", err)
		return false
	}
	return true
}
