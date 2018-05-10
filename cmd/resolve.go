package cmd

import (
	"fmt"

	"github.com/7onetella/dex/internal/dockerw"
	"github.com/spf13/cobra"
)

var resolveCmd = &cobra.Command{
	Use:   "resolve <container-name>",
	Short: "Resolves container's ip by name",
	Long: `Resolves container's ip by name
	
	container-name     docker container name
	`,
	Example:            "resolve redis.local.io",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			cmd.Usage()
			return
		}

		host := args[0]

		resolved := dockerw.ResolveName(host)

		fmt.Println(resolved)
	},
}

func init() {
	rootCmd.AddCommand(resolveCmd)
}
