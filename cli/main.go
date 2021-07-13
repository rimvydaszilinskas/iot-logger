package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/rimvydaszilinskas/iot-logger/models"
	"github.com/rimvydaszilinskas/iot-logger/network"
	"github.com/rimvydaszilinskas/iot-logger/system"
)

func main() {
	token := flag.String("token", "", "IOT device token")
	flag.Parse()

	if len(*token) == 0 {
		log.Fatalf("No token provided")
	}

	networkInterface, err := network.GetActiveNetworkInterface()

	if err != nil {
		log.Fatalf("error collecting network interface - %s\n", err)
	}

	system, err := system.CollectSystem()

	if err != nil {
		log.Fatalf("error collecting system - %s\n", err)
	}

	full := models.FullSystemDetails{
		System:           system,
		NetworkInterface: *networkInterface,
	}

	data, err := json.Marshal(full)

	if err != nil {
		log.Fatalf("error marshalling data - %s\n", err)
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/iot/", bytes.NewBuffer(data))

	if err != nil {
		log.Fatalf("error creating request - %s\n", err)
	}

	req.Header.Set("Authorization", *token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("error executing request - %s\n", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		log.Fatalf("retrieved a bad error code - %d", resp.StatusCode)
	}
	log.Println("OK")
}
