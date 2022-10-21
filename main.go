package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
)

func main() {
	eventServer, err := NewHTTPService("")
	if err != nil {
		log.Fatal(err)
	}

	raftService, err := NewRaftService("")
	if err != nil {
		log.Fatal(err)
	}

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	<-sigs

	log.Println("shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return eventServer.Stop(ctx)
	})
	g.Go(func() error {
		return raftService.Stop()
	})

	err = g.Wait()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("finished")
}
