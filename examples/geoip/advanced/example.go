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

	options := &salesfly.GeoLocationOptions{
		DisplayFields: "country_code", // Return only country code
	}

	result, err := client.GeoLocation.Get("8.8.8.8", options)
	if err != nil {
		log.Fatal("Failed to get IP location")
	}
	fmt.Printf("Location: %+v\n", result)
}
