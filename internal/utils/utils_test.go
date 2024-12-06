package utils

import (
	"devcontainerMaker/internal/model"
	"testing"
)

/*func TestPrettifyJSON(t *testing.T) {
	t.Run("Success to Prettify JSON", func(t *testing.T) {
		dc := model.NewDevContainer().
			withName().
			WithBuild().
			WithShutdownAction()

		dc.SetName()
		dc.SetBuild()
		dc.SetShutdownAction()

		_, err := PrettifyDevContainerJSON(dc)
		if err != nil {
			t.Errorf("PrettifyDevContainerJSON failed with error: %v", err.Error())
		}
	})

	t.Run("Fail to Prettify JSON", func(t *testing.T) {
		dc := model.NewDevContainer().
			withName().
			WithBuild().
			WithShutdownAction()

		_, err := PrettifyDevContainerJSON(dc)
		if err == nil {
			t.Error("Expected PrettifyDevContainerJSON to fail, but it succeeded")
		}
	})
}*/

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
