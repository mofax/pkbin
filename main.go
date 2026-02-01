package main

import (
	"fmt"
	"os"

	"github.com/mofax/pkbin/internal/config"
	"github.com/mofax/pkbin/internal/runner"
	"github.com/spf13/cobra"
)

var Version = "dev"

var rootCmd = &cobra.Command{
	Use:   "pk <script-name>",
	Short: "Run scripts defined in pkbin.jsonc",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		scriptName := args[0]

		cfg, err := config.LoadConfig()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		scriptCommand, err := cfg.FindScript(scriptName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		exitCode, err := runner.RunScript(scriptCommand)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		os.Exit(exitCode)
	},
}

func init() {
	rootCmd.Version = Version
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
