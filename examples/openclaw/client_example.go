package main

import (
	"crypto/tls"
	"fmt"
	"time"

	"github.com/open-nexa/nexa"
)

func runClient() {
	config := &nexa.ClientConfig{
		Addr:        "localhost:443",
		TLSConfig:   &tls.Config{InsecureSkipVerify: true},
		MaxIdleTime: 30 * time.Second,
	}

	client, err := nexa.NewClient(config)
	if err != nil {
		fmt.Printf("Failed to create client: %v\n", err)
		return
	}
	defer client.Close()

	record, err := client.Create(nexa.MessageType("task"), "example task content", map[string]interface{}{
		"priority": "high",
		"tags":     []string{"openclaw", "example"},
	})
	if err != nil {
		fmt.Printf("Failed to create record: %v\n", err)
		return
	}
	fmt.Printf("Created record: %+v\n", record)

	retrieved, err := client.Get(record.ID)
	if err != nil {
		fmt.Printf("Failed to get record: %v\n", err)
		return
	}
	fmt.Printf("Retrieved record: %+v\n", retrieved)

	updated, err := client.Update(record.ID, "updated content", map[string]interface{}{
		"status": "completed",
	})
	if err != nil {
		fmt.Printf("Failed to update record: %v\n", err)
		return
	}
	fmt.Printf("Updated record: %+v\n", updated)

	if err := client.Delete(record.ID); err != nil {
		fmt.Printf("Failed to delete record: %v\n", err)
		return
	}
	fmt.Println("Record deleted successfully")

	records, err := client.Search(nexa.MessageType("task"), map[string]interface{}{
		"status": "completed",
	})
	if err != nil {
		fmt.Printf("Failed to search records: %v\n", err)
		return
	}
	fmt.Printf("Found %d records\n", len(records))

	stream, err := client.OpenMediaStream(nexa.StreamType("events"))
	if err != nil {
		fmt.Printf("Failed to open media stream: %v\n", err)
		return
	}
	defer stream.Close()

	fmt.Fprintf(stream, "Hello from OpenClaw agent\n")
	fmt.Println("Stream message sent")

	fmt.Println("Example completed")
}
