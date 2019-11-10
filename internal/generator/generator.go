package generator

import (
	"fmt"

	"github.com/lukasjarosch/genki-cli/internal"
	"github.com/lukasjarosch/genki-cli/internal/cli"
	"github.com/lukasjarosch/genki-cli/internal/config"
	"github.com/lukasjarosch/genki-cli/pkg/template"
	"github.com/lukasjarosch/genki-cli/pkg/writer"
)

const (
	ReadmeIdentifier = "readme"
	MakefileIdentifier = "makefile"
	KubernetesIdentifer = "kubernetes"
)

type Generator interface {
	Generate(ctx internal.Context) error
	LoadConfiguration(cfg *config.Configuration)
	Configure(cfg *config.Configuration)
	Remove(cfg *config.Configuration)
}

func NewGenerator(identifier string) (Generator, error) {
	switch identifier {
	case ReadmeIdentifier:
		return newReadme(), nil
	case MakefileIdentifier:
		return newMakefile(), nil
	case KubernetesIdentifer:
		return newKubernetes(), nil
	default:
		return nil, fmt.Errorf("unknown generator identifier")
	}
}

type Target struct {
	Path string
	TemplateName string
	TemplatePath string
	GoSource bool
	Overwrite bool
	AppendWrite bool
}

type baseGenerator struct {
	Targets map[string]Target
}

func (g *baseGenerator) Generate(ctx internal.Context) error {
	for _, target := range g.Targets {
		tpl := template.NewTemplate(
			template.Name(target.TemplateName),
			template.GoSource(target.GoSource),
			template.UseFilesystem(internal.Templates),
			template.Path(target.TemplatePath))

		rendered, err := tpl.Render(ctx)
		if err != nil {
			cli.Fatalf("failed to render template: %s", err)
		}
		fw := writer.NewFileWriter(target.Path, writer.Overwrite(true))
		if err := fw.WriteFile(rendered); err != nil {
			return fmt.Errorf("failed to write: %s: %s", target.Path, err)
		}
		cli.Successf("rendered '%s' into: %s", target.TemplateName, target.Path)
	}
	return nil
}