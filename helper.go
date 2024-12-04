package main

import (
	"atomicgo.dev/keyboard/keys"
	"encoding/json"
	"fmt"
	"github.com/pterm/pterm"
	"os"
	"os/exec"
	"runtime"
)

func OutputJsonToFile(data []byte) error {
	file, err := os.Create("devcontainer.json")
	if err != nil {
		return fmt.Errorf("error creating file %v", err.Error())
	}
	defer file.Close()

	// Write the indented JSON to the file
	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("error writing to file %v", err.Error())

	}

	return nil
}

func PrettifyJson(dc *DevContainer) ([]byte, error) {
	d, err := json.MarshalIndent(dc, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("error prettifying json %v", err.Error())
	}
	return d, nil
}

func ReadDevContainerJsonFile() ([]byte, error) {
	file, err := os.ReadFile("devcontainer.json")
	if err != nil {
		return nil, fmt.Errorf("error reading file %v", err.Error())
	}

	return file, nil
}

func JsonToStruct(content []byte, dc *DevContainer) error {
	err := json.Unmarshal(content, &dc)
	if err != nil {
		return fmt.Errorf("error unmarshalling json %v", err.Error())
	}

	return nil
}

func runInteractiveMultiselect(opts []string) ([]string, error) {
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

func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
