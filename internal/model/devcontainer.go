package model

import "errors"

// DevContainer struct represents a configuration for a development container,
// which can include build settings, shutdown actions, feature configurations, and customizations.
// Some fields are optional and will be omitted from the JSON representation if not set, allowing flexibility in the configuration
type DevContainer struct {
	Name           string                            `json:"name"`
	Build          *Build                            `json:"build,omitempty"`
	ShutdownAction string                            `json:"shutdownAction,omitempty"`
	Features       map[string]map[string]interface{} `json:"features,omitempty"`
	Customizations *Customizations                   `json:"customizations,omitempty"`
}

// Build struct encapsulates the build-related configuration for a development container.
// In this case, it contains the path to the Dockerfile,
// allowing the user to define the specifics of how the container should be built. This field is optional
type Build struct {
	Dockerfile string `json:"dockerfile,omitempty"`
}

// Customizations struct allows users to define container-specific customizations, such as VSCode settings or extensions.
// It currently supports VSCode-related configurations, but additional customizations could be added in the future.
type Customizations struct {
	VSCode *VSCode `json:"vscode,omitempty"`
}

// VSCode struct provides configuration options for Visual Studio Code inside the development container.
// It includes settings for extensions to be installed and custom settings to be applied. These fields are optional
type VSCode struct {
	Extensions []string               `json:"extensions,omitempty"`
	Settings   map[string]interface{} `json:"settings,omitempty"`
}

// NewDevContainer creates and returns a pointer to a new instance of a DevContainer.
// It initializes the DevContainer and prepares it for further configuration.
func NewDevContainer() *DevContainer {
	return &DevContainer{}
}

// WithName initializes the DevContainer's Name field with an empty string.
// Subsequent configuration is possible by using the method SetName.
func (d *DevContainer) WithName() *DevContainer {
	d.Name = ""
	return d
}

// WithBuildDockerFile initializes the DevContainer's Build -> Dockerfile field with an empty string
// Subsequent configuration is possible by using the method SetBuildDockerfile.
func (d *DevContainer) WithBuildDockerFile() *DevContainer {
	d.Build = &Build{
		Dockerfile: "",
	}
	return d
}

// WithShutdownAction initializes the DevContainer's ShutdownAction field with an empty string
// Subsequent configuration is possible by using the method SetShutdownAction.
func (d *DevContainer) WithShutdownAction() *DevContainer {
	d.ShutdownAction = ""
	return d
}

// WithFeatures initializes the DevContainer's Feature field with an empty map[string]map[string]interface{}
// Subsequent configuration is possible by using the method SetFeatures.
func (d *DevContainer) WithFeatures() *DevContainer {
	d.Features = make(map[string]map[string]interface{})
	return d
}

// WithExtensions initializes the DevContainer's Customizations -> VSCode -> Extensions field with an empty []string
// Subsequent configuration is possible by using the method SetExtensions.
func (d *DevContainer) WithExtensions() *DevContainer {
	d.Customizations = &Customizations{
		VSCode: &VSCode{
			Extensions: make([]string, 0),
		},
	}
	return d
}

// WithSettings initializes the DevContainer's Customizations -> VSCode -> Settings field with an empty map[string]interface{}
// Subsequent configuration is possible by using the method SetSettings.
func (d *DevContainer) WithSettings() *DevContainer {
	d.Customizations = &Customizations{
		VSCode: &VSCode{
			Settings: make(map[string]interface{}),
		},
	}
	return d
}

// SetName sets the Name field of the DevContainer.
// If the provided name is an empty string, it defaults to "MyDevContainer"
func (d *DevContainer) SetName(name string) {
	if name == "" {
		d.Name = "MyDevContainer"
		return
	}

	d.Name = name
}

// SetBuildDockerfile sets the Dockerfile field within the Build struct of the DevContainer.
// If the Build field is not initialized, the method has no effect
func (d *DevContainer) SetBuildDockerfile(dockerfile string) {
	if d.Build != nil {
		d.Build.Dockerfile = dockerfile
	}
}

// SetShutdownAction sets the ShutdownAction field of the DevContainer.
func (d *DevContainer) SetShutdownAction(shutdownAction string) {
	d.ShutdownAction = shutdownAction
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
