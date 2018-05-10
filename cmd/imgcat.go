package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/7onetella/dex/internal/img"
	"github.com/spf13/cobra"
)

var imageCatCmd = &cobra.Command{
	Use:     "imgcat [file_path|url]",
	Short:   "Renders image from local path on iTerm terminal",
	Long:    `Renders image from local path on iTerm terminal`,
	Example: "imgcat ./foo.jpg",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			cmd.Usage()
			return
		}

		s := args[0]

		term, ok := os.LookupEnv("TERM_PROGRAM")

		if ok && term == "iTerm.app" {
			if !strings.HasPrefix(s, "http") {
				img.Cat(s, "0", "0")
				return
			}

			err := img.CatURL(s, "0", "0")
			if err != nil {
				fmt.Println(err.Error())
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(imageCatCmd)
}
