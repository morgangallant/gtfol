package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// Get the first two arguments, the listening port and the destination address (respectively).
	args := os.Args[1:]
	if len(args) != 2 {
		return fmt.Errorf("Usage: %s <port> <destination>", os.Args[0])
	}

	// Make sure the port is valid.
	port, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("Invalid port: %s", err)
	}

	// Make sure we have a valid destination address.
	parsed, err := url.Parse(args[1])
	if err != nil {
		return fmt.Errorf("Invalid destination address: %s", err)
	}

	// Create a new reverse proxy and start up the web server.
	rp := httputil.NewSingleHostReverseProxy(parsed)
	http.Handle("/", rp)
	return http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil)
}
