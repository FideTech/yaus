package main

import (
	"flag"
	"log"

	"github.com/FideTech/yaus/config"
	"github.com/FideTech/yaus/core"
)

func main() {
	log.Printf("welcome to yet another url shortener (yaus) v%s\n", core.GetVersion())

	configFile := flag.String("configFile", "config.yaml", "Config File full path. Defaults to current folder")

	flag.Parse()

	if err := config.Load(*configFile); err != nil {
		panic(err)
	}

	log.Printf("base url of %s", config.Config.System.BaseURL)
	log.Printf("loaded %d error short links and %d info short links", len(config.Config.Hardcoded.Error), len(config.Config.Hardcoded.Info))
}
