package dockerw

import (
	"fmt"
	"os"
	"strings"

	"github.com/7onetella/mvk/internal/execw"
	dockerpty "github.com/fgrehm/go-dockerpty"
	docker "github.com/fsouza/go-dockerclient"
)

const dockerRepoBase = ""

// BuildTagPush will build tag and push
func BuildTagPush(repo string, push bool) {
	if strings.HasPrefix(repo, "/") {
		repo = repo[1:]
	}
	dockerURI := dockerRepoBase + repo

	// docker build -t nginx .
	cmds := []string{"docker", "build", "-t", dockerURI, ".", "--no-cache=true", "--pull=true"}
	fmt.Println(cmds)
	err := execw.Exec(cmds)

	if err != nil {
		fmt.Println(err)
		return
	}

	if push {
		cmds = []string{"docker", "push", dockerURI}
		fmt.Println(cmds)
		execw.Exec(cmds)
	}
}

// StartDockerConsole does docker pull and then docker run
func StartDockerConsole(image string, env, cmd []string) {
	endpoint := "unix:///var/run/docker.sock"
	client, _ := docker.NewClient(endpoint)

	dimage, err := client.InspectImage(image)
	if dimage == nil && err != nil {
		fmt.Print("\u2717 " + err.Error() + ". ")
	}

	if dimage == nil {
		fmt.Println("performing docker pull " + image)
		// Create container
		err = client.PullImage(docker.PullImageOptions{Repository: image}, docker.AuthConfiguration{})
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	container, err := client.CreateContainer(docker.CreateContainerOptions{
		Config: &docker.Config{
			Image:        image,
			Env:          env,
			Cmd:          cmd,
			OpenStdin:    true,
			StdinOnce:    true,
			AttachStdin:  true,
			AttachStdout: true,
			AttachStderr: true,
			Tty:          true,
		},
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Cleanup when done
	defer func() {
		client.RemoveContainer(docker.RemoveContainerOptions{
			ID:    container.ID,
			Force: true,
		})
	}()

	// Fire up the console
	if err = dockerpty.Start(client, container, &docker.HostConfig{}); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
