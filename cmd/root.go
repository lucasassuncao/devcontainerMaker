package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "devcontainer",
	Short: "Manage DevContainer configurations",
	Long:  "A CLI to add features, extensions, and settings to a DevContainer configuration.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
