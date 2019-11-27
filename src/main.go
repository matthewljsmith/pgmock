package main

import (
	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"

	"github.com/matthewljsmith/pgmock/pgserver"
)

func main() {

	// define some flags that get passed in
	var verbose = flag.Bool("verbose", false, "verbose - pass this to set the debug level to Debug instead of Info")
	flag.Parse()

	// if verbose set log verbosity
	if *verbose {
		log.SetLevel(log.DebugLevel)
	}

	// new up the mocking server
	mock := pgmock.NewServer()

	// kick of the mocking instance
	log.Infof("starting pgmock %s", "127.0.0.1:7432")
	mock.ListenAndServe("127.0.0.1:7432")
}
