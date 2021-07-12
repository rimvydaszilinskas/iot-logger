package main

import (
	"encoding/json"
	"fmt"

	"github.com/rimvydaszilinskas/announcer-backend/models"
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

	full := models.FullSystemDetails{
		System:           system,
		NetworkInterface: *networkInterface,
	}

	d, err := json.Marshal(full)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(d))
}
