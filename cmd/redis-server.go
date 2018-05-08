package cmd

import (
	"strings"

	"github.com/7onetella/mvk/internal/execw"

	"github.com/spf13/cobra"
)

var redisServerCmd = &cobra.Command{
	Use:                "redis-server <redishost>",
	Short:              "Runs redis server locally",
	Long:               `Runs redis server locally`,
	Example:            "redis-server",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {

		command := "docker run --name redis --rm -d -p 6379:6379 redis:3-alpine"
		dockerCmd := strings.Split(command, " ")

		execw.Exec(dockerCmd)

	},
}

func init() {
	rootCmd.AddCommand(redisServerCmd)
}
