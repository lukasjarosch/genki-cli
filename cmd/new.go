package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"

	"github.com/lukasjarosch/genki-cli/internal/cli"
	"github.com/lukasjarosch/genki-cli/internal/config"
)

var (
	projectType string
	serviceName string
	namespace string
)

func init() {
	rootCmd.AddCommand(newCmd)
	flag.StringVarP(&projectType, "type", "t", "", fmt.Sprintf("the type of the project you're about to create (%s)", config.ValidProjectTypes))
	flag.StringVarP(&serviceName, "service", "s", "", "the name of the service to create (lowercase)")
	flag.StringVarP(&namespace, "namespace", "n", "", "the name of the namespace in which the service resices (lowercase)")
	flag.Parse()
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new genki-powered project",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		log := cli.NewCliLogger()

		if projectType == "" {
			projectType, err = cli.Select("what type of project do you want to create?", config.ValidProjectTypes...)
			if err != nil {
				log.Errorf("prompt failed: %s", err)
				return
			}
		}
		if serviceName == "" {
			serviceName, err = cli.Prompt("enter the service name (lowercase)")
			if err != nil {
			    log.Errorf("prompt failed: %s", err)
			    return
			}
		}
		if namespace == "" {
			namespace, err = cli.Prompt("enter the service namespace (lowercase)")
			if err != nil {
				log.Errorf("prompt failed: %s", err)
				return
			}
		}

		cfg := config.NewProjectConfiguration(projectType, serviceName, namespace)
		if err := cfg.Validate(); err != nil {
		   	log.Fatalf("configuration invalid: %s", err)
		}

	},
}

