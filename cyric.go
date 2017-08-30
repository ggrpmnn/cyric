package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

var (
	debug     = true
	targetURL = "http://rbcpc.org/404/"
	//targetURL = "https://www.lehigh.edu/~inwww/form-test.html"
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
	page.Parse(doc)

	if debug {
		for _, f := range page.Fields {
			fmt.Printf("%+v\n", f)
		}
	}

	// res, err = http.PostForm(targetURL, params)
	// if err != nil {
	// 	log.Fatalln("error posting data back to site")
	// }

	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatalln("error reading second response")
	// }
	// fmt.Println(string(data))
}
