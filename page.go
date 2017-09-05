package main

import (
	"strconv"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Page structs represent a page that has been crawled
type Page struct {
	Title  string
	URL    string
	Fields []Input
}

// Input structs represent an input field on a web page
type Input struct {
	Element   string // either `input` or `textarea`
	Name      string
	ID        string
	InputType string // always empty for `textarea` elements
	Classes   []string
	MaxLength int // max size of input the field accepts
	Required  bool
	ReadOnly  bool
}

//ParseHTML parses a page and adds input fields to the page struct
func (p *Page) ParseHTML(node *html.Node) {
	if node.Type == html.ElementNode {
		switch node.DataAtom {
		case atom.Title:
			p.Title = node.FirstChild.Data
		case atom.Input:
			input := Input{Element: atom.Input.String()}
			input.ParseField(node)
			p.Fields = append(p.Fields, input)
		case atom.Textarea:
			input := Input{Element: atom.Textarea.String()}
			input.ParseField(node)
			p.Fields = append(p.Fields, input)
		}
	}
	// recurse down the document tree
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		p.ParseHTML(child)
	}
}

// ParseField parses the content of input fields and populates the input struct
func (i *Input) ParseField(n *html.Node) {
	// loop through attributes and capture relevant data
	for _, a := range n.Attr {
		switch a.Key {
		case "name":
			i.Name = strings.ToLower(a.Val)
		case "id":
			i.ID = strings.ToLower(a.Val)
		case "type":
			i.InputType = strings.ToLower(a.Val)
		case "class":
			i.Classes = strings.Split(a.Val, " ")
		case "maxlength":
			i.MaxLength, _ = strconv.Atoi(a.Val)
		case "required":
			i.Required = true
		case "readonly":
			i.ReadOnly = true
		}
	}
}

// IsInjectable returns true for input fields where text is allowed; false otherwise
func (i *Input) IsInjectable() bool {
	switch i.InputType {
	case "", "text", "password", "email", "url": // for textareas
		return true
	default:
		return false
	}
}
