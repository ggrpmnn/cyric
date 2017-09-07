package main

import (
	"fmt"
	"log"
	"net/url"
)

// placeholder values
var (
	fakeText = "Some placeholder text"
	fakeNum  = "12457"
	fakePass = "myP@sw0rD1"
	fakeMail = "test@test.test"
	fakeURL  = "http://www.test.net"
)

// SetFormValues creates the values to be posted to the target page;
// if a field is "injectable", then we set its data and append injection
// data to it; otherwise, if the field is required, we fill it out with
// appropriate placeholder values
func SetFormValues(p *Page) (v url.Values) {
	v = url.Values{}
	for _, f := range p.Fields {
		if f.IsInjectable() {
			switch f.InputType {
			case "": //textarea
				fallthrough
			case "text":
				v.Add(f.Name, fakeText+"';")
			case "password":
				v.Add(f.Name, fakePass+"';")
			case "email":
				v.Add(f.Name, fakeMail+"';")
			case "url":
				v.Add(f.Name, fakeURL+"';")
			}
		} else if f.Required {
			switch f.InputType {
			case "": //textarea
				fallthrough
			case "text":
				v.Add(f.Name, fakeText)
			case "number":
				v.Add(f.Name, fakeNum)
			case "password":
				v.Add(f.Name, fakePass)
			case "email":
				v.Add(f.Name, fakeMail)
			case "url":
				v.Add(f.Name, fakeURL)
			default:
				// TODO - handle non-text inputs (radio buttons, check boxes, etc.)
				log.Println("error: required value is non-text field")
			}
		}
	}

	if debug {
		log.Println("URL Values:")
		for k, d := range v {
			fmt.Printf("'%s': '%s'\n", k, d)
		}
	}

	return
}
