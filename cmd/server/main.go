package main

import (
	"context"
	"fmt"
	"github.com/upper/db/v4/adapter/cockroachdb"
	"github.com/upper/db/v4/adapter/postgresql"
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

	//BD
	var settings = cockroachdb.ConnectionURL{
		Database: `postgres`,
		Host:     `127.0.0.1:5432`,
		User:     `postgres`,
		Password: `postgres`,
	}
	sess, err := postgresql.Open(settings)
	if err != nil {
		fmt.Printf("postgresql.Open:", err)
	}
	defer sess.Close()

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
	eventRepository := event.NewRepository(sess)
	eventService := event.NewService(&eventRepository)
	eventController := controllers.NewEventController(&eventService)

	// HTTP Server
	err2 := http.Server(
		ctx,
		http.Router(
			eventController,
		),
	)

	if err2 != nil {
		fmt.Printf("http server error: %s", err)
		exitCode = 2
		return
	}
}
