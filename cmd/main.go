package main

import (
	"context"
	"fmt"
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
	fmt.Println("app has been started")
	<-mainCtx.Done()

	stop()
	err = myApp.Stop(mainCtx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("app has been stopped")
}
