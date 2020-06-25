package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arrrden/orchestrator/handlers"
)

// PORT variable
var PORT = "7070"

func main() {
	// Logger
	l := log.New(os.Stdout, "api", log.LstdFlags)

	// Handlers
	c := handlers.NewCatalogue(l)

	// ServeMux
	sm := http.NewServeMux()
	sm.Handle("/catalogue", c)

	// Serve struct
	s := &http.Server{
		Addr:         ":" + PORT, // set TCP address to listen on
		Handler:      sm,
		ErrorLog:     l,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	// Instantiate serve
	go func() {
		l.Println("The server is listening on port: 7070 🤩")
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// Create cancel conext defer shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	defer func() {
		err := s.Shutdown(ctx)
		if err != nil {
			l.Fatal(err)
		}
	}()

	// Channel blocks server from closing
	// sender
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)

	// receiver
	sig := <-sigChan
	l.Println("Signal received: shutting down", sig)
}
