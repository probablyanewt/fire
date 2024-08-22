package cli

import (
	"os"
	"strings"

	"github.com/probablyanewt/fire/internal/logger"
	"github.com/probablyanewt/fire/internal/page"
	"github.com/probablyanewt/fire/internal/server"
)

func Run() {
	logger.LogLogo()
	logger.Debug("Args: %+v", strings.Join(os.Args[1:], " "))
	pageTree := page.ParseCompleteTree()
	server.Start(pageTree)
}
