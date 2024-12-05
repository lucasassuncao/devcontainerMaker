package repository

import (
	"fmt"
	"os"
)

func SaveToJSONFile(data []byte) error {
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

func ReadJSONFile() ([]byte, error) {
	file, err := os.ReadFile("devcontainer.json")
	if err != nil {
		return nil, fmt.Errorf("error reading file %v", err.Error())
	}

	return file, nil
}
