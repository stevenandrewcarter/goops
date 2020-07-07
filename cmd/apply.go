package cmd

import (
	"log"

	"github.com/hashicorp/hcl/v2/hclsimple"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(applyCmd)
}

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Executes the Given configuration.",
	Long:  `TODO: Long Description`,
	Run: func(cmd *cobra.Command, args []string) {
		var config config
		err := hclsimple.DecodeFile("config.hcl", nil, &config)
		if err != nil {
			log.Fatalf("Failed to load configuration: %s", err)
		}
		log.Printf("Configuration is %#v", config)
	},
}

type config struct {
	LogLevel string `hcl:"log_level"`
}
