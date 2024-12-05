package config

var Features = map[string]map[string]interface{}{
	"ghcr.io/devcontainers/features/common-utils:2": {
		"installZsh":      "false",
		"username":        "none",
		"userUid":         "automatic",
		"userGid":         "automatic",
		"upgradePackages": "false",
	},
	"ghcr.io/devcontainers/features/powershell:1": {
		"version": "latest",
	},
	"ghcr.io/devcontainers/features/git:1": {
		"version": "latest",
		"ppa":     "true",
	},
	"ghcr.io/devcontainers/features/terraform:1": {
		"version":              "latest",
		"tflint":               "latest",
		"installTerraformDocs": "true",
	},
	"ghcr.io/devcontainers/features/go:1": {
		"golangciLintVersion": "latest",
		"version":             "1.23.3",
	},
	"ghcr.io/eitsupi/devcontainer-features/jq-likes:2.1.0": {
		"jqVersion": "latest",
		"yqVersion": "latest",
	},
	"ghcr.io/dhoeric/features/google-cloud-cli:1": {
		"version":                    "latest",
		"installGkeGcloudAuthPlugin": false,
	},
	"ghcr.io/devcontainers/features/aws-cli:1": {
		"version": "latest",
	},
	"ghcr.io/devcontainers/features/kubectl-helm-minikube:1": {
		"version": "latest",
		"helm":    "latest",
	},
	"ghcr.io/mpriscella/features/kind:1": {
		"version": "latest",
	},
	"ghcr.io/dhoeric/features/k9s:1": {
		"version": "latest",
	},
}
