func main() {
  // Seed the random number generator
  rand.Seed(time.Now().UnixNano())

	// Define a slice of messages
	messages := []string{
    "blue",
    "red",
		"green",
		"yellow",
		"silver",
		"black",
		"purple",
		"white",
		"gold",
		"pink",
  }

	// Get the pplication name from APP_NAME env var
  appName := os.Getenv("APP_NAME")
  if appName == "" {
    appName = "color-message"
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
