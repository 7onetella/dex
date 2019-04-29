package httpw

import (
	"io/ioutil"
	"net/http"
	"time"
)

// Get does http get
func Get(url string) ([]byte, error) {
	client := http.Client{
		Timeout: time.Second * 3,
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)

	return data, err
}
