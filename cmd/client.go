package main

import (
	"log"
	"os"
	"strings"

	"github.com/gongt/remote-shell/internal/actions/handlers"
	"github.com/gongt/remote-shell/internal/broadcaster"
	"github.com/gongt/remote-shell/internal/helpers"
	"github.com/gongt/remote-shell/internal/receiver"
)

func main() {
	log.Println("hello, client.", os.Args)

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
		if strings.HasPrefix(f, "file:///") {
			f = f[7:]
		}
		if strings.HasPrefix(f, "/") {
			rootId, relFile, err := helpers.FindRoot(f)
			if err != nil {
				log.Printf("can not open %s: %v", f, err)
			}
			log.Println("open file:", rootId, relFile)
			succ = succ || open(rootId, relFile)
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

func open(root, f string) bool {
	err := broadcaster.Action(handlers.NewOpenAction(root, f))
	if err != nil {
		log.Println("error open file:", err)
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
