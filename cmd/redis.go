package cmd

import (
	"strings"

	"github.com/7onetella/mvk/internal/execw"

	"github.com/spf13/cobra"
)

var redisServerCmd = &cobra.Command{
	Use:                "redis",
	Short:              "Runs redis server locally",
	Long:               `Runs redis server locally`,
	Example:            "redis",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {

		// command := "docker run --name redis-server --net macvlan --rm -d -p 6379:6379 redis:3-alpine"
		command := "docker run --name redis.local.io --rm -d -p 6379:6379 redis:3-alpine"
		dockerCmd := strings.Split(command, " ")

		//dockerw.StartDockerConsolePre()

		execw.Exec(dockerCmd)

	},
}

func init() {
	rootCmd.AddCommand(redisServerCmd)
}
