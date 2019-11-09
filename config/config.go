package config

type Configurator interface {
	// Save the config to file
	Save() error
	// FileExists returns true if the config file exists in the current working directory
	FileExists() bool
	// CreateFile will create a new configuration file
	CreateFile() error
	// Initialize will load a config if it exists
	Initialize() error
}
