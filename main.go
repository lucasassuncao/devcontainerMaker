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
	dc := model.NewDevContainer().Initialize()

	err := v.Struct(dc)
	if err != nil {
		pterm.Error.Print(err.Error())
		return
	}

	// EXTENSIONS SECTION
	utils.ClearScreen()
	pterm.DefaultBasicText.Println(pterm.LightBlue("Configuring Devcontainer's vscode Extensions..."))
	selectedExtensions := service.GetMultiselectOptionsFromMap(config.Extensions, service.RunInteractiveMultiselect)
	var se []string
	for _, s := range selectedExtensions {
		se = append(se, s)
	}
	err = dc.SetExtensions(se)
	if err != nil {
		pterm.Fatal.Println(err.Error())
	}

	// SETTINGS SECTION
	utils.ClearScreen()
	pterm.DefaultBasicText.Println(pterm.LightBlue("Configuring Devcontainer's vscode Settings..."))
	selectedSettings := service.GetMultiselectOptionsFromMap(config.Settings, service.RunInteractiveMultiselect)
	err = dc.SetSettings(selectedSettings)
	if err != nil {
		pterm.Fatal.Println(err.Error())
	}

	// FEATURES SECTION
	utils.ClearScreen()
	pterm.DefaultBasicText.Println(pterm.LightBlue("Configuring Devcontainer's Features..."))
	selectedFeatures := service.GetMultiselectOptionsFromMap(config.Features, service.RunInteractiveMultiselect)
	err = dc.SetFeatures(selectedFeatures)
	if err != nil {
		pterm.Fatal.Println(err.Error())
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

	utils.ClearScreen()
	fmt.Println("Your Dev Container is ready c:")
}
