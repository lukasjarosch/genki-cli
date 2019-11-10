package config

var ValidProjectTypes = []string{"grpc-server", "http-server", "amqp-consumer", "http-gateway", "plain"}

const FileName = "genki"
const FileType = "yaml"

type Configuration struct {
	rootPath string
	Project  Project  `yaml:"project"`
	Grpc     Grpc     `yaml:"grpc"`
	Http     Http     `yaml:"http"`
}

type Project struct {
	Namespace      string `yaml:"namespace"`
	Service        string `yaml:"service"`
	DockerRegistry string `yaml:"dockerRegistry"`
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
