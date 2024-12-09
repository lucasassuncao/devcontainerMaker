package cmd

import (
	"devcontainerMaker/internal/model"
	"devcontainerMaker/internal/repository"
	"devcontainerMaker/internal/utils"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var featureKey string
var featureValues []string

var addFeaturesCmd = &cobra.Command{
	Use:   "add-features",
	Short: "Add a feature to the DevContainer",
	Long: `Add a feature to the DevContainer. 
You can provide multiple key-value pairs for the feature using the --value flag.`,
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

		featureMap := make(map[string]interface{}) // Parse each --value into the map
		for _, v := range featureValues {
			var keyValue map[string]interface{}
			err := json.Unmarshal([]byte(v), &keyValue)
			if err != nil {
				fmt.Printf("Error parsing value: %s\n", v)
				continue
			}
			for k, val := range keyValue {
				featureMap[k] = val
			}
		}

		dc.AddFeature(featureKey, featureMap)

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
	rootCmd.AddCommand(addFeaturesCmd)
	addFeaturesCmd.Flags().StringVarP(&featureKey, "key", "k", "", "Feature key (required)")
	addFeaturesCmd.Flags().StringArrayVarP(&featureValues, "value", "v", []string{}, "Feature values in JSON format (e.g., '{\"enabled\":true}') (required)")
	addFeaturesCmd.MarkFlagRequired("key")
	addFeaturesCmd.MarkFlagRequired("value")
}
