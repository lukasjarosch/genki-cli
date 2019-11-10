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
	viper.Set("genki", cfg)
	if err := viper.WriteConfig(); err != nil {
		cli.Errorf("unable to write config file: %s", err)
	}
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

func (cfg *Configuration) WithProjectType(typ string) error {
	if !cfg.isValidType(typ) {
		return fmt.Errorf("invalid project type: %s", typ)
	}
	switch typ {
	case "grpc-server":
		cfg.Grpc.Server = GrpcServer{
			Enabled:true,
			Name: cfg.Project.Service,
		}
		break
	case "http-server":
		cfg.Http.Server = HttpServer{
			Enabled:true,
		}
		break
	case "amqp-consumer":
		break
	case "http-gateway":
		break
	case "plain":
		break
	}

	return nil
}

func (cfg *Configuration) AbsPath() string {
	return path.Join(cfg.rootPath, fmt.Sprintf("%s.%s", FileName, FileType))
}

func (cfg *Configuration) Load() error {
	return viper.ReadInConfig()
}

func (cfg *Configuration) Validate() error {
	if !cfg.isValidServiceName() {
		return fmt.Errorf("invalid service name: %s", cfg.Project.Service)
	}
	if !cfg.isValidNamespace() {
		return fmt.Errorf("invalid namespace: %s", cfg.Project.Namespace)
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
	if cfg.Project.Service == "" {
		return false
	}
	cfg.Project.Service = strings.ToLower(cfg.Project.Service)
	return true
}

func (cfg *Configuration) isValidNamespace() bool {
	if cfg.Project.Namespace == "" {
		return false
	}
	cfg.Project.Namespace = strings.ToLower(cfg.Project.Namespace)
	return true
}
