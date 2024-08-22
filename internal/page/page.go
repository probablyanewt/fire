package page

import (
	"fmt"
	"html/template"
	"io"
	"strings"
)

type Page struct {
	Name     string
	template *template.Template
	children []*Page
	parent   *Page
}

// GetDeepChildByUri uses a breadth first search to find a page which matches the uri path exactly from the current page
// It returns a page if one is found, and any errors encountered
func (p *Page) GetDeepChildByUri(uri string) (*Page, error) {
	if p.Name == uri {
		return p, nil
	}

	// Split uri by / and remove first item as it will be empty
	uriSections := strings.Split(uri, "/")
	_, uriSections = uriSections[0], uriSections[1:]

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

// HasTemplate is used to determine whether a template has been parsed for a given page
// It returns a boolean
func (p *Page) HasTemplate() bool {
	return p.template != nil
}

func (p *Page) RenderTemplate(w io.Writer) {
	p.template.Execute(w, struct{}{})
}

// addChild adds a child to the current page,
// It returns that child.
func (p *Page) addChild(name string) *Page {
	newChild := NewPage(name, p)
	p.children = append(p.children, newChild)
	return newChild
}

// getRootPage will recursively ascend the tree until a node without a parent is found
// It returns the root page
func (p *Page) getRootPage() *Page {
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
func (p *Page) buildFullUri() string {
	node := p
	uriSections := make([]string, 0)
	for {
		if node.parent == nil {
			return strings.Join(uriSections, "/")
		}
		uriSections = append([]string{node.Name}, uriSections...)
		node = node.parent
	}
}

// getChildByName finds the child of a page with a given name if it exists.
// It returns a page if one is found, and any errors encountered.
func (p *Page) getChildByName(name string) (*Page, error) {
	for _, child := range p.children {
		if child.Name == name {
			return child, nil
		}
	}
	return nil, fmt.Errorf("No child %v found on page %v", name, p.Name)
}

// crawlTreeByUri uses a breadth first search from the root node to find a page which matches as much of the uri as possible.
// It returns the deepest page it was able to find, the remainder of the uri , and any error encountered.
func (p *Page) crawlTreeByUri(uri string) (*Page, *string, error) {
	if p.Name == uri {
		return p, nil, nil
	}

	// Split uri by / and remove first item as it will be empty
	uriSections := strings.Split(uri, "/")
	_, uriSections = uriSections[0], uriSections[1:]

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
func (p *Page) addToTreeFromUri(uri string) (*Page, error) {
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

func NewPage(name string, parent *Page) *Page {
	return &Page{
		Name:     name,
		children: make([]*Page, 0),
		parent:   parent,
	}
}
