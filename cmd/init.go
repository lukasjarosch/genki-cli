package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/lukasjarosch/genki-cli/internal/cli"
	"github.com/lukasjarosch/genki-cli/internal/config"
)
func init() {
	rootCmd.AddCommand(newCmd)
}

var newCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a genki project in the current working directory",
	Long:  `If you specify a type, genki-cli will prepare the project configuration for that scenario. Then you can simply generate a boilerplate.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := initProjectConfiguration(args[0])
		cli.Successf("genki project initialized, %s written", "genki.yaml")
		_ = cfg

		/*
			tplContext := internal.Context{
				Config: *cfg,
			}

			// render makefile
			tpl := template.NewTemplate(
				template.Name("makefile"),
				template.GoSource(false),
				template.UseFilesystem(internal.Templates),
				template.Path("makefile.go.tmpl"),
			)
			rendered, err := tpl.Render(tplContext)
			if err != nil {
				cli.Fatalf("failed to render template: %s", err)
			}
			fw := writer.NewFileWriter("./Makefile", writer.Overwrite(true))
			if err := fw.WriteFile(rendered); err != nil {
				cli.Errorf("failed to write path: %s: %s", "./Makefile", err)
				return
			}
		*/
	},
}

func initProjectConfiguration(serviceName string) *config.Configuration {
	var err error
	if serviceName == "" {
		serviceName, err = cli.Prompt("enter the service name (lowercase)")
		if err != nil {
			cli.Errorf("prompt failed: %s", err)
			os.Exit(2)
		}
	}
	pwd, _ := os.Getwd()

	cfg := config.NewProjectConfiguration(pwd)
	if cfg.Exists() {
		cli.Errorf("genki project already initialized")
		os.Exit(1)
	}

	cfg.Project.Name = serviceName
	if err := cfg.Validate(); err != nil {
		cli.Fatalf("configuration invalid: %s", err)
	}
	if err := cfg.CreateConfigFile(); err != nil {
		cli.Fatalf("failed to create config file: %s", err)
	}

	cfg.Project.Version = "1.0.0"

	cfg.Save()

	return cfg
}
