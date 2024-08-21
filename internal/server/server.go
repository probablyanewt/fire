package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
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

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	shudownContext, shutdownContextRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownContextRelease()

	err := server.Shutdown(shudownContext)
	if err != nil {
		logger.Error("\nHTTP shutdown error: %v", err)
	}
	logger.Info("Toodle pip")
}

func generateHandler(pageTree *page.Page) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Incoming: %v %v", r.Method, r.URL.Path)

		result, err := pageTree.GetDeepChildByUri(r.URL.Path)
		if err != nil || !result.HasTemplate() {
			logger.Info("404: %v %v", r.Method, r.URL.Path)
			http.NotFound(w, r)
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
			log.Fatalf("HTTP server error: %v", err)
		}
		logger.Info("\nStopped serving new connections.")
	}
}
