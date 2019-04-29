package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/7onetella/dex/internal/ecsw"
	"github.com/spf13/cobra"
)

var portMappingCmdTestMode bool

var myBool bool

/*
	this command assumes the first container's data is the data to be used in generating the output
*/
var portMappingCmd = &cobra.Command{
	Use:   "port-mapping",
	Short: "Gets ECS port mapping",
	Long:  `Gets ECS port mapping`,
	Example: `port-mapping -t
export EC2_HOST=192.168.0.100
export PORT_TCP_8080=32778
export PORT_TCP_9000=32779`,
	Run: func(cmd *cobra.Command, args []string) {

		e := ecsw.NewECSMetaDataRetriever()
		e.Testmode = portMappingCmdTestMode

		dockerID, err := e.GetDockerID()
		handleError("", err)

		data, err := e.GetTask(dockerID)
		handleError("", err)
		task := ecsw.TaskJSON{}
		json.Unmarshal(data, &task)

		ec2Host, err := e.GetECSHost()
		handleError("", err)

		fmt.Printf("export EC2_HOST=%s\n", ec2Host)
		for i, port := range task.Containers[0].Ports {
			if i == 0 {
				fmt.Printf("export PORT_%s_%d=%d", strings.ToUpper(port.Protocol), port.ContainerPort, port.HostPort)
				continue
			} else {
				fmt.Printf("\nexport PORT_%s_%d=%d", strings.ToUpper(port.Protocol), port.ContainerPort, port.HostPort)
			}
		}

	},
}

func init() {
	ecsCmd.AddCommand(portMappingCmd)
	flags := portMappingCmd.Flags()

	flags.BoolVarP(&portMappingCmdTestMode, "test-mode", "t", false, "test mode")
}

func handleError(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %v", msg, err)
		os.Exit(1)
	}
}
