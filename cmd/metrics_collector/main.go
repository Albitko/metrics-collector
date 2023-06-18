package main

import (
	"log"

	"github.com/Albitko/metrics-collector/internal/app"
	"github.com/Albitko/metrics-collector/internal/config/contracts_settings"
)

func main() {

	contractsCfg, err := contracts_settings.New()
	if err != nil {
		log.Fatal(err)
	}

	app.Run(contractsCfg)
}
