package cmd

import (
	"fmt"

	"github.com/7onetella/mvk/internal/dockerw"
	"github.com/7onetella/mvk/internal/net"
	"github.com/spf13/cobra"
)

var psqlCmd = &cobra.Command{
	Use:   "psql <dbhost> <database> <user> [<password>]",
	Short: "Runs psql client",
	Long: `Runs psql client
	
	dbhost     database hostname
	database   database name
	user       user name	
	password   optional paramater. password will be prompted	

	`,
	Example: "psql db.venus.com tododb admin pass1234",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 3 {
			cmd.Usage()
			return
		}

		dbhost := args[0]
		database := args[1]
		username := args[2]
		password := args[3]

		ok := net.IsTCPConnValid(dbhost, "5432")
		if !ok {
			fmt.Println()
			fmt.Println(failed + " check network connection. i.e. vpn")
			return
		}

		env := []string{"PGPASSWORD=" + password}
		dockerCmd := []string{"psql", "-h", dbhost, "-U", username, database}

		dockerw.StartDockerConsole("postgres:9.6-alpine", env, dockerCmd)
	},
}

func init() {
	rootCmd.AddCommand(psqlCmd)
}
