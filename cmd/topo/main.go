package main

import (
	"fmt"
	"log"

	"github.com/ecelis/topo/cmd/topo/service"
)

func main() {
	if err := service.Run(); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}
