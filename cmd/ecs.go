package cmd

import (
	"github.com/spf13/cobra"
)

var ecsCmd = &cobra.Command{
	Use:   "ecs [command]",
	Short: "ECS automation",
	Long:  `ECS automation`,
}

func init() {
	rootCmd.AddCommand(ecsCmd)
}
