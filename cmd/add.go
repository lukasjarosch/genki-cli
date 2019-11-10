/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/lukasjarosch/genki-cli/internal"
	"github.com/lukasjarosch/genki-cli/internal/cli"
	"github.com/lukasjarosch/genki-cli/internal/config"
	"github.com/lukasjarosch/genki-cli/internal/generator"
)

var cfg *config.Configuration

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <generator>",
	Short: "Add generators to your project",
	Args: cobra.MinimumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		cfg = config.NewProjectConfiguration(wd)
		if err := cfg.Load(); err != nil {
			cli.Fatalf("failed to load configuration: %s", err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		generatorName := args[0]
		tplContext := internal.Context{
			Config: *cfg,
		}

		// resolve generator by name or fail
		gen, err := generator.NewGenerator(generatorName)
		if err != nil {
		    cli.Errorf("failed to add generator '%s': %s", generatorName, err)
		    os.Exit(1)
		}

		// add to project config
		if err := cfg.Project.AddGenerator(generatorName); err != nil {
			cli.Warningf("generator '%s' is already enabled", generatorName)
			return
		}
		cli.Infof("generator '%s' enabled", generatorName)

		// configure and run generator
		gen.Configure(cfg)
		if err := gen.Generate(tplContext); err != nil {
			cli.Errorf("%s generator failed: %s", generatorName, err)
			return
		}
		cfg.Save()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
