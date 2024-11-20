package cmd

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

type serveFlags struct {
	directory string
	port      string
}

func cmdServe(opts serveFlags) error {
	fs := http.FileServer(http.Dir(opts.directory))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		fs.ServeHTTP(w, r)
	})

	absDir, err := filepath.Abs(opts.directory)
	if err != nil {
		return fmt.Errorf("getting absolute directory path: %w", err)
	}

	log.Printf("serving files from %q. Navigate to http://localhost:%s.\n", absDir, opts.port)
	if err := http.ListenAndServe(":"+opts.port, nil); err != nil {
		return fmt.Errorf("error starting the file server: %w", err)
	}
	return nil
}
