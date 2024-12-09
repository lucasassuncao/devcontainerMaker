package cmd

import (
	"devcontainerMaker/internal/model"
	"devcontainerMaker/internal/repository"
	"devcontainerMaker/internal/utils"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var settingKey string
var settingValue string

var addSettingsCmd = &cobra.Command{
	Use:   "add-settings",
	Short: "Add a VSCode setting to the DevContainer",
	Run: func(cmd *cobra.Command, args []string) {
		// Try to parse the input value
		var parsedValue interface{}
		err := json.Unmarshal([]byte(settingValue), &parsedValue)
		if err != nil {
			// If parsing as JSON fails, fallback to treating it as a raw string
			parsedValue = settingValue
		}

		file, err := repository.ReadDevContainerJSONFile()
		if err != nil {
			fmt.Printf("couldn't read the DevContainer JSON file: %v", err)
		}

		var dc model.DevContainer
		err = utils.JSONToStruct(file, &dc)
		if err != nil {
			fmt.Printf("couldn't unmarshall the DevContainer JSON: %v", err)
		}

		dc.AddSetting(settingKey, parsedValue)

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
	rootCmd.AddCommand(addSettingsCmd)
	addSettingsCmd.Flags().StringVarP(&settingKey, "key", "k", "", "Setting key (required)")
	addSettingsCmd.Flags().StringVarP(&settingValue, "value", "v", "", "Setting value (required)")
	_ = addSettingsCmd.MarkFlagRequired("key")
	_ = addSettingsCmd.MarkFlagRequired("value")
}
