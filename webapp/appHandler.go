package webapp

import (
	"embed"
	"fmt"
	"net/http"
	"path/filepath"
)

var (
	//go:embed all:build/*

	res embed.FS
)

func AppHandler(w http.ResponseWriter, r *http.Request) {

	var resource string

	path := r.URL.Path

	if path == "/" {
		resource = "build/index.html"
	} else {
		resource = fmt.Sprintf("build%s", path)
	}

	b, err := res.ReadFile(resource)
	if err != nil {
		// Just in case we're missing the /index.html, add it and try again...
		resource += "/index.html"
		b, err = res.ReadFile(resource)
		if err != nil {
			resource = "build/index.html"
			b, err = res.ReadFile(resource)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				_, _ = w.Write([]byte("Not found"))
				return
			}
		}
	}

	var mimeType string

	extension := filepath.Ext(resource)
	switch extension {
	case ".css":
		mimeType = "text/css"
	case ".js":
		mimeType = "text/javascript"
	case ".png":
		mimeType = "image/png"
	case ".map":
		mimeType = "application/json"
	default:
		mimeType = "text/html"
	}

	w.Header().Set("Content-Type", mimeType)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(b)
}
