package repository

import (
	"fmt"
	"os"
)

// SaveDevContainerJSONFile creates a file named devcontainer.json and writes the provided JSON data to it
func SaveDevContainerJSONFile(data []byte) error {
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

// ReadDevContainerJSONFile reads the contents of the devcontainer.json file and returns it as a byte slice
func ReadDevContainerJSONFile() ([]byte, error) {
	file, err := os.ReadFile("devcontainer.json")
	if err != nil {
		return nil, fmt.Errorf("error reading file %v", err.Error())
	}

	return file, nil
}
