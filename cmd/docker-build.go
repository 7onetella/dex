package cmd

import (
	"fmt"

	"github.com/7onetella/mvk/internal/execw"
	"github.com/spf13/cobra"
)

var dockerNoPush bool
var dockerBuildFile string

var dockerBuildTagPushCmd = &cobra.Command{
	Use:     "build <repo/namespace:tag>",
	Short:   "Docker build tag & push",
	Long:    `Docker build tag & push`,
	Example: "build 7onetella/alpinegovim:1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
			return
		}

		dockerURI := args[0]

		// docker build -t alpinegovim .
		cmds := []string{"docker", "build"}

		if dockerBuildFile != "" {
			cmds = append(cmds, "-f")
			cmds = append(cmds, dockerBuildFile)
		}

		cmds = append(cmds, "-t")
		cmds = append(cmds, dockerURI)
		cmds = append(cmds, "--no-cache=true")
		cmds = append(cmds, "--pull=true")
		cmds = append(cmds, ".")

		fmt.Println(cmds)
		err := execw.Exec(cmds)

		if err != nil {
			fmt.Println(err)
			return
		}

		// docker push 7onetella/alpinegovim:1.0.0
		cmds = []string{"docker", "push", dockerURI}
		fmt.Println(cmds)
		execw.Exec(cmds)

	},
}

func init() {
	dockerCmd.AddCommand(dockerBuildTagPushCmd)

	dockerBuildTagPushCmd.Flags().BoolVarP(&dockerNoPush, "nopush", "n", false, "docker build & tag, no push")
	dockerBuildTagPushCmd.Flags().StringVarP(&dockerBuildFile, "file", "f", "", "Dockerfile location")

}
