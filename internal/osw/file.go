package osw

import (
	"io/ioutil"
	"os"
)

// ReadFile reads from file
func ReadFile(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}
