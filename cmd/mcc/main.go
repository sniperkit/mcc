package main

import (
	"log"
	"os"

	// internal
	"github.com/sniperkit/snk.golang.mcc/pkg/default/controller"
)

func main() {
	if err := controller.Run(); err != nil {
		log.Panicln(err)
		os.Exit(1)
	}
}
