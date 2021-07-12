package main

import (
	"fmt"

	"github.com/rimvydaszilinskas/announcer-backend/network"
	"github.com/rimvydaszilinskas/announcer-backend/system"
)

func main() {
	networkInterface, err := network.GetActiveNetworkInterface()

	if err != nil {
		panic(err)
	}

	system, err := system.CollectSystem()

	if err != nil {
		panic(err)
	}

	fmt.Println(networkInterface, system)

}
