package cmd

import (
	"fmt"

	"encoding/hex"

	"github.com/7onetella/dex/internal/crypto"
	"github.com/spf13/cobra"
)

var decryptCmdUseDefault bool

var decryptCmd = &cobra.Command{
	Use:     "dec <ciphertext> <key>",
	Short:   "Decrypts specified value with key",
	Long:    `Decrypts specified value with key`,
	Example: "decrypt foo the-key-has-to-be-32-bytes-long!",
	Aliases: []string{"decrypt"},
	Hidden:  true,
	Run: func(cmd *cobra.Command, args []string) {

		if decryptCmdUseDefault {
			goto UseDefault
		}

		if len(args) < 2 {
			cmd.Usage()
			return
		}

	UseDefault:
		// use the default key if not specified
		key := []byte(crypto.DefaultKey)

		if len(args) == 2 {
			key = []byte(args[1])
		}

		ds, err := hex.DecodeString(args[0])
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		ciphertext := []byte(ds)

		plaintext, err := crypto.Decrypt(ciphertext, key)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Printf("%x => %s\n", ciphertext, plaintext)
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)

	flags := decryptCmd.Flags()

	flags.BoolVarP(&decryptCmdUseDefault, "use", "u", false, "use default")

	flags.MarkHidden("use")
}
