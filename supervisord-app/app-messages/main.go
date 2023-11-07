package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Define a slice of messages
	messages := []string{
		"Hello, world!",
		"How's it going?",
		"Another second passed!",
		"Keep smiling!",
		"Stay positive!",
		"Keep calm and code on!",
		"Believe in yourself!",
		"One moment at a time!",
		"Make every second count!",
		"Embrace the randomness!",
	}

	// Get the pplication name from APP_NAME env var
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "random-message"
	}

	// Infinite loop to print a message every second
	for {
		// Get a random index
		index := rand.Intn(len(messages))
		// Print a random message
		fmt.Printf("app=%s message=\"%s\"\n", appName, messages[index])
		// Wait for a second
		time.Sleep(1 * time.Second)
	}
}
