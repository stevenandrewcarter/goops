package cmd

import (
	"bytes"
	"log"
	"os/exec"

	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(applyCmd)
}

func executeCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	log.Print(out.String())
	return err
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
		for i := range config.Command {
			err = executeCommand(config.Command[i].Name, config.Command[i].Args...)
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

type config struct {
	LogLevel string     `hcl:"log_level"`
	Command  []*command `hcl:"command,block"`
}

type command struct {
	Name string   `hcl:",label"`
	Args []string `hcl:"args,optional"`
}
