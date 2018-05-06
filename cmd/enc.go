package cmd

import (
	"fmt"

	"github.com/7onetella/mvk/internal/crypto"
	"github.com/spf13/cobra"
)

var encryptCmdUseDefault bool

var encryptCmd = &cobra.Command{
	Use:     "enc <plaintext> <key>",
	Short:   "Encrypts specified value with key",
	Long:    `Encrypts specified value with key`,
	Example: "enc foo the-key-has-to-be-32-bytes-long!",
	Aliases: []string{"encrypt"},
	Hidden:  true,
	Run: func(cmd *cobra.Command, args []string) {
		if encryptCmdUseDefault {
			goto UseDefault
		}

		if len(args) < 2 {
			cmd.Usage()
			return
		}

	UseDefault:
		plaintext := []byte(args[0])

		// use the default key if not specified
		key := []byte(crypto.DefaultKey)

		if len(args) == 2 {
			key = []byte(args[1])
		}

		if len(key) < 16 {
			fmt.Println("the-key-has-to-be-32-bytes-long!")
		}

		ciphertext, err := crypto.Encrypt(plaintext, key)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("%s => %x\n", plaintext, ciphertext)
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)

	flags := encryptCmd.Flags()

	flags.BoolVarP(&encryptCmdUseDefault, "use", "u", false, "use default")

	flags.MarkHidden("use")

}
