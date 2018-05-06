package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/7onetella/mvk/internal/mdcat"
	"github.com/spf13/cobra"
)

var mdcatCmd = &cobra.Command{
	Use:     "mdcat <filepath>",
	Short:   "Markdown automation",
	Long:    `Markdown automation`,
	Example: "mdcat ./README.md",
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.Open(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		content, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(mdcat.MarkDown(string(content)))
	},
}

func init() {
	rootCmd.AddCommand(mdcatCmd)
}
