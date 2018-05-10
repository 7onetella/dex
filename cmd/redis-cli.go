package cmd

import (
	"fmt"
	"strings"

	"github.com/7onetella/dex/internal/dockerw"
	"github.com/spf13/cobra"
)

var redisCmd = &cobra.Command{
	Use:   "redis-cli <redishost>",
	Short: "Runs redis client",
	Long: `Runs redis client
	
	redishost     redis hostname
	`,
	Example:            "redis-cli redis.local.io",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			cmd.Usage()
			return
		}

		host := args[0]
		port := "6379"

		if len(args) > 1 {
			port = args[1]
		}

		resolved := dockerw.ResolveName(host)

		dockerCmd := []string{"redis-cli", "-h", resolved, "-p", port}

		fmt.Println(strings.Join(dockerCmd, " "))
		dockerw.StartDockerConsole("redis:3-alpine", []string{}, dockerCmd)
	},
}

func init() {
	rootCmd.AddCommand(redisCmd)
}
