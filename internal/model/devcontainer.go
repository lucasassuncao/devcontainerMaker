package model

import (
	"devcontainerMaker/internal/service"
	"errors"
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

func (d *DevContainer) Initialize() *DevContainer {
	d.withName()
	d.withType()
	d.withShutdownAction()
	d.withFeatures()
	d.withExtensions()
	d.withSettings()

	return d
}

// withName initializes the DevContainer's Name field with an empty string.
// Subsequent configuration is possible by using the method SetName.
func (d *DevContainer) withName() *DevContainer {
	d.SetName()
	return d
}

func (d *DevContainer) withType() *DevContainer {
	d.SetType()

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
	d.SetImage()
	return d
}

// WithBuild initializes the DevContainer's build -> Dockerfile field with an empty string
// Subsequent configuration is possible by using the method SetBuild.
func (d *DevContainer) withBuild() *DevContainer {
	d.SetBuild()
	return d
}

func (d *DevContainer) withDockerCompose() *DevContainer {
	d.SetDockerComposeFile()
	return d
}

func (d *DevContainer) withService() *DevContainer {
	d.SetService()
	return d
}

// WithShutdownAction initializes the DevContainer's ShutdownAction field with an empty string
// Subsequent configuration is possible by using the method SetShutdownAction.
func (d *DevContainer) withShutdownAction() *DevContainer {
	d.SetShutdownAction()
	return d
}

// WithFeatures initializes the DevContainer's Feature field with an empty map[string]map[string]interface{}
// Subsequent configuration is possible by using the method SetFeatures.
func (d *DevContainer) withFeatures() *DevContainer {
	d.Features = make(map[string]map[string]interface{})
	return d
}

// WithExtensions initializes the DevContainer's customizations -> vscode -> Extensions field with an empty []string
// Subsequent configuration is possible by using the method SetExtensions.
func (d *DevContainer) withExtensions() *DevContainer {
	d.Customizations = &customizations{
		VSCode: &vscode{
			Extensions: make([]string, 0),
		},
	}
	return d
}

// WithSettings initializes the DevContainer's customizations -> vscode -> Settings field with an empty map[string]interface{}
// Subsequent configuration is possible by using the method SetSettings.
func (d *DevContainer) withSettings() *DevContainer {
	d.Customizations = &customizations{
		VSCode: &vscode{
			Settings: make(map[string]interface{}),
		},
	}
	return d
}

func (d *DevContainer) SetType() {
	opts := []string{"image", "dockerfile", "dockercompose"}
	typ, _ := service.RunInteractiveSelect(opts, "Select your Dev Container type")

	d.Type = typ
}

// SetName sets the Name field of the DevContainer.
// If the provided name is an empty string, it defaults to "MyDevContainer"
func (d *DevContainer) SetName() {
	name, _ := service.RunInteractiveTextInput("Enter your Dev Container name", "")
	d.Name = name
}

func (d *DevContainer) SetImage() {
	if d.Type != "image" {
		fmt.Printf("WARNING: Dev Container is set to use a '%s' not Image", d.Type)
		return
	}

	image, _ := service.RunInteractiveTextInput("Enter the Docker image to use", "")
	d.Image = image
}

// SetBuild sets the Dockerfile field within the build struct of the DevContainer.
// If the build field is not initialized, the method has no effect
func (d *DevContainer) SetBuild() {
	if d.Type != "dockerfile" {
		fmt.Printf("WARNING: Dev Container is set to use a '%s' not Dockerfile", d.Type)
		return
	}

	file, _ := service.RunInteractiveTextInput("Please enter the name of your Dockerfile", "Dockerfile")
	context, _ := service.RunInteractiveTextInput("Enter the path where the Docker build should be executed from (default is the current directory)", ".")

	d.Build = &build{
		Dockerfile: file,
		Context:    context,
	}
}

func (d *DevContainer) SetDockerComposeFile() {
	if d.Type != "dockercompose" {
		fmt.Printf("WARNING: Dev Container is set to use a '%s' not DockerCompose", d.Type)
		return
	}

	compose, _ := service.RunInteractiveTextInput("Enter the Docker Compose file name", "")
	d.DockerComposeFile = compose
}

func (d *DevContainer) SetService() {
	if d.Type != "dockercompose" {
		fmt.Printf("WARNING: Dev Container is set to use a '%s' not DockerCompose", d.Type)
		return
	}

	svc, _ := service.RunInteractiveTextInput("Enter the service name", "")
	d.Service = svc
}

// SetShutdownAction sets the ShutdownAction field of the DevContainer.
func (d *DevContainer) SetShutdownAction() {
	var opts []string

	switch d.Type {
	case "image":
		opts = []string{"none", "stopContainer"}
	case "dockerfile":
		opts = []string{"none", "stopContainer"}
	case "dockercompose":
		opts = []string{"none", "stopCompose"}

	}

	sa, _ := service.RunInteractiveSelect(opts, "Select your Dev Container shutdown action")

	d.ShutdownAction = sa
}

// SetFeatures sets the Features field of the DevContainer with the provided map.
// If the Features field is not initialized, it returns an error indicating the field must be initialized first.
func (d *DevContainer) SetFeatures(features map[string]map[string]interface{}) error {
	if d.Features != nil {
		d.Features = features
		return nil
	}

	return errors.New("feature field not initialized")
}

// SetExtensions sets the Extensions field of the DevContainer with the provided slice.
// If the Extensions field is not initialized, it returns an error indicating the field must be initialized first.
func (d *DevContainer) SetExtensions(extensions []string) error {
	if d.Customizations.VSCode != nil {
		d.Customizations.VSCode.Extensions = extensions
		return nil
	}

	return errors.New("extension field not initialized")
}

// SetSettings sets the Settings field of the DevContainer with the provided map.
// If the Settings field is not initialized, it returns an error indicating the field must be initialized first.
func (d *DevContainer) SetSettings(settings map[string]interface{}) error {
	if d.Customizations.VSCode != nil {
		d.Customizations.VSCode.Settings = settings
		return nil
	}

	return errors.New("setting field not initialized")
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
