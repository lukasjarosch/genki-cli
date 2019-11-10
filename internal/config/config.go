package config

import (
	"fmt"
	"strings"
)

var ValidProjectTypes = []string{"grpc-server", "http-server", "amqp-consumer", "http-gateway", "plain"}

type ProjectConfiguration struct {
	ProjectType string
	Namespace   string
	ServiceName string
}

func NewProjectConfiguration(projectType, serviceName, namespace string) *ProjectConfiguration {
	cfg := &ProjectConfiguration{
		ProjectType: projectType,
		Namespace:   namespace,
		ServiceName: serviceName,
	}

	return cfg
}

func (cfg *ProjectConfiguration) Validate() error {
	if !cfg.isValidType(cfg.ProjectType) {
		return fmt.Errorf("invalid project type: %s", cfg.ProjectType)
	}
	cfg.ServiceName = cfg.ensureLowercase(cfg.ServiceName)
	cfg.Namespace = cfg.ensureLowercase(cfg.Namespace)
	return nil
}

func (cfg *ProjectConfiguration) isValidType(typ string) bool {
	for _, t := range ValidProjectTypes {
		if t == typ {
			return true
		}
	}
	return false
}

func (cfg *ProjectConfiguration) ensureLowercase(str string) string {
	return strings.ToLower(str)
}
