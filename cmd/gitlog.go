package cmd

import (
	"fmt"
	"strings"

	"github.com/7onetella/dex/internal/execw"
	"github.com/spf13/cobra"
)

var gitLogCmd = &cobra.Command{
	Use:     "gitlog",
	Short:   "Git log with flattned view and formatting",
	Long:    `Git log with flattned view and formatting`,
	Example: "gitlog",
	Run: func(cmd *cobra.Command, args []string) {
		cmds := []string{"git",
			"log",
			"--graph",
			"--all",
			"--color",
			"--date=short",
			"-40",
			`--pretty=format:"%C(yellow)%h%x20%C(white)%cd%C(green)%d%C(reset)%x20%s%x20%C(bold)(%an)%Creset"`}

		fmt.Println(strings.Join(cmds, " "))
		output, err := execw.Execute(cmds)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(gitLogCmd)
}
