package main

import (
	"devcontainerMaker/internal/config"
	"devcontainerMaker/internal/model"
	"devcontainerMaker/internal/repository"
	"devcontainerMaker/internal/service"
	"devcontainerMaker/internal/utils"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/pterm/pterm"
)

var v = validator.New()

func main() {
	dc, err := model.NewDevContainer().Initialize("")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = dc.SetName("")
	if err != nil {
		fmt.Println(err.Error())
	}

	switch dc.Type {
	case "image":
		err = dc.SetImage("")
		if err != nil {
			fmt.Println(err.Error())
		}
	case "dockerfile":
		err = dc.SetBuildDockerfile("")
		if err != nil {
			fmt.Println(err.Error())
		}

		err = dc.SetBuildContext("")
		if err != nil {
			fmt.Println(err.Error())
		}
	case "dockercompose":
		err = dc.SetDockerComposeFile("")
		if err != nil {
			fmt.Println(err.Error())
		}

		err = dc.SetService("")
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	err = dc.SetShutdownAction("")
	if err != nil {
		fmt.Println(err.Error())
	}

	// EXTENSIONS SECTION
	//utils.ClearScreen()
	pterm.DefaultBasicText.Println(pterm.LightBlue("Configuring Devcontainer's vscode Extensions..."))
	selectedExtensions := service.GetMultiselectOptionsFromMap(config.DefaultExtensions, service.RunInteractiveMultiselect)
	var se []string
	for _, s := range selectedExtensions {
		se = append(se, s)
	}
	err = dc.SetExtensions(se)
	if err != nil {
		fmt.Println(err.Error())
	}

	// SETTINGS SECTION
	//utils.ClearScreen()
	pterm.DefaultBasicText.Println(pterm.LightBlue())
	selectedSettings := service.GetMultiselectOptionsFromMap(config.DefaultSettings, service.RunInteractiveMultiselect)
	err = dc.SetSettings(selectedSettings)
	if err != nil {
		fmt.Println(err.Error())
	}

	// FEATURES SECTION
	//utils.ClearScreen()
	pterm.DefaultBasicText.Println(pterm.LightBlue("Configuring Devcontainer's Features..."))
	selectedFeatures := service.GetMultiselectOptionsFromMap(config.DefaultFeatures, service.RunInteractiveMultiselect)
	err = dc.SetFeatures(selectedFeatures)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Validate struct to check fields
	err = v.Struct(dc)
	if err != nil {
		pterm.Error.Print(err.Error())
		return
	}

	d, err := utils.PrettifyDevContainerJSON(dc)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = repository.SaveDevContainerJSONFile(d)
	if err != nil {
		_ = fmt.Errorf("failed to create devcontainer.json file %v", err.Error())
	}

	//utils.ClearScreen()
	fmt.Println("Your Dev Container is ready c:")
}
