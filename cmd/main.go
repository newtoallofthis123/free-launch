package main

import (
	"fmt"
	"github.com/newtoallofthis123/free-launch/internal/launcher"
	"github.com/newtoallofthis123/free-launch/internal/models"
	"github.com/newtoallofthis123/free-launch/internal/picker"
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

	if err := models.EnsureModels(); err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching models: %v\n", err)
		os.Exit(1)
	}

	modelsList, err := models.LoadModels()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading models: %v\n", err)
		os.Exit(1)
	}

	var modelID string

	if len(os.Args) >= 3 && os.Args[2] != "-h" && os.Args[2] != "--help" {
		query := os.Args[2]
		m, err := models.FindModel(query, modelsList)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		modelID = m.ID
	} else {
		opted, err := picker.PickModel(modelsList)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		modelID = opted.ID
	}

	if err := launcher.LaunchClaude(modelID); err != nil {
		fmt.Fprintf(os.Stderr, "Error launching claude: %v\n", err)
		os.Exit(1)
	}
}
