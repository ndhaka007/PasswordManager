package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

var envWithModeBasedConfig = []string{"stage", "prod"}

// loadConfigFromFile: load/overwrite config values from given basepath, filename and env
func LoadConfigFromFile(
	filePath string,
	basePath string,
	filename string,
	configStruct interface{},
	env string,
	isRequired bool) {
	path := getFilePath(filePath, basePath, filename, env)

	_, err := os.Stat(path)

	if err != nil {
		if isRequired {
			//logger.Get(nil).WithError(err).Panic("config file not found")
		} else {
			//logger.Get(nil).WithError(err).Info("config file not found")

			return
		}
	}

	if _, err := toml.DecodeFile(path, configStruct); err != nil {
		//logger.Get(nil).WithError(err).Panic("Invalid data type in configuration")
	}
}

// getFilePath: gives the file path based on the environment provided
// file path will be relative to the application and determined by basePath
// file will be chosen based on env and mode. if the mode based config is available for an env then will be loaded
func getFilePath(filePath string, basePath string, fileName string, env string) string {
	envFile := env

	if env != "" {
		fileName = fmt.Sprintf(fileName, envFile)
	}

	path := fmt.Sprintf(filePath, basePath, fileName)

	return path
}
