package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/htquangg/go-ethereum-poc/cmd/commands"
)

func Execute() {
	if err := run(); err != nil {
		fmt.Printf("store service exitted abnormally: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())

	signals := make(chan os.Signal, 1)
	signal.Notify(signals,
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGHUP,
	)
	defer func() {
		signal.Stop(signals)
		cancel()
	}()

	go func() {
		<-signals
		cancel()
		os.Exit(1)
	}()

	cmd := commands.NewRootCommand(ctx)

	return cmd.Execute()
}
