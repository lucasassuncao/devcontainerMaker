package utils

import (
	"devcontainerMaker/internal/model"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func PrettifyJSON(dc *model.DevContainer) ([]byte, error) {
	d, err := json.MarshalIndent(dc, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("error prettifying json %v", err.Error())
	}
	return d, nil
}

func JSONToStruct(content []byte, dc *model.DevContainer) error {
	err := json.Unmarshal(content, &dc)
	if err != nil {
		return fmt.Errorf("error unmarshalling json %v", err.Error())
	}

	return nil
}

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
