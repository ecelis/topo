package main

import (
	"log"

	"github.com/ecelis/topo/cmd/topo/service"
)

func main() {
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
