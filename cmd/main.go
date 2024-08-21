package main

import (
	"github.com/probablyanewt/fire/internal/logger"
	"github.com/probablyanewt/fire/internal/page"
	"github.com/probablyanewt/fire/internal/server"
)

func main() {
	logger.LogLogo()
	logger.Info("Parsing templates")
	pageTree := page.ParseCompleteTree()
	server.Start(pageTree)
}
