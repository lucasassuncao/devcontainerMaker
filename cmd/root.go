package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "devcontainer",
	Short: "Manage DevContainer configurations",
	Long:  "A CLI to create DevContainer configuration files",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
