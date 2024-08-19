package main

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func pageParser() *page {
	root := newPageTree()

	filepath.Walk("./pages", filePathWalker)

	return root
}

func filePathWalker(path string, info os.FileInfo, err error) error {

	if info.IsDir() {
		return err
	}

	splitPath := strings.Split(path, "/")

	_, trimmedPath := splitPath[0], splitPath[1:]

	gohtmlRegEx := regexp.MustCompile(".*.gohtml")
	if gohtmlRegEx.MatchString(trimmedPath[0]) {

	}

	return err
}
