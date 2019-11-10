package config

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/lukasjarosch/genki-cli/internal/cli"
)

func NewProjectConfiguration(rootPath string) *Configuration {
	cfg := &Configuration{
		rootPath: rootPath,
	}

	viper.AddConfigPath(cfg.rootPath)
	viper.SetConfigName(FileName)
	viper.SetConfigType(FileType)

	return cfg
}

func (cfg *Configuration) Save() {
	viper.Set("genki", *cfg)
	if err := viper.WriteConfig(); err != nil {
		cli.Errorf("unable to write config file: %s", err)
	}
}

func (cfg *Configuration) Set(key string, value interface{}) {
	viper.Set(key, value)
}

func (cfg *Configuration) GetString(key string) string {
	return viper.GetString(key)
}

func (cfg *Configuration) Delete(key string) {
	delete(viper.Get(key).(map[string]interface{}), key)
}

func (cfg *Configuration) Exists() bool {
	if _, err := os.Stat(cfg.AbsPath()); os.IsNotExist(err) {
		return false
	}
	return true
}

func (cfg *Configuration) CreateConfigFile() error {
	if !cfg.Exists() {
		if _, err := os.Create(cfg.AbsPath()); err != nil {
			return errors.Wrap(err, "failed to create config file")
		}
	}
	return nil
}

func (cfg *Configuration) AbsPath() string {
	return path.Join(cfg.rootPath, fmt.Sprintf("%s.%s", FileName, FileType))
}

func (cfg *Configuration) Load() error {
	if !cfg.Exists() {
		return fmt.Errorf("missing config file")
	}
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config file")
	}
	return viper.UnmarshalKey("genki", &cfg)
}

func (cfg *Configuration) Validate() error {
	if !cfg.isValidServiceName() {
		return fmt.Errorf("invalid service name: %s", cfg.Project.Name)
	}
	return nil
}

func (cfg *Configuration) isValidType(typ string) bool {
	for _, t := range ValidProjectTypes {
		if t == typ {
			return true
		}
	}
	return false
}

func (cfg *Configuration) isValidServiceName() bool {
	if cfg.Project.Name == "" {
		return false
	}
	cfg.Project.Name = strings.ToLower(cfg.Project.Name)
	return true
}


