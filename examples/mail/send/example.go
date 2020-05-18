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

	message := salesfly.NewMessage("from@example.com", "Test", "This is just a test", "to@example.com")
	receipt, err := client.Mail.Send(message)
	if err != nil {
		log.Fatal("Failed to send mail")
	}
	fmt.Printf("Location: %+v\n", receipt)
}
