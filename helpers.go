package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/AckeeDevOps/envdocksec/config"
)

// get configuration from env variables
func getConfig() *config.PluginConfig {
	cfg, err := config.Create()
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}

// parse input file in dotenv format
func getInput(c *config.PluginConfig) map[string]interface{} {

	log.Printf("Getting input data from %s ...", c.InputFile)

	inputFile, err := ioutil.ReadFile(c.InputFile)
	if err != nil {
		log.Fatalf("Could not open input file: %s", err)
	}

	inputVars := map[string]interface{}{}
	err = json.Unmarshal(inputFile, &inputVars)
	if err != nil {
		log.Fatalf("Could not unmarshal input file: %s", err)
	}

	return inputVars
}

func handleOutputDir(c *config.PluginConfig) {
	if !dirExist(c.OuputDirectory) && c.CreateOutputDirectory {
		createDir(c.OuputDirectory)
	} else if dirExist(c.OuputDirectory) {
		log.Printf("Output directory %s already exists", c.OuputDirectory)
	} else {
		log.Fatalf("Output directory %s does not exist", c.OuputDirectory)
	}
}

func dirExist(name string) bool {
	fi, err := os.Stat(name)
	if err != nil {
		return false
	}
	return fi.Mode().IsDir()
}

func createDir(name string) {
	log.Printf("creating output directory %s", name)
	err := os.Mkdir(name, os.ModePerm)
	// handle exceptions such as regular file with the same name
	if err != nil {
		log.Fatalf("Could not create directory %s: %s", name, err)
	}
}

// create files with secrets values, filename = KEY
func createSecretFiles(vars map[string]interface{}, c *config.PluginConfig) {
	for key, value := range vars {

		filename := fmt.Sprintf("%s/%s", c.OuputDirectory, key)
		text := []byte(fmt.Sprintf("%v", value))

		log.Printf("Creating %s ...", filename)

		err := ioutil.WriteFile(filename, text, os.ModePerm)
		if err != nil {
			log.Fatalf("Could not create %s: %s", filename, err)
		}
	}
}

func createOutputManifest(vars map[string]interface{}, c *config.PluginConfig) []string {
	mapping := []string{}
	for key := range vars {
		entry := fmt.Sprintf("%s/%s:%s/%s", c.OuputDirectory, key, c.DockerTargetDirectory, key)
		mapping = append(mapping, entry)
	}
	return mapping
}

func writeOutputManifest(manifest []string, c *config.PluginConfig) {
	res, err := json.Marshal(manifest)
	if err != nil {
		log.Fatalf("Could not marshal JSON: %s", err)
	}

	log.Printf("Creating output manifest %s", c.OutputManifest)
	err = ioutil.WriteFile(c.OutputManifest, res, os.ModePerm)
	if err != nil {
		log.Fatalf("Could not write manifest to file: %s", err)
	}
}
