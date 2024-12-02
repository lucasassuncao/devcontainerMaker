package main

type DevContainer struct {
	Name           string                            `json:"name,omitempty"`
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
