// Copyright 2017 Google Inc. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to writing, software distributed
// under the License is distributed on a "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
package img

import (
	"encoding/base64"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Copy copies the given image reader and encodes it as an
// iTerm2 image into the writer.
func copy(w io.Writer, r io.Reader, width, height string) error {
	widthVal, _ := strconv.Atoi(width)
	heightVal, _ := strconv.Atoi(height)

	var header *strings.Reader

	if widthVal > 0 && heightVal > 0 {
		header = strings.NewReader("\033]1337;File=inline=1;width=" + width + "px;heigh=" + height + "px:")
	} else {
		header = strings.NewReader("\033]1337;File=inline=1:")
	}

	footer := strings.NewReader("\a\n")

	pr, pw := io.Pipe()
	go func() {
		defer pw.Close()

		wc := base64.NewEncoder(base64.StdEncoding, pw)
		_, err := io.Copy(wc, r)
		if err != nil {
			pw.CloseWithError(errors.Wrap(err, "could not encode image"))
			return
		}

		if err := wc.Close(); err != nil {
			pw.CloseWithError(errors.Wrap(err, "could not close base64 encoder"))
			return
		}
	}()

	_, err := io.Copy(w, io.MultiReader(header, pr, footer))
	return err
}

// NewWriter returns a new imgcat writer.
func NewWriter(w io.Writer, width, height string) io.WriteCloser {
	pr, pw := io.Pipe()

	wc := &writer{pw, make(chan struct{})}
	go func() {
		defer close(wc.done)
		err := copy(w, pr, width, height)
		pr.CloseWithError(err)
	}()
	return wc
}

type writer struct {
	pw   *io.PipeWriter
	done chan struct{}
}

func (w *writer) Write(data []byte) (int, error) {
	return w.pw.Write(data)
}

func (w *writer) Close() error {
	if err := w.pw.Close(); err != nil {
		return err
	}
	<-w.done
	return nil
}

// Cat image cat
func Cat(path string, width, height string) error {
	f, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "could not open image")
	}
	defer f.Close()

	wc := NewWriter(os.Stdout, width, height)
	if _, err = io.Copy(wc, f); err != nil {
		return err
	}
	return wc.Close()
}

// CatS3 image cat from s3 bucket
func CatS3(r io.Reader, width, height string) error {
	wc := NewWriter(os.Stdout, width, height)
	if _, err := io.Copy(wc, r); err != nil {
		return err
	}
	return wc.Close()
}

// CatURL image at from web resource
func CatURL(url string, width, height string) error {
	req, _ := http.NewRequest("GET", url, nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode > 200 {
		return errors.New("issue with downloading")
	}

	wc := NewWriter(os.Stdout, width, height)
	if _, err := io.Copy(wc, res.Body); err != nil {
		return err
	}
	return wc.Close()
}
