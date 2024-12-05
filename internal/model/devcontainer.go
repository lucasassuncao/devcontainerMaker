package model

import "errors"

type DevContainer struct {
	Name           string                            `json:"name"`
	Build          *Build                            `json:"build,omitempty"`
	ShutdownAction string                            `json:"shutdownAction,omitempty"`
	Features       map[string]map[string]interface{} `json:"features,omitempty"`
	Customizations *Customizations                   `json:"customizations,omitempty"`
}

type Build struct {
	Dockerfile string `json:"dockerfile,omitempty"`
}

type Customizations struct {
	VSCode *VSCode `json:"vscode,omitempty"`
}

type VSCode struct {
	Extensions []string               `json:"extensions,omitempty"`
	Settings   map[string]interface{} `json:"settings,omitempty"`
}

// NewDevContainer creates and returns a pointer to a new instance of a DevContainer.
// It initializes the DevContainer and prepares it for further configuration.
func NewDevContainer() *DevContainer {
	return &DevContainer{}
}

func (d *DevContainer) WithName() *DevContainer {
	d.Name = ""
	return d
}

func (d *DevContainer) WithBuildDockerFile() *DevContainer {
	d.Build = &Build{
		Dockerfile: "",
	}
	return d
}

func (d *DevContainer) WithShutdownAction() *DevContainer {
	d.ShutdownAction = ""
	return d
}

func (d *DevContainer) WithFeatures() *DevContainer {
	d.Features = make(map[string]map[string]interface{})
	return d
}

func (d *DevContainer) WithExtensions() *DevContainer {
	d.Customizations = &Customizations{
		VSCode: &VSCode{
			Extensions: make([]string, 0),
		},
	}
	return d
}

func (d *DevContainer) WithSettings() *DevContainer {
	d.Customizations = &Customizations{
		VSCode: &VSCode{
			Settings: make(map[string]interface{}),
		},
	}
	return d
}

func (d *DevContainer) SetName(name string) {
	if name == "" {
		d.Name = "MyDevContainer"
		return
	}

	d.Name = name
}

func (d *DevContainer) SetBuildDockerfile(dockerfile string) {
	if d.Build != nil {
		d.Build.Dockerfile = dockerfile
	}
}

func (d *DevContainer) SetShutdownAction(shutdownAction string) {
	d.ShutdownAction = shutdownAction
}

func (d *DevContainer) SetFeatures(features map[string]map[string]interface{}) error {
	if d.Features != nil {
		d.Features = features
		return nil
	}

	return errors.New("feature field not initialized")
}

func (d *DevContainer) SetExtensions(extensions []string) error {
	if d.Customizations.VSCode != nil {
		d.Customizations.VSCode.Extensions = extensions
		return nil
	}

	return errors.New("extension field not initialized")
}

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
