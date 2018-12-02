package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/alecthomas/template"
	"github.com/spf13/cobra"
)

var scriptDockerBuildCmdDebug bool
var scriptDockerBuildCmd = &cobra.Command{
	Use:     "dockerbuild",
	Short:   "Generates docker build script",
	Long:    `Generates docker build script`,
	Example: "dockerbuild",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println()

		filename := fmt.Sprintf("build__%s", "dex_temp_file")

		fmt.Println("Generating", filename)

		f, _ := os.Create(filename)
		w := bufio.NewWriter(f)
		data := map[string]string{
			"pre_docker_build_step": `# Write your pre docker build step here
# e.g. maven build using docker container
# e.g. sbt   build using docker container`,
		}
		templateExecute(data, w)
		w.Flush()

	},
}

func init() {
	scriptCmd.AddCommand(scriptDockerBuildCmd)

	flags := scriptDockerBuildCmd.Flags()

	flags.BoolVarP(&scriptDockerBuildCmdDebug, "debug", "d", false, "debug script dockerbuild")
}

// RecordFailedLinksAsTXT writes filed links as txt
func templateExecute(ctx interface{}, w io.Writer) {
	tmpl, err := template.New("docker build template").Parse(scriptDockerBuildCmdTemplate)
	if err == nil {
		tmpl.Execute(w, ctx)
	}
}

const scriptDockerBuildCmdTemplate = `#!/bin/sh

tag=$1                  # 1.0.0
namespace_repo=$2       # acme/app  <= no slash in front
aws_account_number=$3   #
docker_file_path=$4     #

repo_uri=${aws_account_number}.dkr.ecr.us-east-1.amazonaws.com/${namespace_repo}

if [ "$1" = "" ]; then
    echo "Docker tag is required"
    echo "Usage: build 1.0.0"
    exit 1
fi

if [ "$4" = "" ]; then
   docker_file_path=.
fi

{{.pre_docker_build_step}}

# ECR login
# This assumes that aws credential is provided
aws ecr get-login --no-include-email --region us-east-1 | awk '{ print $6 }' | docker login --username AWS --password-stdin https://${aws_account_number}.dkr.ecr.us-east-1.amazonaws.com

# build local docker image and tag it
docker build -t ${namespace_repo}:${tag} ${docker_file_path}

# tag the local docker image with remote repo and tag
docker tag ${namespace_repo}:${tag} ${repo_uri}:${tag}

# push to remote
docker push ${repo_uri}:${tag}

# clean up temporary docker images
docker system prune --force > /dev/null

# remove the local docker image tag in local storage
docker rmi ${repo_uri}:${tag} || true

# remove the remote docker image tag in local storage
docker rmi ${namespace_repo}:${tag} || true`
