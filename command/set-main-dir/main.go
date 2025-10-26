package commands

import (
	"fmt"
	"miyako-git/helper"
	"os"

	"github.com/spf13/cobra"
)

var SetMainDir = func(cmd *cobra.Command, args []string) {
	fmt.Println("this will set main dir!")

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error!:", err)
		return
	}
	fmt.Println("Current working directory:", dir)
	helper.SaveConfig(helper.Config{
		MainDir: dir,
	})
}
