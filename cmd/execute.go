package cmd

import "flag"

func Execute() error {
	var showHelp bool
	flagBoolVarP(&showHelp, "help", "h", false, "Display the help text and exit")

	var showVersion bool
	flagBoolVarP(&showVersion, "version", "v", false, "Display the version of the application and exit")

	var asJSON bool
	flagBoolVarP(&asJSON, "json", "j", false, "Output the result as JSON")

	var opts serveFlags
	flagStringVarP(&opts.directory, "dir", "d", ".", "Set the directory path containing static files to be served. Default is the current working directory")
	flagStringVarP(&opts.port, "port", "p", "8100", "Specify the port number for the server to listen on. Default is 8100")
	flag.Parse()

	if showHelp {
		cmdHelp()
		return nil
	}
	if showVersion {
		return cmdVersion(asJSON)
	}
	return cmdServe(opts)
}
