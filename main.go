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

	//time.Sleep(time.Second * 5)

	/*	content, err := ReadDevContainerJsonFile()
		if err != nil {
			_ = fmt.Errorf("failed to read devcontainer.json file %v", err.Error())
		}

		var dev = NewDevContainer()
		err = JsonToStruct(content, dev)
		if err != nil {
			_ = fmt.Errorf("failed to unmarshal devcontainer.json file %v", err.Error())
		}

		dev.AddExtension("lucao.extension")

		d, err = PrettifyJson(dev)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = OutputJsonToFile(d)
		if err != nil {
			_ = fmt.Errorf("failed to create devcontainer.json file %v", err.Error())
		}*/

	clearScreen()

	pterm.DefaultBasicText.Println(pterm.LightBlue("Configuring Devcontainer's VSCode Extensions..."))
	selectedExtensions, _ := runInteractiveMultiselect(extensions)
	dc.setExtensions(selectedExtensions)

	clearScreen()

	pterm.DefaultBasicText.Println(pterm.LightBlue("Configuring Devcontainer's VSCode Settings..."))
	selectedSettings := getMultiselectOptionsFromMap(settings, runInteractiveMultiselect)
	dc.setSettings(selectedSettings)

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
