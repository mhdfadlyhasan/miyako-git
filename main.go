package main

import (
	commands "miyako-git/command/set-main-dir"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "repo-pr",
		Short: "A CLI to automate PRs across multiple repositories",
	}
	var setMainDir = &cobra.Command{
		Use:   "set",
		Short: "Set the main directory",
		Run:   commands.SetMainDir,
	}

	rootCmd.AddCommand(setMainDir)
	rootCmd.Execute()
}
