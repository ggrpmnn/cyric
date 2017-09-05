package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"

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

	// Procedure:
	// 1. Determine which fields must be filled out and which fields
	//    are text-based ("injectable") and add them to the parameters
	// 2. Provide data to form fields (valid data for required fields,
	//    injection/fuzz data for text fields)
	// 3. POST form to site
	// 4. ???
	form := url.Values{}
	for _, f := range page.Fields {
		if f.IsInjectable() {
			form.Add(f.Name, "test") //add injection value
		} else if f.Required {
			form.Add(f.Name, "test")
		}
	}
	if debug {
		log.Println("URL Values:")
		for k, v := range form {
			fmt.Printf("'%s': '%s'\n", k, v)
		}
	}

}

// debugOut prints a debug message based on the debug flag state
func debugOut(s string) {
	if debug {
		log.Println(s)
	}
}
