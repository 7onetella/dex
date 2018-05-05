package xpath

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
	xmlpath "gopkg.in/xmlpath.v2"
)

// ScreenScrape hits the given URL and screen scrape  then return dom like object for searching
func ScreenScrape(url string) (*xmlpath.Node, error) {
	failed := "\u2717"

	client := http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, errors.New(failed)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New(failed)
	}

	pageContent, err := ioutil.ReadAll(resp.Body)

	reader := strings.NewReader(string(pageContent))
	root, err := html.Parse(reader)
	if err != nil {
		log.Fatal(err)
	}

	var b bytes.Buffer
	html.Render(&b, root)
	fixedHTML := b.String()

	reader = strings.NewReader(fixedHTML)
	xmlroot, xmlerr := xmlpath.ParseHTML(reader)

	if xmlerr != nil {
		log.Fatal(xmlerr)
	}

	return xmlroot, nil
}

// SearchByXPath will walk down the node and children using xpath expression
func SearchByXPath(context *xmlpath.Node, xpath string) []*xmlpath.Node {
	path := xmlpath.MustCompile(xpath)

	nodes := make([]*xmlpath.Node, 0, 100)

	iter := path.Iter(context)
	for iter.Next() {
		nodes = append(nodes, iter.Node())
	}

	return nodes
}

// XPathGet xpath get by index
func XPathGet(context *xmlpath.Node, xpath string, index int) string {
	nodes := SearchByXPath(context, xpath)
	if index >= len(nodes) {
		fmt.Println("failed to get ", xpath, " index:", index)
		return ""
	}
	return strings.TrimSpace(nodes[index].String())
}
