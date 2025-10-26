package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var SetMainDir = func(cmd *cobra.Command, args []string) {
	fmt.Println("this will set main dir!")

}
