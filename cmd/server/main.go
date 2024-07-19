package main

import (
	"fmt"
	"inicio/api"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	absPath, err := filepath.Abs("web/vue/dist")
	if err != nil {
		fmt.Println("Error determining absolute path:", err)
		return
	}

	// Serve static files from the Vue build directory
	fs := http.FileServer(http.Dir(absPath))

	// Setup API endpoints
	api.SetupApiEndpoints()

	// Handle requests and log only non-static file routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		path := r.URL.Path
		filePath := filepath.Join(absPath, path)

		isStaticFile := strings.HasPrefix(path, "/css/") || strings.HasPrefix(path, "/js/") || strings.HasPrefix(path, "/img/") || strings.HasPrefix(path, "/favicon.ico")
		if !isStaticFile {
			defer func() {
				elapsed := time.Since(start)
				elapsedMillis := float64(elapsed.Nanoseconds()) / 1e6
				fmt.Printf("%s Requested path: '%s' Executed in %.3fms\n", time.Now().UTC().Format("2006-01-02 15:04:05"), path, elapsedMillis)
			}()
		}

		if _, err := os.Stat(filePath); os.IsNotExist(err) || filepath.Ext(filePath) == "" {
			http.ServeFile(w, r, filepath.Join(absPath, "index.html"))
		} else {
			fs.ServeHTTP(w, r)
		}
	})

	// Start the HTTP server
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
