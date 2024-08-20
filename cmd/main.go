package main

import (
	"github.com/probablyanewt/fire/internal/page"
	"github.com/probablyanewt/fire/internal/server"
)

func main() {
	println("Parsing templates")
	pageTree := page.ParseCompleteTree()

	// result, err := pageTree.getByUri("abc")
	// if err != nil {
	// 	println("oh no")
	// }
	// println(result.template.Execute(os.Stdout, struct{}{}))

	server.Start(pageTree)

}
