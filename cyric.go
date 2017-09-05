package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

var (
	debug     = true
	targetURL = "https://www.lehigh.edu/~inwww/form-test.html"
)

func main() {
	// make a request to the target page
	res, err := http.Get(targetURL)
	if err != nil {
		log.Fatalln("error making request")
	}
	defer res.Body.Close()

	page := Page{URL: targetURL}
	doc, err := html.Parse(res.Body)
	if err != nil {
		log.Fatalln("error parsing HTML response")
	}
	page.ParseHTML(doc)

	if debug {
		for _, f := range page.Fields {
			fmt.Printf("%+v\n", f)
		}
	}

	// page struct now contains the list of input fields
}
