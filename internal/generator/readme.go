package generator

import (
	"fmt"

	"github.com/lukasjarosch/genki-cli/internal"
	"github.com/lukasjarosch/genki-cli/internal/cli"
	"github.com/lukasjarosch/genki-cli/internal/config"
	"github.com/lukasjarosch/genki-cli/pkg/template"
	"github.com/lukasjarosch/genki-cli/pkg/writer"
)

type readme struct {
	target       string
	templateName string
	templatePath string
}

func NewReadme() *readme {
	return &readme{
		target:       "./README.md",
		templateName: "readme",
		templatePath: "readme.go.tmpl",
	}
}

func NewReadmeFromConfig(cfg *config.Configuration) *readme {
	g := NewReadme()
	g.target = cfg.GetString("readme.path")

	return g
}

func (g *readme) Generate(ctx internal.Context) error {
	// render makefile
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

func (g *readme) Configure(cfg *config.Configuration) {
	cfg.Set("readme.path", g.target)
}

func (g *readme) Remove(cfg *config.Configuration) {
	cfg.Delete("readme")
}

