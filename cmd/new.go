package cmd

import (
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"

	"github.com/lukasjarosch/genki-cli/internal/cli"
)

var projectTemplate = "none"

func init() {
	rootCmd.AddCommand(newCmd)
	flag.StringVar(&projectTemplate, "type", "none", "the type of the project you're about to create")
	flag.Parse()
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new genki-powered project",
	Run: func(cmd *cobra.Command, args []string) {
		log := cli.NewCliLogger()

		var err error
		projectTemplate, err = cli.Select("what type of project do you want to create?", "grpc-server", "http-server", "amqp-consumer", "http-gateway", "none")
		if err != nil {
			log.Error("invalid template: %s", projectTemplate)
		}
		log.Infof("project template selected: %s", projectTemplate)
	},
}
