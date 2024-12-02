package main

import (
	"fmt"
	"github.com/pterm/pterm"
)

func main() {
	dc := NewDevContainer()

	dc.setName("TestDevContainer")
	dc.setBuildDockerfile("Dockerfile")
	dc.setShutdownAction("stopContainer")

	// EXTENSIONS SECTION
	clearScreen()
	pterm.DefaultBasicText.Println(pterm.LightBlue("Configuring Devcontainer's VSCode Extensions..."))
	selectedExtensions := getMultiselectOptionsFromMap(extensions, runInteractiveMultiselect)
	var se []string
	for _, s := range selectedExtensions {
		se = append(se, s)
	}
	dc.setExtensions(se)

	// SETTINGS SECTION
	clearScreen()
	pterm.DefaultBasicText.Println(pterm.LightBlue("Configuring Devcontainer's VSCode Settings..."))
	selectedSettings := getMultiselectOptionsFromMap(settings, runInteractiveMultiselect)
	dc.setSettings(selectedSettings)

	// FEATURES SECTION
	clearScreen()
	pterm.DefaultBasicText.Println(pterm.LightBlue("Configuring Devcontainer's Features..."))
	selectedFeatures := getMultiselectOptionsFromMap(features, runInteractiveMultiselect)
	dc.setFeatures(selectedFeatures)

	d, err := PrettifyJson(dc)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = OutputJsonToFile(d)
	if err != nil {
		_ = fmt.Errorf("failed to create devcontainer.json file %v", err.Error())
	}

}
