package generator

import (
	"fmt"

	"github.com/lukasjarosch/genki-cli/internal"
	"github.com/lukasjarosch/genki-cli/internal/config"
)

const (
	ReadmeIdentifier = "Readme"
	MakefileIdentifier = "Makefile"
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
	default:
		return nil, fmt.Errorf("unknown generator identifier")
	}
}