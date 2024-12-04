package main

import (
	"errors"
)

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
