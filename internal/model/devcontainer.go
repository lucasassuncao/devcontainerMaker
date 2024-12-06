package model

import (
	"devcontainerMaker/internal/service"
	"fmt"
)

// DevContainer struct represents a configuration for a development container,
// which can include build settings, shutdown actions, feature configurations, and customizations.
// Some fields are optional and will be omitted from the JSON representation if not set, allowing flexibility in the configuration
type DevContainer struct {
	Name              string                            `json:"name" validate:"required"`
	Type              string                            `json:"-" validate:"required,oneof=image dockerfile dockercompose"`
	Image             string                            `json:"image,omitempty"`
	Build             *build                            `json:"build,omitempty"`
	DockerComposeFile string                            `json:"dockerComposeFile,omitempty"`
	Service           string                            `json:"service,omitempty"`
	ShutdownAction    string                            `json:"shutdownAction,omitempty" validate:"oneof=none stopContainer stopCompose"`
	Features          map[string]map[string]interface{} `json:"features,omitempty"`
	Customizations    *customizations                   `json:"customizations,omitempty"`
}

// build struct encapsulates the build-related configuration for a development container,
// allowing the user to define the specifics of how the container should be built. This field is optional
type build struct {
	Dockerfile string `json:"dockerfile"`
	Context    string `json:"context"`
}

// customizations struct allows users to define container-specific customizations, such as vscode settings or extensions.
// It currently supports vscode-related configurations, but additional customizations could be added in the future.
type customizations struct {
	VSCode *vscode `json:"vscode,omitempty"`
}

// vscode struct provides configuration options for Visual Studio Code inside the development container.
// It includes settings for extensions to be installed and custom settings to be applied. These fields are optional
type vscode struct {
	Extensions []string               `json:"extensions,omitempty"`
	Settings   map[string]interface{} `json:"settings,omitempty"`
}

// NewDevContainer creates and returns a pointer to a new instance of a DevContainer.
// It initializes the DevContainer and prepares it for further configuration.
func NewDevContainer() *DevContainer {
	return &DevContainer{}
}

func (d *DevContainer) Initialize(typ string) (*DevContainer, error) {
	var opts = []string{"image", "dockerfile", "dockercompose"}

	// Prompt for type if not provided
	if typ == "" {
		var err error
		typ, err = service.RunInteractiveSelect(opts, "Select your DevContainer type")
		if err != nil {
			return nil, fmt.Errorf("error during initialization: %w", err)
		}
	}
	d.Type = typ

	if err := validateOptions(d.Type, opts); err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	// Initialize the DevContainer fields
	d.withName()
	d.withType()
	d.withShutdownAction()
	d.withFeatures()
	d.withExtensions()
	d.withSettings()

	return d, nil
}

func (d *DevContainer) withName() *DevContainer {
	d.Name = ""
	return d
}
func (d *DevContainer) withType() *DevContainer {
	switch d.Type {
	case "image":
		d.withImage()
	case "dockerfile":
		d.withBuild()
	case "dockercompose":
		d.withDockerCompose()
		d.withService()
	}
	return d
}
func (d *DevContainer) withImage() *DevContainer {
	d.Image = ""
	return d
}
func (d *DevContainer) withBuild() *DevContainer {
	d.Build = &build{
		Dockerfile: "",
		Context:    "",
	}
	return d
}
func (d *DevContainer) withDockerCompose() *DevContainer {
	d.DockerComposeFile = ""
	return d
}
func (d *DevContainer) withService() *DevContainer {
	d.Service = ""
	return d
}
func (d *DevContainer) withShutdownAction() *DevContainer {
	d.ShutdownAction = ""
	return d
}
func (d *DevContainer) withFeatures() *DevContainer {
	d.Features = make(map[string]map[string]interface{})
	return d
}
func (d *DevContainer) withExtensions() *DevContainer {
	d.Customizations = &customizations{
		VSCode: &vscode{
			Extensions: make([]string, 0),
		},
	}
	return d
}
func (d *DevContainer) withSettings() *DevContainer {
	d.Customizations = &customizations{
		VSCode: &vscode{
			Settings: make(map[string]interface{}),
		},
	}
	return d
}

func (d *DevContainer) SetName(name string) error {
	if name == "" {
		name, err := service.RunInteractiveTextInput("Enter your DevContainer name", "")
		if err != nil {
			return fmt.Errorf("could not set DevContainer name: %v", err)
		}
		d.Name = name
		return nil
	}

	d.Name = name
	return nil
}
func (d *DevContainer) SetImage(image string) error {
	if d.Type != "image" {
		return fmt.Errorf("WARNING: DevContainer is set to '%s' not 'image'", d.Type)
	}

	if image == "" {
		image, err := service.RunInteractiveTextInput("Enter the Docker image to use", "")
		if err != nil {
			return fmt.Errorf("could not set DevContainer image: %v", err)
		}
		d.Image = image
		return nil
	}

	d.Image = image
	return nil
}
func (d *DevContainer) SetBuildDockerfile(file string) error {
	if d.Type != "dockerfile" {
		return fmt.Errorf("WARNING: DevContainer is set to '%s' not 'Dockerfile'", d.Type)
	}

	if file == "" {
		file, err := service.RunInteractiveTextInput("Please enter the name of your Dockerfile", "Dockerfile")
		if err != nil {
			return fmt.Errorf("could not set DevContainer Dockerfile: %v", err)
		}
		d.Build.Dockerfile = file
		return nil
	}

	d.Build.Dockerfile = file
	return nil
}
func (d *DevContainer) SetBuildContext(path string) error {
	if d.Type != "dockerfile" {
		return fmt.Errorf("WARNING: DevContainer is set to '%s' not 'dockerfile'", d.Type)
	}

	if path == "" {
		path, err := service.RunInteractiveTextInput("Enter the path where the Docker build should be executed from (default is the current directory)", ".")
		if err != nil {
			return fmt.Errorf("could not set DevContainer Build Context: %v", err)
		}
		d.Build.Context = path
		return nil
	}

	d.Build.Context = path
	return nil
}
func (d *DevContainer) SetDockerComposeFile(file string) error {
	if d.Type != "dockercompose" {
		return fmt.Errorf("WARNING: DevContainer is set to '%s' not 'dockercompose'", d.Type)
	}

	if file == "" {
		file, err := service.RunInteractiveTextInput("Enter the Docker Compose file name", "")
		if err != nil {
			return fmt.Errorf("could not set DevContainer DockerComposeFile: %v", err)
		}
		d.DockerComposeFile = file
		return nil
	}

	d.DockerComposeFile = file
	return nil
}
func (d *DevContainer) SetService(svcName string) error {
	if d.Type != "dockercompose" {
		return fmt.Errorf("WARNING: DevContainer is set to '%s' not 'dockercompose'", d.Type)
	}

	if svcName == "" {
		svcName, err := service.RunInteractiveTextInput("Enter the service name", "")
		if err != nil {
			return fmt.Errorf("could not set DevContainer Service: %v", err)
		}
		d.Service = svcName
		return nil
	}

	d.Service = svcName
	return nil
}
func (d *DevContainer) SetShutdownAction(action string) error {
	var opts []string

	switch d.Type {
	case "image", "dockerfile":
		opts = []string{"none", "stopContainer"}
	case "dockercompose":
		opts = []string{"none", "stopCompose"}
	}

	if action == "" {
		action, err := service.RunInteractiveSelect(opts, "Select your DevContainer shutdown action")
		if err != nil {
			return fmt.Errorf("could not set DevContainer ShutdownAction: %v", err)
		}
		d.ShutdownAction = action
		return nil
	}

	if err := validateOptions(action, opts); err != nil {
		return fmt.Errorf("error: %w", err)
	}

	d.ShutdownAction = action
	return nil
}

func (d *DevContainer) SetFeatures(features map[string]map[string]interface{}) error {
	if len(features) == 0 {
		return fmt.Errorf("no features selected... skipping feature set")
	}

	if d.Features != nil {
		d.Features = features
	}

	return nil
}
func (d *DevContainer) SetExtensions(extensions []string) error {
	if len(extensions) == 0 {
		return fmt.Errorf("no extensions selected... skipping extension set")
	}

	if d.Customizations.VSCode != nil {
		d.Customizations.VSCode.Extensions = extensions
	}

	return nil
}
func (d *DevContainer) SetSettings(settings map[string]interface{}) error {
	if len(settings) == 0 {
		return fmt.Errorf("no settings selected... skipping settings set")
	}

	if d.Customizations.VSCode != nil {
		d.Customizations.VSCode.Settings = settings
	}

	return nil
}

func (d *DevContainer) AddFeature(key string, value map[string]interface{}) {
	if d.Features == nil {
		d.Features = make(map[string]map[string]interface{})
	}

	d.Features[key] = value
}
func (d *DevContainer) AddExtension(extension string) {
	if d.Customizations.VSCode.Extensions == nil {
		d.Customizations.VSCode.Extensions = []string{}
	}

	d.Customizations.VSCode.Extensions = append(d.Customizations.VSCode.Extensions, extension)
}
func (d *DevContainer) AddSetting(key string, value interface{}) {
	if d.Customizations.VSCode.Settings == nil {
		d.Customizations.VSCode.Settings = make(map[string]interface{})
	}

	d.Customizations.VSCode.Settings[key] = value
}

func validateOptions(s string, validOptions []string) error {
	for _, validOption := range validOptions {
		if s == validOption {
			return nil
		}
	}
	return fmt.Errorf("invalid DevContainer type: %s. Valid types are: %v", s, validOptions)
}
