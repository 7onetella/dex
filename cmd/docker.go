package cmd

import (
	"github.com/spf13/cobra"
)

var dockerCmd = &cobra.Command{
	Use:   "docker [command]",
	Short: "Docker automation",
	Long:  `Docker automation`,
}

func init() {
	rootCmd.AddCommand(dockerCmd)
}
