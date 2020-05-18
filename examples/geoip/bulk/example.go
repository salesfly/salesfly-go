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

	result, err := client.GeoLocation.GetBulk([]string{"1.2.3.4", "8.8.8.8"}, nil)
	if err != nil {
		log.Fatal("Failed to get IP location")
	}

	for _, v := range result {
		fmt.Printf("Location: %+v\n", v)
	}
}
