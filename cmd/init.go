package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lukasjarosch/genki-cli/internal/cli"
	"github.com/lukasjarosch/genki-cli/internal/config"
)

var (
	projectType    string
	serviceName    string
	namespace      string
	dockerRegistry string
)

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().StringVarP(&projectType, "type", "t", "", fmt.Sprintf("specify a type to prepare the config for the desired use-case. Avaliable: (%s)", config.ValidProjectTypes))
	newCmd.Flags().StringVarP(&serviceName, "service", "s", "", "the name of the service to create (lowercase)")
	newCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "the name of the namespace in which the service resides (lowercase)")
	newCmd.Flags().StringVarP(&dockerRegistry, "docker-registry", "d", "docker.io", "specify which docker-registry to use")
	if err := newCmd.Flags().Parse(newCmd.Flags().Args()); err != nil {
		cli.Fatalf("failed to parse flags")
	}
}

var newCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a genki project in the current working directory",
	Long:  `If you specify a type, genki-cli will prepare the project configuration for that scenario. Then you can simply generate a boilerplate.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := initProjectConfiguration()
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

func initProjectConfiguration() *config.Configuration {
	var err error
	if projectType == "" {
		projectType, err = cli.Select("what type of project do you want to create?", config.ValidProjectTypes...)
		if err != nil {
			cli.Errorf("prompt failed: %s", err)
			os.Exit(2)
		}
	}
	if serviceName == "" {
		serviceName, err = cli.Prompt("enter the service name (lowercase)")
		if err != nil {
			cli.Errorf("prompt failed: %s", err)
			os.Exit(2)
		}
	}
	if namespace == "" {
		namespace, err = cli.Prompt("enter the service namespace (lowercase)")
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

	cfg.Project.Service = serviceName
	cfg.Project.Namespace = namespace
	cfg.Project.DockerRegistry = dockerRegistry
	if err := cfg.Validate(); err != nil {
		cli.Fatalf("configuration invalid: %s", err)
	}
	if err := cfg.WithProjectType(projectType); err != nil {
		cli.Fatalf("unable to create config with type: %s: %s", projectType, err)
	}
	if err := cfg.CreateConfigFile(); err != nil {
		cli.Fatalf("failed to create config file: %s", err)
	}
	cfg.Save()

	return cfg
}
