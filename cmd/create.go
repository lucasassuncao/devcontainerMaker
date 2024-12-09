package cmd

import (
	"devcontainerMaker/internal/config"
	"devcontainerMaker/internal/model"
	"devcontainerMaker/internal/repository"
	"devcontainerMaker/internal/service"
	"devcontainerMaker/internal/utils"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var (
	v                        = validator.New()
	typ               string = ""
	name              string = ""
	image             string = ""
	dockerFile        string = ""
	context           string = ""
	dockerComposeFile string = ""
	svc               string = ""
	shutdownAction    string = ""
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Add a VSCode extension to the DevContainer",
	Run: func(cmd *cobra.Command, args []string) {
		dc, err := model.NewDevContainer().Initialize(typ)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = dc.SetName(name)
		if err != nil {
			fmt.Println(err.Error())
		}

		switch dc.Type {
		case "image":
			err = dc.SetImage(image)
			if err != nil {
				fmt.Println(err.Error())
			}
		case "dockerfile":
			err = dc.SetBuildDockerfile(dockerFile)
			if err != nil {
				fmt.Println(err.Error())
			}

			err = dc.SetBuildContext(context)
			if err != nil {
				fmt.Println(err.Error())
			}
		case "dockercompose":
			err = dc.SetDockerComposeFile(dockerComposeFile)
			if err != nil {
				fmt.Println(err.Error())
			}

			err = dc.SetService(svc)
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		err = dc.SetShutdownAction(shutdownAction)
		if err != nil {
			fmt.Println(err.Error())
		}

		// EXTENSIONS SECTION
		utils.ClearScreen()
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
		utils.ClearScreen()
		pterm.DefaultBasicText.Println(pterm.LightBlue("Configuring Devcontainer's vscode Settings..."))
		selectedSettings := service.GetMultiselectOptionsFromMap(config.DefaultSettings, service.RunInteractiveMultiselect)
		err = dc.SetSettings(selectedSettings)
		if err != nil {
			fmt.Println(err.Error())
		}

		// FEATURES SECTION
		utils.ClearScreen()
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

		utils.ClearScreen()
		fmt.Println("Prettifying JSON...")
		d, err := utils.PrettifyDevContainerJSON(dc)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("Generating devcontainer.json file...")
		err = repository.SaveDevContainerJSONFile(d)
		if err != nil {
			_ = fmt.Errorf("failed to create devcontainer.json file %v", err.Error())
		}

		fmt.Println("Your Dev Container is ready c:")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&typ, "type", "t", "", "Specify the type of the container (image, dockerfile or dockercompose)")
	createCmd.Flags().StringVarP(&name, "name", "n", "", "Provide a name for the container")
	createCmd.Flags().StringVarP(&image, "image", "i", "", "Define the image for the container")
	createCmd.Flags().StringVarP(&dockerFile, "dockerFile", "d", "", "Path to the Dockerfile (Default is \"Dockerfile\")")
	createCmd.Flags().StringVarP(&context, "context", "c", "", "Define the build context (Default is \".\")")
	createCmd.Flags().StringVarP(&dockerComposeFile, "dockerComposeFile", "D", "", "Path to the Docker Compose file")
	createCmd.Flags().StringVarP(&svc, "serviceName", "k", "s", "Specify the service name")
	createCmd.Flags().StringVarP(&shutdownAction, "shutdownAction", "a", "", "Define the container shutdown action.\nIF type = image or dockerfile (none, stopContainer)\nIF type = dockercompose (none, stopCompose)")
}
