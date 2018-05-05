package cmd

import (
	"fmt"

	"github.com/7onetella/mvk/internal/dockerw"
	"github.com/7onetella/mvk/internal/net"
	"github.com/spf13/cobra"
)

var redisCmd = &cobra.Command{
	Use:   "redis <redishost>",
	Short: "Runs redis client",
	Long: `Runs redis client
	
	redishost     redis hostname
	`,
	Example:            "redis redis.mars.com",
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

		ok := net.IsTCPConnValid(host, port)
		if !ok {
			fmt.Println()
			fmt.Println(failed + " check network connection. i.e. vpn")
			return
		}

		env := []string{}
		dockerCmd := []string{"redis-cli", "-h", host, "-p", port}

		dockerw.StartDockerConsole("redis:3-alpine", env, dockerCmd)
	},
}

func init() {
	rootCmd.AddCommand(redisCmd)
}
