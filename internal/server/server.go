package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	"time"

	"github.com/probablyanewt/fire/internal/logger"
	"github.com/probablyanewt/fire/internal/page"
)

func Start(pageTree *page.Page) {
	server := &http.Server{
		Addr: ":42069",
	}

	http.HandleFunc("/", generateHandler(pageTree))

	go generateServeRoutine(server)()
	logger.Info("Running...")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	shudownContext, shutdownContextRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownContextRelease()

	err := server.Shutdown(shudownContext)
	if err != nil {
		logger.Error("HTTP shutdown error: %v", err)
	}
	logger.Info("Toodle pip o/")
}

func generateHandler(pageTree *page.Page) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Debug("Incoming: %v %v", r.Method, r.URL.Path)

		resourceUrlRegEx := regexp.MustCompile("resource:/")
		if resourceUrlRegEx.MatchString(r.URL.Path) {
			logger.Debug("404: %v %v", r.Method, r.URL.Path)
			http.NotFound(w, r)
			return

		}

		result, err := pageTree.GetDeepChildByUri(r.URL.Path)
		if err != nil || !result.HasTemplate() {
			logger.Info("404: %v %v", r.Method, r.URL.Path)
			notFound(w, r)
			return
		}

		logger.Info("200: %v %v", r.Method, r.URL.Path)
		result.RenderTemplate(w)
	}
}

func generateServeRoutine(server *http.Server) func() {
	return func() {
		err := server.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal("HTTP server error: %v", err)
		}
		logger.Info("\nStopped the server.")
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Oops! Doesn't look like you have a template for this path!\nTry adding one in pages", r.URL.Path, ".gohtml")
}
