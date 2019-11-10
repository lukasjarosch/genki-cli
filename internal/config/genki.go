package config

import "fmt"

var ValidProjectTypes = []string{"grpc-server", "http-server", "amqp-consumer", "http-gateway", "plain"}

const FileName = "genki"
const FileType = "yaml"

type Configuration struct {
	rootPath string
	Project  Project  `yaml:"project"`
}

type Project struct {
	Name       string   `yaml:"name"`
	Version    string   `yaml:"version"`
	Generators []string `yaml:"generators"`
}

func (p *Project) GeneratorEnabled(identifier string) bool {
	for _, gen := range p.Generators {
		if gen == identifier {
			return true
		}
	}
	return false
}

func (p *Project) RemoveGenerator(identifier string) {
	var newList []string
	for _, gen := range p.Generators {
		if gen == identifier {
			continue
		}
		newList = append(newList, gen)
	}
	p.Generators = newList
}

func (p *Project) AddGenerator(identifier string) error {
	if p.GeneratorEnabled(identifier) {
		return fmt.Errorf("generator %s already enabled", identifier)
	}
	p.Generators = append(p.Generators, identifier)
	return nil
}

type Grpc struct {
	Server GrpcServer
}

type GrpcServer struct {
	Enabled bool
	Name    string
}

type Http struct {
	Server HttpServer
}

type HttpServer struct {
	Enabled bool
	Name    string
}
