package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := flag.String("port", "8080", "port number")
	host := flag.String("host", "http://localhost", "host address (including protocol)")
	flag.Parse()

	resp, err := http.Get(fmt.Sprintf("%v:%v/health", host, *port)) // Note pointer dereference using *

	// If there is an error or non-200 status, exit with 1 signaling unsuccessful check.
	if err != nil || resp.StatusCode != 200 {
		os.Exit(1)
	}
	os.Exit(0)
}
