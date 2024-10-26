package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
)

type cliFlags struct {
	port      string
	directory string
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func main() {
	var opts cliFlags
	flag.StringVar(&opts.port, "p", "8100", "port to serve on")
	flag.StringVar(&opts.directory, "d", ".", "static files directory to host")
	flag.Parse()

	http.Handle("/", corsMiddleware(
		http.FileServer(
			http.Dir(opts.directory),
		),
	))

	absDir, err := filepath.Abs(opts.directory)
	if err != nil {
		log.Fatalf("rrror getting absolute directory path: %v", err)
	}

	log.Printf("serving file from %q on port: %s\n", absDir, opts.port)
	if err := http.ListenAndServe(":"+opts.port, nil); err != nil {
		log.Fatalf("error starting the file server: %v", err)
	}
}
