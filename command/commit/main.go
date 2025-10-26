package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var gitCommit = &cobra.Command{
		Use:   "mycli",
		Short: "My CLI tool",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from CLI!")
		},
	}
	gitCommit.Execute()
}
