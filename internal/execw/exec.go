package execw

import (
	"fmt"
	"os/exec"
)

// Execute execute
func Execute(args []string) (string, error) {
	output, err := exec.Command(args[0], args[1:]...).Output()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(output), err
}
