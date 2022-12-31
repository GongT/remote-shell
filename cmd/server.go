package main

import (
	"fmt"
	"log"

	"github.com/gongt/remote-shell/internal/networking"
	"github.com/gongt/remote-shell/internal/receiver"
)

func main() {
	log.Println("hello, server.")

	s, err := networking.CreateServer(receiver.MessageHandler)
	if err != nil {
		panic(fmt.Errorf("Can not create server: %w", err))
	}

	go s.Start()
}
