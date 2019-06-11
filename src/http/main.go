package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		response, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v %v\n", err, response.Status)
			os.Exit(1)
		}

		body, err := ioutil.ReadAll(response.Body)

		response.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s\t%s", body, response.Status)
	}
}

// go run main.go http://baidu.com
