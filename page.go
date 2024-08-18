package main

import (
	"errors"
	"html/template"
	"log"
)

type page struct {
	uri      string
	template *template.Template
	children []*page
}

func newPage(uri string, template *template.Template) page {
	return page{
		uri:      uri,
		template: template,
	}
}

func newPageFromFilePath(uri, filePath string) page {
	templ, err := template.ParseFiles(filePath, "components/**/*.gohtml")
	if err != nil {
		log.Fatal("Failed to parse template at", filePath)
	}

	return newPage(uri, templ)
}

func (p *page) addChild(page *page) {
	p.children = append(p.children, page)
}

// This can be improved, but it'll do for now
func (p *page) getByUri(path string) (*page, error) {
	if p.uri == path {
		return p, nil
	}

	if len(p.children) > 0 {
		for _, child := range p.children {
			result, _ := child.getByUri(path)
			if result != nil {
				return result, nil
			}
		}
	}

	return nil, errors.New("No page found")
}

func (p *page) setTemplateByPath(path string) {
	templ, err := template.ParseFiles(path, "components/**/*.gohtml")
	if err != nil {
		log.Fatal("Failed to parse template at: ", path)
	}
	p.template = templ
}
