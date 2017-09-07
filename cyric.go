package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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
		log.Fatalln("error making GET request")
	}
	defer res.Body.Close()

	debugOut("parsing HTML response")
	page := Page{URL: targetURL}
	doc, err := html.Parse(res.Body)
	if err != nil {
		log.Fatalln("error parsing GET response HTML")
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
		log.Fatalln("error making POST request to form")
	}
	defer res.Body.Close()

	// TODO - eval the response; for now, print it out
	debugOut("printing POST response")
	if res.Body != nil {
		bytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalln("error parsing POST response body")
		}
		fmt.Println(string(bytes))
	}
}

// debugOut prints a debug message based on the debug flag state
func debugOut(s string) {
	if debug {
		log.Println(s)
	}
}
