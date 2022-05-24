package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"server/internal/domein/event"
	"server/internal/infra/http"
	"server/internal/infra/http/controllers"
	"syscall"
)

func main() {
	exitCode := 0
	ctx, cancel := context.WithCancel(context.Background())

	// Recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("The system panicked!: %v\n", r)
			fmt.Printf("Stack trace form panic: %s\n", string(debug.Stack()))
			exitCode = 1
		}
		os.Exit(exitCode)
	}()

	// Signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-c
		fmt.Printf("Received signal '%s', stopping... \n", sig.String())
		cancel()
		fmt.Printf("Sent cancel to all threads...")
	}()

	// Event
	eventRepository := event.NewRepository()
	eventService := event.NewService(&eventRepository)
	eventController := controllers.NewEventController(&eventService)
	eventUserWeb := http.NewUseWeb(&eventService)

	// HTTP Server
	err := http.Server(
		ctx,
		http.Router(
			eventController, eventUserWeb,
		),
	)

	if err != nil {
		fmt.Printf("http server error: %s", err)
		exitCode = 2
		return
	}
}
