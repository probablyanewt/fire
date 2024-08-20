package server

import (
	"errors"
	"fmt"
	"github.com/probablyanewt/fire/internal/page"
	"net/http"
	"os"
)

func Start(pageTree *page.Page) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request received: %q %q\n", r.Method, r.URL.Path)
		result, err := pageTree.GetDeepChildByUri(r.URL.Path)
		if err != nil || result.GetTemplate() == nil {
			fmt.Printf("404: %q %q\n", r.Method, r.URL.Path)
			http.NotFound(w, r)
			return
		}

		fmt.Printf("200: %q %q\n", r.Method, r.URL.Path)
		result.GetTemplate().Execute(w, struct{}{})
	})

	println("Starting server...")
	err := http.ListenAndServe(":42069", nil)
	println("here")
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
