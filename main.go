package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Fprintf(os.Stderr, "Usage: free-launch claude [model]\n")
		os.Exit(1)
	}

	if os.Args[1] != "claude" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\nUsage: free-launch claude [model]\n", os.Args[1])
		os.Exit(1)
	}

	if err := ensureModels(); err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching models: %v\n", err)
		os.Exit(1)
	}

	models, err := loadModels()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading models: %v\n", err)
		os.Exit(1)
	}

	var modelID string

	if len(os.Args) >= 3 && os.Args[2] != "-h" && os.Args[2] != "--help" {
		query := os.Args[2]
		m, err := findModel(query, models)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		modelID = m.ID
	} else {
		picked, err := pickModel(models)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		modelID = picked.ID
	}

	if err := launchClaude(modelID); err != nil {
		fmt.Fprintf(os.Stderr, "Error launching claude: %v\n", err)
		os.Exit(1)
	}
}
