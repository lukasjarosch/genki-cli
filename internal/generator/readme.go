package generator

import (
	"fmt"

	"github.com/lukasjarosch/genki-cli/internal"
	"github.com/lukasjarosch/genki-cli/internal/cli"
	"github.com/lukasjarosch/genki-cli/internal/config"
	"github.com/lukasjarosch/genki-cli/pkg/template"
	"github.com/lukasjarosch/genki-cli/pkg/writer"
)

type Readme struct {
	target       string
	templateName string
	templatePath string
}

func newReadme() *Readme {
	return &Readme{
		target:       "./README.md",
		templateName: "Readme",
		templatePath: "Readme.go.tmpl",
	}
}

func (g *Readme) Generate(ctx internal.Context) error {
	// render Makefile
	tpl := template.NewTemplate(
		template.Name(g.templateName),
		template.GoSource(false),
		template.UseFilesystem(internal.Templates),
		template.Path(g.templatePath),
	)
	rendered, err := tpl.Render(ctx)
	if err != nil {
		cli.Fatalf("failed to render template: %s", err)
	}
	fw := writer.NewFileWriter(g.target, writer.Overwrite(true))
	if err := fw.WriteFile(rendered); err != nil {
		return fmt.Errorf("failed to write: %s: %s", g.target, err)
	}
	cli.Successf("rendered '%s' into: %s", g.templateName, g.target)

	return nil
}

func (g *Readme) Configure(cfg *config.Configuration) {
	cfg.Set("Readme.path", g.target)
}

func (g *Readme) Remove(cfg *config.Configuration) {
	cfg.Delete("Readme")
}

func (g *Readme) LoadConfiguration(cfg *config.Configuration) {
	g.target = cfg.GetString("Readme.path")
}
