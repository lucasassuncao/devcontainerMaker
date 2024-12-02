package main

func (d *DevContainer) setName(name string) {
	d.Name = name
}

func (d *DevContainer) setBuildDockerfile(dockerfile string) {
	d.Build.Dockerfile = dockerfile
}

func (d *DevContainer) setShutdownAction(shutdownAction string) {
	d.ShutdownAction = shutdownAction
}

func (d *DevContainer) setFeatures(features map[string]map[string]interface{}) {
	d.Features = features
}

func (d *DevContainer) setExtensions(extensions []string) {
	d.Customizations.VSCode.Extensions = extensions
}

func (d *DevContainer) setSettings(settings map[string]interface{}) {
	d.Customizations.VSCode.Settings = settings
}

func (d *DevContainer) addFeature(key string, value map[string]interface{}) {
	if d.Features == nil {
		d.Features = make(map[string]map[string]interface{})
	}

	d.Features[key] = value
}

func (d *DevContainer) addExtension(extension string) {
	if d.Customizations.VSCode.Extensions == nil {
		d.Customizations.VSCode.Extensions = []string{}
	}

	d.Customizations.VSCode.Extensions = append(d.Customizations.VSCode.Extensions, extension)
}

func (d *DevContainer) addSetting(key string, value interface{}) {
	if d.Customizations.VSCode.Settings == nil {
		d.Customizations.VSCode.Settings = make(map[string]interface{})
	}

	d.Customizations.VSCode.Settings[key] = value
}
