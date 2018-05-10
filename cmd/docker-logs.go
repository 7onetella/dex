package cmd

import (
	"fmt"
	"strings"

	"github.com/7onetella/dex/internal/execw"
	"github.com/spf13/cobra"
)

var dockerLogsCmd = &cobra.Command{
	Use:     "logs [name]",
	Short:   "Docker logs",
	Long:    `Docker logs`,
	Example: "logs foo",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Usage()
			return
		}

		name := args[0]

		cmds := "docker logs -t --follow " + name
		fmt.Println(cmds)

		execw.Exec(strings.Split(cmds, " "))
	},
}

func init() {
	dockerCmd.AddCommand(dockerLogsCmd)
}
