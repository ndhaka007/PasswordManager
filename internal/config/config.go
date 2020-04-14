// Package config will manage all application level configurations
// config file will be taken based on the application environment
// all the configuration available in vault file will be overwritten
// this will be immutable as it always provides the value of the struct
package config

const (
	// FilePath - relative path to the config directory
	FilePath = "%s/config/%s"

	// DefaultFilename - Filename format of default config file
	DefaultFilename = "env.default.toml"

	// EnvFilename - Filename format of env specific config file
	EnvFilename = "env.%s.toml"
)

var (
	// config : this will hold all the application configuration
	config appConfig
)

// appConfig global configuration struct definition
type appConfig struct {
	Application application `toml:"application"`
	Database    Database    `toml:"database"`
}

// LoadConfig will load the configuration available in the cnf directory available in basePath
// config file will takes based on the env provided
// vault file content will override the config available
func LoadConfig(basePath string, env string) {
	//logger.Get(nil).Info("Loading environment for " + env)

	config.Application.Environment = env

	// reading config based on default environment
	LoadConfigFromFile(FilePath, basePath, DefaultFilename, &config, "", true)

	// reading env file and override config values; if env file exists
	LoadConfigFromFile(FilePath, basePath, EnvFilename, &config, env, false)

	err := LoadEnvironmentVariables(&config)
	if err != nil {
		//logger.Get(nil).WithError(err).Panic("Having trouble loading environment variables...")
	}
}

// GetConfig : will give the struct as value so that the actual config doesn't get tampered
func GetConfig() appConfig {
	return config
}
