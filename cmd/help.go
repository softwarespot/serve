package cmd

import "fmt"

func cmdHelp() {
	helpText := `Usage: ./serve [OPTIONS]

This application is a simple static file server that serves files from either the current working directory or specified directory over HTTP with CORS support, allowing easy access for development and testing

Options:
  -d, --dir       Set the directory path containing static files to be served. Default is the current working directory.
  -p, --port      Port number for the server to listen on. Default is 8100.
  -h, --help      Show this help text and exit.
  -v, --version   Display the version of the application and exit.
  -j, --json      Output the version as JSON.

Examples:
  ./serve`
	fmt.Println(helpText)
}
