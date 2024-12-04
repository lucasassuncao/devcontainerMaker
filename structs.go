package main

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
