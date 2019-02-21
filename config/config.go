package config

import (
	"os"
	"strconv"
)

type PluginConfig struct {
	InputFile             string
	OuputDirectory        string
	CreateOutputDirectory bool
	OutputManifest        string
	DockerTargetDirectory string
}

func Create() (*PluginConfig, error) {
	p := PluginConfig{}

	inputFile := os.Getenv("ENVDOCKSEC_INPUT_FILE")
	outputDirectory := os.Getenv("ENVDOCKSEC_OUTPUT_DIRECTORY")
	outputManifest := os.Getenv("ENVDOCKSEC_OUTPUT_MANIFEST")
	dockerTargetDirectory := os.Getenv("ENVDOCKSEC_DOCKER_TARGET_DIRECTORY")

	var createOutputDirectory bool
	createOutputDirectory, err := strconv.ParseBool(os.Getenv("ENVDOCKSEC_CREATE_OUTPUT_DIRECTORY"))
	if err != nil {
		createOutputDirectory = false
	}

	p.InputFile = inputFile
	p.OuputDirectory = outputDirectory
	p.CreateOutputDirectory = createOutputDirectory
	p.OutputManifest = outputManifest
	p.DockerTargetDirectory = dockerTargetDirectory

	return &p, nil
}
