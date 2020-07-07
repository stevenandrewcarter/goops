package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goops",
	Short: "Configuration Management and Convergance",
	Long:  `Goops is a tool for applying configuration management concepts and convergance to a given host.`,
}

// Execute is the Entry point for Cobra
func Execute() error {
	return rootCmd.Execute()
}
