package generator

import (
	"github.com/lukasjarosch/genki-cli/internal/config"
)

type Makefile struct {
	baseGenerator
}

func newMakefile() *Makefile {
	targets := map[string]Target{
		"makefile": {
			Path:         "./Makefile",
			TemplateName: "makefile",
			TemplatePath: "/makefile.go.tmpl",
			GoSource:     false,
			Overwrite:    true,
			AppendWrite:  false,
		},
	}

	return &Makefile{
		baseGenerator{Targets: targets},
	}
}

func (g *Makefile) LoadConfiguration(cfg *config.Configuration) {
	newCfg := g.Targets["makefile"]
	newCfg.Path = cfg.GetString("rakefile.path")

	g.Targets["makefile"] = newCfg
}

func (g *Makefile) Configure(cfg *config.Configuration) {
	cfg.Set("makefile.path", g.Targets["makefile"].Path)
}

func (g *Makefile) Remove(cfg *config.Configuration) {
	cfg.Delete("makefile")
}

