package docs

import (
	"embed"
	"github.com/betterde/focusly/internal/journal"
	"io/fs"
	"net/http"
)

//go:embed api/*
var FS embed.FS

func Serve() http.FileSystem {
	dist, err := fs.Sub(FS, "api")
	if err != nil {
		journal.Logger.Panicw("Error mounting front-end static resources!", err)
	}

	return http.FS(dist)
}
