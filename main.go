package main

import (
	"log"
)

func main() {
	log.Print("starting envdocksec ...")
	cfg := getConfig()
	handleOutputDir(cfg)
	inputVars := getInput(cfg)

	createSecretFiles(inputVars, cfg)
	manifest := createOutputManifest(inputVars, cfg)

	writeOutputManifest(manifest, cfg)
}
