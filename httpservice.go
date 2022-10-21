package main

import (
	"context"
	"log"
	"net/http"
)

type HTTPService struct {
	httpServer *http.Server
}

func NewHTTPService(address string) (*HTTPService, error) {
	if address == "" {
		address = "127.0.0.1:8080"
	}

	h := &HTTPService{}

	mux := http.NewServeMux()
	mux.HandleFunc("/event", h.Event)

	h.httpServer = &http.Server{
		Addr:    address,
		Handler: mux,
	}

	go func() {
		log.Printf("http service listening on %s\n", address)
		if err := h.httpServer.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	return h, nil
}

func (h *HTTPService) Event(w http.ResponseWriter, r *http.Request) {
	log.Print("new request")
}

func (h *HTTPService) Stop(ctx context.Context) error {
	log.Println("stopping HTTP Service")
	return h.httpServer.Shutdown(ctx)
}
