package utils

import (
	"devcontainerMaker/internal/model"
	"testing"
)

// TestPrettifyDevContainerJSON tests the functionality of PrettifyDevContainerJSON.
func TestPrettifyDevContainerJSON(t *testing.T) {
	// Success Test: Valid DevContainer
	t.Run("Success to Prettify JSON", func(t *testing.T) {
		dc, _ := model.NewDevContainer().Initialize("dockerfile")

		dc.SetName("TestDevContainer")
		dc.SetBuildDockerfile("Dockerfile")
		dc.SetBuildContext(".")
		dc.SetShutdownAction("none")

		_, err := PrettifyDevContainerJSON(dc)
		if err != nil {
			t.Errorf("Expected success but got error: %v", err.Error())
		}
	})

	// Failure Test: Missing Required Fields
	t.Run("Fail to Prettify JSON", func(t *testing.T) {
		dc, _ := model.NewDevContainer().Initialize("dockerfile")

		// Intentionally not setting required fields

		_, err := PrettifyDevContainerJSON(dc)
		if err == nil {
			t.Error("Expected error but got success")
		}
	})
}

func TestJSONToStruct(t *testing.T) {
	type args struct {
		content []byte
		dc      *model.DevContainer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid JSON",
			args: args{
				content: []byte(`{"name": "my-dev-container", "image": "ubuntu:20.04"}`),
				dc:      &model.DevContainer{},
			},
			wantErr: false,
		},
		{
			name: "Invalid JSON",
			args: args{
				content: []byte(`{"name": "my-dev-container", "image": "ubuntu:20.04"`), // Missing closing brace
				dc:      &model.DevContainer{},
			},
			wantErr: true,
		},
		{
			name: "Empty JSON",
			args: args{
				content: []byte(``),
				dc:      &model.DevContainer{},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := JSONToStruct(tt.args.content, tt.args.dc); (err != nil) != tt.wantErr {
				t.Errorf("JSONToStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
