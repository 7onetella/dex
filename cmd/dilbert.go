package cmd

import (
	"fmt"
	"os"

	"github.com/7onetella/dex/internal/img"
	"github.com/7onetella/dex/internal/xpath"
	"github.com/spf13/cobra"
)

var dilbertCmd = &cobra.Command{
	Use:   "dilbert",
	Short: "Shows today's dilbert comic",
	Long:  `Shows today's dilbert comic`,
	Run: func(cmd *cobra.Command, args []string) {
		i := 0
		if len(args) == 1 {
			day := args[0]
			switch day {
			case "today":
				i = 0
			case "yesterday":
				i = 1
			}
		}
		dilbert(i)
	},
}

func init() {
	rootCmd.AddCommand(dilbertCmd)
}

func dilbert(index int) {
	root, err := xpath.ScreenScrape("http://dilbert.com/")
	if err != nil {
		fmt.Println(err.Error())
	}

	images := xpath.SearchByXPath(root, "//div[@class='img-comic-container']/a/img/@src")

	term, ok := os.LookupEnv("TERM_PROGRAM")

	fmt.Println()
	for i, image := range images {
		if i != index {
			continue
		}
		s := image.String()

		if ok && term == "iTerm.app" {
			err := img.CatURL(s, "0", "0")
			if err != nil {
				fmt.Println(err.Error())
			}
			// always do the first one which is today's
			break
		}

		fmt.Println(s)
		// always do the first one which is today's
		break
	}

}
