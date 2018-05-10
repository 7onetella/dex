package cmd

import (
	"fmt"

	"github.com/7onetella/dex/internal/dockerw"
	"github.com/7onetella/dex/internal/net"
	"github.com/spf13/cobra"
)

const succeed = "\u2713"
const failed = "\u2717"

var mysqlCmd = &cobra.Command{
	Use:   "mysql <dbhost> <database> <user> [<password>]",
	Short: "Runs mysql client",
	Long: `Runs mysql client
	
	dbhost     database hostname
	database   database name
	user       user name	
	password   optional paramater. password will be prompted	

	`,
	Example: "mysql db.mars.com tododb admin pass1234",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 3 {
			cmd.Usage()
			return
		}

		dbhost := args[0]
		database := args[1]
		username := args[2]
		password := args[3]

		ok := net.IsTCPConnValid(dbhost, "3306")
		if !ok {
			fmt.Println()
			fmt.Println(failed + " check network connection. i.e. vpn")
			return
		}

		env := []string{}
		dockerCmd := []string{"mysql", "-h", dbhost, "-u", username, "-p" + password, "-D" + database}

		dockerw.StartDockerConsole("mysql:5.6.34", env, dockerCmd)
	},
}

func init() {
	rootCmd.AddCommand(mysqlCmd)
}
