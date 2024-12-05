package service

import (
	"atomicgo.dev/keyboard/keys"
	"github.com/pterm/pterm"
)

// RunInteractiveMultiselect creates and displays an interactive multiselect menu using the pterm library.
// It allows the user to navigate through the provided options and select one or more items.
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

// GetMultiselectOptionsFromMap extracts the keys from the provided map[string]T and use them as options for selection using the runMultiselect function
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
