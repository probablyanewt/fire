package main

import (
	// "errors"
	"fmt"
	"html/template"
	"log"
	// "path/filepath"
	"strings"
)

type page struct {
	name     string
	template *template.Template
	children []*page
	parent   *page
}

// addChild adds a child to the current page,
// It returns that child.
func (p *page) addChild(name string) *page {
	newChild := newPage(name, p)
	p.children = append(p.children, newChild)
	return newChild
}

// getRootPage will recursively ascend the tree until a node without a parent is found
// It returns the root page
func (p *page) getRootPage() *page {
	lastNode := p
	for {
		if lastNode.parent == nil {
			return lastNode
		}
		lastNode = lastNode.parent
	}
}

// buildFullUri will recursively ascend the tree to build a full uri path for the page.
// It returns a uri as a string.
func (p *page) buildFullUri() string {
	node := p
	uriSections := make([]string, 0)
	for {
		if node.parent == nil {
			return strings.Join(uriSections, "/")
		}
		uriSections = append([]string{node.name}, uriSections...)
		node = node.parent
	}
}

// getChildByName finds the child of a page with a given name if it exists.
// It returns a page if one is found, and any errors encountered.
func (p *page) getChildByName(name string) (*page, error) {
	for _, child := range p.children {
		if child.name == name {
			return child, nil
		}
	}
	return nil, fmt.Errorf("No child %v found on page %v", name, p.name)
}

// getByUri uses a breadth first search to find a page which matches the uri path exactly from the current page
// It returns a page if one is found, and any errors encountered
func (p *page) getByUri(uri string) (*page, error) {
	if p.name == uri {
		return p, nil
	}

	uriSections := strings.Split(uri, "/")
	lastNode := p
	for _, uriSection := range uriSections {
		result, err := lastNode.getChildByName(uriSection)
		if err != nil {
			return nil, err
		}

		lastNode = result
	}

	return lastNode, nil
}

// crawlTreeByUri uses a breadth first search from the root node to find a page which matches as much of the uri as possible.
// It returns the deepest page it was able to find, the remainder of the uri , and any error encountered.
func (p *page) crawlTreeByUri(uri string) (*page, *string, error) {
	if p.name == uri {
		return p, nil, nil
	}

	uriSections := strings.Split(uri, "/")
	lastNode := p
	for i, uriSection := range uriSections {
		result, err := lastNode.getChildByName(uriSection)
		if err != nil {
			remaining := strings.Join(uriSections[i:], "/")
			return lastNode, &remaining, nil
		}

		lastNode = result
	}

	return lastNode, nil, nil
}

// addToTreeFromUri attempts to deeply add a node to the page tree by traversing the uri
// It returns the last node added and any errors encountered
func (p *page) addToTreeFromUri(uri string) (*page, error) {
	node, remainder, err := p.crawlTreeByUri(uri)
	if err != nil {
		return nil, err
	}

	if remainder == nil {
		return node, nil
	}

	uriSections := strings.Split(*remainder, "/")
	lastNode := node
	for _, path := range uriSections {
		lastNode = lastNode.addChild(path)
	}

	return lastNode, nil
}

// addToTreeByFilePath attempts to deeply add a node to the page tree by parsing the filePath of the template.
// It returns any errors which occured.
// func (p *page) addToTreeByFilePath(filePath string) error {
// 	uri := filePathToUri(filePath)
// 	newNode, err := p.addToTreeByUri(uri)
// 	if err != nil {
// 		return err
// 	}
// 	newNode.setTemplateByFilePath(filePath)
// 	return nil
// }

// setTemplateByFilePath sets the template on a page if non exists.
// It returns any errors that occur.
func (p *page) setTemplateByFilePath(filePath string) error {
	if p.template != nil {
		return fmt.Errorf("Template for %v exists", filePath)
	}

	templ, err := template.ParseFiles(filePath, "components/**/*.gohtml")
	if err != nil {
		log.Fatal("Failed to parse template at: ", filePath)
	}

	p.template = templ
	return nil
}

func newPage(name string, parent *page) *page {
	return &page{
		name:     name,
		children: make([]*page, 0),
		parent:   parent,
	}
}

func newPageTree() *page {
	return newPage("/", nil)
}

// filePathToUri converts a filepath for a template, into an equivalent web uri.
// Returns a uri as a string
func filePathToUri(filePath string) string {
	splitPath := strings.Split(filePath, "/")
	// Remove pages/ from the beginning
	_, splitUri := splitPath[0], splitPath[1:]
	fileName := &splitUri[len(splitUri)-1]

	*fileName = strings.Replace(*fileName, ".gohtml", "", -1)

	if *fileName == "index" {
		splitUri = splitUri[:len(splitUri)-1]
	}

	return strings.Join(splitUri, "/")
}
