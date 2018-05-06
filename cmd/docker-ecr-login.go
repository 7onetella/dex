package cmd

import (
	"fmt"

	"strings"

	"github.com/7onetella/mvk/internal/execw"
	"github.com/spf13/cobra"
)

var dockerLoginCmd = &cobra.Command{
	Use:     "ecr-login",
	Short:   "Docker ecr login",
	Long:    `Docker ect login`,
	Example: "ecr-login",
	Run: func(cmd *cobra.Command, args []string) {
		ecrLogin := "aws ecr get-login --no-include-email"
		fmt.Println(ecrLogin)
		output, err := execw.Execute(strings.Split(ecrLogin, " "))
		if err != nil {
			return
		}
		fmt.Println()

		fmt.Println(output)
		cmds := strings.Split(strings.TrimSpace(output), " ")

		output, err = execw.Execute(cmds)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print(output)
	},
}

func init() {
	dockerCmd.AddCommand(dockerLoginCmd)
}
