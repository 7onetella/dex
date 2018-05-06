package cmd

import (
	"fmt"
	"os"
	"runtime"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var autoCompleteSh = "mvk_autocomplete.sh"
var autoCompleteZsh = "_mvk_autocomplete"

// ecsCmd represents the ecs command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates autocomplete scripts",
	Long:  `Generates autocomplete script: `,
	Run: func(cmd *cobra.Command, args []string) {

		rootCmd.GenBashCompletionFile(autoCompleteSh)
		rootCmd.GenZshCompletionFile(autoCompleteZsh)

		// TODO: add support for Linux
		if runtime.GOOS == "darwin" {
			fmt.Printf("Generating bash autocomplete script: /usr/local/etc/bash_completion.d/%s\n", autoCompleteSh)
			err := os.Rename("./"+autoCompleteSh, "/usr/local/etc/bash_completion.d/"+autoCompleteSh)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("Generating zsh  autocomplete script: ~/.oh-my-zsh/completions/%s\n", autoCompleteZsh)
			homedir, _ := homedir.Dir()
			err = os.Rename("./"+autoCompleteZsh, homedir+"/.oh-my-zsh/completions/"+autoCompleteZsh)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
