package cmd

import (
	"github.com/spf13/cobra"
)

var scriptCmd = &cobra.Command{
	Use:   "script [command]",
	Short: "Script automation",
	Long:  `Script automation`,
}

func init() {
	rootCmd.AddCommand(scriptCmd)
}
