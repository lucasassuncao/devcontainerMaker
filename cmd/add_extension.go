package cmd

import (
	"devcontainerMaker/internal/model"
	"devcontainerMaker/internal/repository"
	"devcontainerMaker/internal/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var extension string

var addExtensionCmd = &cobra.Command{
	Use:   "add-extensions",
	Short: "Add a VSCode extension to the DevContainer",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := repository.ReadDevContainerJSONFile()
		if err != nil {
			fmt.Printf("couldn't read the DevContainer JSON file: %v", err)
		}

		var dc model.DevContainer
		err = utils.JSONToStruct(file, &dc)
		if err != nil {
			fmt.Printf("couldn't unmarshall the DevContainer JSON: %v", err)
		}

		dc.AddExtension(extension)

		newFile, err := utils.PrettifyDevContainerJSON(&dc)
		if err != nil {
			fmt.Printf("couldn't format the DevContainer JSON: %v", err)
		}

		err = repository.SaveDevContainerJSONFile(newFile)
		if err != nil {
			fmt.Printf("couldn't write the DevContainer JSON file: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addExtensionCmd)
	addExtensionCmd.Flags().StringVarP(&extension, "extension", "e", "", "Extension ID (required)")
	addExtensionCmd.MarkFlagRequired("extension")
}
