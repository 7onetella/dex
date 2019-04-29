package ecsw

import (
	"strings"

	"github.com/7onetella/dex/internal/httpw"
	"github.com/7onetella/dex/internal/osw"
)

// ECSMetaDataRetriever retrieves ECS meta data
type ECSMetaDataRetriever struct {
	Testmode bool
}

// NewECSMetaDataRetriever creates a new instance
func NewECSMetaDataRetriever() ECSMetaDataRetriever {
	e := ECSMetaDataRetriever{}
	return e
}

// GetDockerID retrieves docker id
func (e ECSMetaDataRetriever) GetDockerID() (string, error) {
	if e.Testmode {
		return "abc239847234958734985734957348957349875987347598347589734957234d", nil
	}

	data, err := osw.ReadFile("/proc/1/cpuset")
	if err != nil {
		return "", err
	}

	dockerid := strings.Replace(string(data), "/docker/", "", 1)

	return dockerid, err
}

// GetTask hits http://172.17.0.1:51678/v1/tasks?dockerid=
func (e ECSMetaDataRetriever) GetTask(dockerid string) ([]byte, error) {
	if e.Testmode {
		return osw.ReadFile("./data/taskmetadata")
	}

	return httpw.Get("http://172.17.0.1:51678/v1/tasks?dockerid=" + dockerid)
}

// GetECSHost hits http://172.17.0.1:51678/v1/tasks?dockerid=
func (e ECSMetaDataRetriever) GetECSHost() (string, error) {
	if e.Testmode {
		return "192.168.0.100", nil
	}

	data, err := httpw.Get("http://169.254.169.254/latest/meta-data/local-ipv4")
	if err != nil {
		return "", err
	}
	return string(data), err
}

type TaskJSON struct {
	Arn           string `json:"Arn"`
	DesiredStatus string `json:"DesiredStatus"`
	KnownStatus   string `json:"KnownStatus"`
	Family        string `json:"Family"`
	Version       string `json:"Version"`
	Containers    []struct {
		DockerID   string `json:"DockerId"`
		DockerName string `json:"DockerName"`
		Name       string `json:"Name"`
		Ports      []struct {
			ContainerPort int    `json:"ContainerPort"`
			Protocol      string `json:"Protocol"`
			HostPort      int    `json:"HostPort"`
		} `json:"Ports"`
	} `json:"Containers"`
}
