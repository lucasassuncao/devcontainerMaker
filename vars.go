package main

var extensions = map[string]string{
	"JSON":                          "zainchen.json",
	"YAML":                          "redhat.vscode-yaml",
	"Kubernetes":                    "ms-kubernetes-tools.vscode-kubernetes-tools",
	"Docker":                        "ms-azuretools.vscode-docker",
	"Terraform":                     "hashicorp.terraform",
	"Go":                            "golang.go",
	"GitLens":                       "eamodio.gitlens",
	"GitLab Workflow":               "gitlab.gitlab-workflow",
	"Dev Containers":                "ms-vscode-remote.remote-containers",
	"PowerShell":                    "ms-vscode.PowerShell",
	"Even Better TOML":              "tamasfe.even-better-toml",
	"Markdown Preview Enhanced":     "shd101wyy.markdown-preview-enhanced",
	"Syntax Highlight for Cucumber": "stevejpurves.cucumber",
	"Makefile Tools":                "ms-vscode.makefile-tools",
	"WSL":                           "ms-vscode-remote.remote-wsl",
}

var settings = map[string]interface{}{
	"editor.formatOnSave":                      true,
	"files.autoSave":                           "afterDelay",
	"terminal.integrated.defaultProfile.linux": "bash",
	"go.inlayHints.functionTypeParameters":     true,
	"go.inlayHints.parameterNames":             true,
	"go.inlayHints.rangeVariableTypes":         true,
	"go.inlayHints.assignVariableTypes":        true,
	"go.inlayHints.compositeLiteralFields":     true,
	"go.inlayHints.compositeLiteralTypes":      true,
	"go.inlayHints.constantValues":             true,
	"go.toolsManagement.checkForUpdates":       "off",
	"go.useLanguageServer":                     true,
	"go.lintTool":                              "golangci-lint",
	"go.lintFlags":                             []string{"--fast"},
	"go.lintOnSave":                            "file",
	"goTestExplorer.profiler.showCodeLens":     true,
	"go.enableCodeLens":                        map[string]interface{}{"runtest": true},
	"gopls":                                    map[string]interface{}{"build.buildFlags": []string{}, "ui.diagnostic.staticcheck": true},
}

var features = map[string]map[string]interface{}{
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
