package main

import (
	"fmt"
	"os"

	// internal - plugins
	controller "github.com/sniperkit/snk.golang.mcc/plugin/service/client/xcc" // mcc + snk cli controller
)

func main() {
	if err := controller.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
