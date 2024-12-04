package main

import (
	"fmt"
	"github.com/pterm/pterm"
)

func main() {
	dc := NewDevContainer().
		WithName().
		WithBuildDockerFile().
		WithShutdownAction().
		WithFeatures().
		WithExtensions().
		WithSettings()

	devContainerName, err := pterm.DefaultInteractiveTextInput.WithDefaultText(pterm.FgGreen.Sprint("Enter your Dev Container name")).Show()

	dc.SetName(devContainerName)
	dc.SetBuildDockerfile("Dockerfile")
	dc.SetShutdownAction("stopContainer")

	// EXTENSIONS SECTION
	ClearScreen()
	pterm.DefaultBasicText.Println(pterm.LightBlue("Configuring Devcontainer's VSCode Extensions..."))
	selectedExtensions := GetMultiselectOptionsFromMap(extensions, runInteractiveMultiselect)
	var se []string
	for _, s := range selectedExtensions {
		se = append(se, s)
	}
	err = dc.SetExtensions(se)
	if err != nil {
		pterm.Fatal.Println(err.Error())
	}

	// SETTINGS SECTION
	ClearScreen()
	pterm.DefaultBasicText.Println(pterm.LightBlue("Configuring Devcontainer's VSCode Settings..."))
	selectedSettings := GetMultiselectOptionsFromMap(settings, runInteractiveMultiselect)
	err = dc.SetSettings(selectedSettings)
	if err != nil {
		pterm.Fatal.Println(err.Error())
	}

	// FEATURES SECTION
	ClearScreen()
	pterm.DefaultBasicText.Println(pterm.LightBlue("Configuring Devcontainer's Features..."))
	selectedFeatures := GetMultiselectOptionsFromMap(features, runInteractiveMultiselect)
	err = dc.SetFeatures(selectedFeatures)
	if err != nil {
		pterm.Fatal.Println(err.Error())
	}

	d, err := PrettifyJson(dc)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = OutputJsonToFile(d)
	if err != nil {
		_ = fmt.Errorf("failed to create devcontainer.json file %v", err.Error())
	}

	ClearScreen()
	fmt.Println("Your Dev Container is ready c:")
}
