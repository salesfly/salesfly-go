package main

import (
	"fmt"
	"log"
	"os"

	"github.com/salesfly/salesfly-go"
)

func main() {
	apiKey := os.Getenv("SALESFLY_APIKEY")
	client, err := salesfly.NewClient(apiKey, nil)

	result, err := client.GeoLocation.GetCurrent(nil)
	if err != nil {
		log.Fatal("Failed to get IP location")
	}
	fmt.Printf("Location: %+v\n", result)
}
