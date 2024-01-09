package main

import (
	"context"
	"github.com/ainurqa95/mood-lifter/internal/app"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	myApp, err := app.NewApp(mainCtx)
	if err != nil {
		log.Fatal(err)
	}
	err = myApp.Run(mainCtx)
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}

	<-mainCtx.Done()

	stop()
}
