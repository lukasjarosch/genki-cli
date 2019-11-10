package generator

import (
	"github.com/lukasjarosch/genki-cli/internal/config"
)

type Kubernetes struct {
	baseGenerator
}

func newKubernetes() *Kubernetes {
	targets := map[string]Target{
		"service": {
			Path:         "./k8s/service.yaml",
			TemplateName: "k8s-service",
			TemplatePath: "/k8s/service.go.tmpl",
			GoSource:     false,
			Overwrite:    true,
			AppendWrite:  false,
		},
	}

	return &Kubernetes{
		baseGenerator{Targets: targets},
	}
}

func (g *Kubernetes) Configure(cfg *config.Configuration) {
	cfg.Set("kubernetes.service.path", g.Targets["service"].Path)
}

func (g *Kubernetes) Remove(cfg *config.Configuration) {
	cfg.Delete("kubernetes")
}

func (g *Kubernetes) LoadConfiguration(cfg *config.Configuration) {
	newCfg := g.Targets["service"]
	newCfg.Path = cfg.GetString("kubernetes.service.path")

	g.Targets["service"] = newCfg
}
