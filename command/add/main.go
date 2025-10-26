package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var gitAdd = &cobra.Command{
		Use:   "mycli",
		Short: "My CLI tool",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from CLI!")
		},
	}
	gitAdd.Execute()
}
