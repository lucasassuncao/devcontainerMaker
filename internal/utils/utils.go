package utils

import (
	"devcontainerMaker/internal/model"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// PrettifyDevContainerJSON takes a DevContainer struct
// and converts it into a prettified (indented) JSON format using json.MarshalIndent
func PrettifyDevContainerJSON(dc *model.DevContainer) ([]byte, error) {
	if dc.Name == "" {
		return nil, errors.New("error prettifying json, field 'name' is empty")
	}

	d, err := json.MarshalIndent(dc, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("error prettifying json %v", err.Error())
	}

	return d, nil
}

// JSONToStruct unmarshals a given JSON byte slice into a DevContainer struct
func JSONToStruct(content []byte, dc *model.DevContainer) error {
	err := json.Unmarshal(content, &dc)
	if err != nil {
		return fmt.Errorf("error unmarshalling json %v", err.Error())
	}

	return nil
}

// ClearScreen clears the terminal screen based on the operating system
func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}
