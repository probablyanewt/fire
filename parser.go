package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func pageParser() page {
	root := page{
		uri: "/",
	}

	templatePaths := make([]string, 0)

	filepath.Walk("./pages", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return err
		}

		templatePaths = append(templatePaths, path)

		return err
	})

	fmt.Printf(strings.Join(templatePaths, "\n"))

	return root
}
