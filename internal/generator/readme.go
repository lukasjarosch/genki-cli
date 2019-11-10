package generator

import (
	"github.com/lukasjarosch/genki-cli/internal/config"
)

type Readme struct {
	baseGenerator
}

func newReadme() *Readme {
	targets := map[string]Target{
		"readme": {
			Path:         "./README.md",
			TemplateName: "readme",
			TemplatePath: "/readme.go.tmpl",
			GoSource:     false,
			Overwrite:    true,
			AppendWrite:  false,
		},
	}
	return &Readme{
		baseGenerator{Targets:targets},
	}
}

func (g *Readme) Configure(cfg *config.Configuration) {
	cfg.Set("readme.path", g.Targets["readme"].Path)
}

func (g *Readme) Remove(cfg *config.Configuration) {
	cfg.Delete("readme")
}

func (g *Readme) LoadConfiguration(cfg *config.Configuration) {
	newCfg := g.Targets["readme"]
	newCfg.Path = cfg.GetString("readme.path")
}
