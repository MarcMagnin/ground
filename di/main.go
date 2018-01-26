// How can you make this code more flexible for testing?

package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	url string
)

func init() {
	flag.StringVar(&url, "url", "http://google.com", "Which URL do we want to parse?")
	flag.Parse()
}

// This simple program will do an http request to http://google.com and print the result on stdout.
// Unless the url flag is passed with a value
func main() {
	err := get(url)

	if err != nil {
		panic(err)
	}
}

// We want to test that get() with a valid response always works
// How can we do that?
func get(link string) error {
	client := http.Client{}
	response, err := client.Get(link)

	if err != nil {
		return err
	}

	if response == nil {
		return errors.New("response is nil")
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return err
}
