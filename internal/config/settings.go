package config

var Settings = map[string]interface{}{
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
