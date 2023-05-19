package main

import (
	"context"
	"errors"
	"fmt"
	"os/signal"

	"github.com/kelledge/devcontainers/golang"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background())
	defer cancel()

	fmt.Printf("starting work\n")

	err := golang.Run(ctx)
	if err != nil && !errors.Is(err, context.Canceled) {
		fmt.Printf("error: %s\n", err.Error())
	}

	fmt.Printf("exiting\n")
}
