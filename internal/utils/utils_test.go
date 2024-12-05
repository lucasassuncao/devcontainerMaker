package utils

import (
	"devcontainerMaker/internal/model"
	"testing"
)

func TestPrettifyJSON(t *testing.T) {
	t.Run("Success to Prettify JSON", func(t *testing.T) {
		dc := model.NewDevContainer().
			WithName().
			WithBuildDockerFile().
			WithShutdownAction()

		dc.SetName("TestDevContainer")
		dc.SetBuildDockerfile("Dockerfile")
		dc.SetShutdownAction("stopContainer")

		_, err := PrettifyDevContainerJSON(dc)
		if err != nil {
			t.Errorf("PrettifyDevContainerJSON failed with error: %v", err.Error())
		}
	})

	t.Run("Fail to Prettify JSON", func(t *testing.T) {
		dc := model.NewDevContainer().
			WithName().
			WithBuildDockerFile().
			WithShutdownAction()

		_, err := PrettifyDevContainerJSON(dc)
		if err == nil {
			t.Error("Expected PrettifyDevContainerJSON to fail, but it succeeded")
		}
	})
}
