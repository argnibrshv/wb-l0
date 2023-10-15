package main

import (
	"context"
	"os"
	"os/signal"
	"sub/internals/app"
	"sub/internals/cfg"

	log "github.com/sirupsen/logrus"
)

func main() {
	config := cfg.LoadAndStoreConfig()

	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	server := app.NewServer(ctx, config)

	go func() {
		oscall := <-c
		log.Printf("system call: %+v\n", oscall)
		server.ShutDown()
		cancel()
	}()

	server.Serve()
}
