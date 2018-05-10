package cmd

import (
	"github.com/7onetella/dex/internal/dockerw"
	"github.com/spf13/cobra"
)

var siegeCmd = &cobra.Command{
	Use:                "siege <siege parameters>",
	Short:              "Runs siege",
	Long:               `Runs siege`,
	Example:            "siege -d1 -r10 -c25 example.com",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			cmd.Usage()
			return
		}

		params := args[0:]

		env := []string{}
		dockerCmd := []string{"siege"}
		for _, p := range params {
			dockerCmd = append(dockerCmd, p)
		}

		dockerw.StartDockerConsole("yokogawa/siege", env, dockerCmd)
	},
}

func init() {
	rootCmd.AddCommand(siegeCmd)
}
