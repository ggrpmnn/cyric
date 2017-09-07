package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

var (
	// print debug messages?
	debug bool
	// a test URL with a lot of good form and input elements
	targetURL = "https://www.lehigh.edu/~inwww/form-test.html"
)

func main() {
	flag.BoolVar(&debug, "d", false, "print debugging messages")
	flag.StringVar(&targetURL, "t", "", "the target site to test against")
	flag.Parse()

	debugOut(fmt.Sprintf("making GET request to '%s'", targetURL))
	res, err := http.Get(targetURL)
	if err != nil {
		log.Fatalln("error making request")
	}
	defer res.Body.Close()

	debugOut("parsing HTML response")
	page := Page{URL: targetURL}
	doc, err := html.Parse(res.Body)
	if err != nil {
		log.Fatalln("error parsing HTML response")
	}
	page.ParseHTML(doc)

	if debug {
		log.Printf("HTML fields:\n")
		for _, f := range page.Fields {
			fmt.Printf("%+v\n", f)
		}
	}

	// get values to post to page
	debugOut("POSTing data to form at " + targetURL)
	form := SetFormValues(&page)
	res, err = http.PostForm(targetURL, form)
	if err != nil {
		log.Fatalln("error POSTing data to form")
	}

	// TODO - eval the response
}

// debugOut prints a debug message based on the debug flag state
func debugOut(s string) {
	if debug {
		log.Println(s)
	}
}
