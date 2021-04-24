package main

import (
	"github.com/gongt/remote-shell/internal/broadcaster"
	"github.com/gongt/remote-shell/internal/receiver"
	"log"
)

func main() {
	log.Println("hello, server.")

	broadcaster.Init()
	receiver.StartListener()
}
