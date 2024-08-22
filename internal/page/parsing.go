package page

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/probablyanewt/fire/internal/logger"
)

// ParseCompleteTree parses the pages directory and constructs a complete page tree with parsed templates.
// It returns the root page
func ParseCompleteTree() *Page {
	logger.Info("Parsing templates")
	root := NewPage("/", nil)

	components := make([]string, 0)
	filepath.Walk("./components", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return err
		}

		components = append(components, path)
		return err
	})

	filepath.Walk("./pages", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return err
		}

		tmplErr := root.addTemplateToTreeByFilePath(path, components)
		if tmplErr != nil {
			logger.Fatal("Error parsing template: %v", tmplErr)
		}
		return err
	})

	return root
}

// setTemplateByFilePath sets the template on a page if non exists.
// It returns any errors that occur.
func (p *Page) setTemplateByFilePath(filePath string, components []string) error {
	if p.template != nil {
		return fmt.Errorf("Template for %v exists", filePath)
	}

	templ, err := template.ParseFiles(append([]string{filePath}, components...)...)
	// templ, err := template.ParseFiles(filePath)
	if err != nil {
		println(err.Error())
		logger.Fatal("Failed to parse template at: ", filePath)
	}

	p.template = templ
	return nil
}

// addTemplateToTreeByFilePath attempts to deeply add a node and template to the page tree by parsing the filePath of the template.
// It returns any errors which occured.
func (p *Page) addTemplateToTreeByFilePath(filePath string, components []string) error {
	uri := filePathToUri(filePath)
	logger.Debug("Adding template %v to node %v", filePath, uri)
	newNode, err := p.addToTreeFromUri(uri)
	if err != nil {
		return err
	}
	err = newNode.setTemplateByFilePath(filePath, components)
	return err
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

	if len(splitUri) < 1 {
		return "/"
	}

	return "/" + strings.Join(splitUri, "/")
}
