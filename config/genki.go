package config

import "github.com/spf13/viper"

const fileName = "genki"
const fileType = "yaml"

type GenkiConfig struct {
	rootPath string
}

// NewGenkiConfiguration initializes a new config struct. The rootPath is the
// path where the config file will be created/searched.
func NewGenkiConfiguration(rootPath string) Configurator {
	viper.AddConfigPath(rootPath)
	viper.SetConfigName(fileName)
	viper.SetConfigType(fileType)

	return &GenkiConfig{
		rootPath: rootPath,
	}
}

func (cfg *GenkiConfig) Save() error {
	return nil
}

func (cfg *GenkiConfig) Initialize() error {
	return nil
}


func (cfg *GenkiConfig) FileExists() bool {
	return false
}

func (cfg *GenkiConfig) CreateFile() error {
	return nil
}

