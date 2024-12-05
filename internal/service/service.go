package service

import (
	"atomicgo.dev/keyboard/keys"
	"github.com/pterm/pterm"
)

func RunInteractiveMultiselect(opts []string) ([]string, error) {
	printer := pterm.DefaultInteractiveMultiselect.
		WithOptions(opts).
		WithFilter(false).
		WithKeyConfirm(keys.Enter).
		WithKeySelect(keys.Space).
		WithMaxHeight(len(opts))

	return printer.Show()
}

/*func runInteractiveMultiselectWithDefaultOptions(opts []string) ([]string, error) {
	printer := pterm.DefaultInteractiveMultiselect.
		WithDefaultOptions(opts).
		WithFilter(false).
		WithKeyConfirm(keys.Enter).
		WithKeySelect(keys.Space).
		WithMaxHeight(len(opts))

	return printer.Show()
}*/

func GetMultiselectOptionsFromMap[T any](data map[string]T, runMultiselect func([]string) ([]string, error)) map[string]T {
	var opts []string
	for key := range data {
		opts = append(opts, key)
	}

	selectedKeys, _ := runMultiselect(opts)

	selectedItems := make(map[string]T)
	for _, key := range selectedKeys {
		if value, exists := data[key]; exists {
			selectedItems[key] = value
		}
	}
	return selectedItems
}
