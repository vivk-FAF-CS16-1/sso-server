package main

import (
	"context"
	"cs-lab-6/internal/application"
	"cs-lab-6/internal/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	isDone := make(chan os.Signal)
	signal.Notify(isDone, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	config.Init()

	mainApp := application.NewApp(ctx)
	go mainApp.Start()

	<-isDone
	cancel()
	mainApp.Shutdown()
}
