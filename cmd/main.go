package main

import (
	"context"
	"github.com/ainurqa95/mood-lifter/internal/bots"
	"github.com/ainurqa95/mood-lifter/internal/config"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	cfg := config.NewConfig()

	bot, err := bots.DefineBot(cfg)
	if err != nil {
		log.Fatal(err)
	}
	//go func() {
	go func() {
		bot.Start(mainCtx)
	}()

	<-mainCtx.Done()

	stop()
}
