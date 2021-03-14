package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/FideTech/yaus/config"
	"github.com/FideTech/yaus/core"
	"github.com/FideTech/yaus/router"
)

func main() {
	log.Printf("welcome to yet another url shortener (yaus) v%s\n", core.GetVersion())

	configFile := flag.String("configFile", "config.yaml", "Config File full path. Defaults to current folder")

	flag.Parse()

	if err := config.Load(*configFile); err != nil {
		panic(err)
	}

	if err := core.Start(); err != nil {
		panic(err)
	}

	sl, err := core.GetAllDynamicShortLinks()
	if err != nil {
		panic(err)
	}

	log.Printf("loaded %d error, %d info short links and %d dynamic short links", len(config.Config.Hardcoded.Error), len(config.Config.Hardcoded.Info), len(sl))

	go func() {
		if err := router.Start(); err != nil {
			log.Println("failed to start the router")
			panic(err)
		}
	}()

	waitForQuitSignal()
	log.Println("yaus received the signal to exit 👋")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := router.Shutdown(ctx); err != nil {
		log.Fatal("Router forced to shutdown:", err.Error())
	}

	if err := core.Shutdown(); err != nil {
		log.Fatal("failed to shutdown the core:", err.Error())
	}

	log.Println("goodbye. See you soon™️")
}

func waitForQuitSignal() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
}
