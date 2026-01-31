package runner

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// RunScript executes a script command via shell and returns the exit code.
// It streams stdout and stderr to the terminal in real-time and passes through
// all environment variables. Returns the exit code (0 on success, non-zero on failure).
func RunScript(command string) (int, error) {
	if command == "" {
		return 1, fmt.Errorf("command cannot be empty")
	}

	// Determine shell based on OS
	var shell string
	var shellArg string
	if runtime.GOOS == "windows" {
		shell = "cmd"
		shellArg = "/c"
	} else {
		shell = "/bin/sh"
		shellArg = "-c"
	}

	// Create command with shell
	cmd := exec.Command(shell, shellArg, command)

	// Set working directory to current directory
	cmd.Dir, _ = os.Getwd()

	// Pass through all environment variables
	cmd.Env = os.Environ()

	// Stream stdout and stderr to terminal in real-time
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Execute command
	if err := cmd.Run(); err != nil {
		// Check if it's an exit error to extract exit code
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode(), nil
		}
		// If it's not an exit error, return error with exit code 1
		return 1, fmt.Errorf("failed to execute command: %w", err)
	}

	// Success
	return 0, nil
}
