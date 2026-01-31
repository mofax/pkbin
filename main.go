package main

import (
	"fmt"
	"os"

	"github.com/mofax/pkbin/internal/config"
	"github.com/mofax/pkbin/internal/runner"
)

func main() {
	// Check if script name was provided
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: pkbin <script-name>\n")
		os.Exit(1)
	}

	scriptName := os.Args[1]

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Find the script
	scriptCommand, err := cfg.FindScript(scriptName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Run the script
	exitCode, err := runner.RunScript(scriptCommand)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Exit with the script's exit code
	os.Exit(exitCode)
}
