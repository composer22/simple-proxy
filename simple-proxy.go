// simple-proxy is a reverse proxy server for passing requests onward to a single endpoint.
package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/composer22/simple-proxy/logger"
	flag "github.com/spf13/pflag"
)

const (
	applicationName = "simple-proxy" // Application name.
	version         = "1.0.0"        // Application and server version.
	usageText       = `
Description: simple-proxy is a server that will pass on requests it receives
to a backendpoint.

Usage: simple-proxy  [options...]

Server options:
    -s, --backend-scheme SCHEME      http or https
    -b, --backend-target TARGET      Host to forward to.

    -d, --debug                      Enable debugging output (default: false)

Common options:
    -h, --help                       Show this message
    -V, --version                    Show version

Example:

    simple-proxy --backend-scheme "https" --backend-target "www.bar.com"

`
)

var (
	log *logger.Logger
)

func init() {
	log = logger.New(logger.UseDefault, false)
}

// PrintVersionAndExit prints the version of the server then exits.
func PrintVersionAndExit() {
	fmt.Printf("%s version %s\n", applicationName, version)
	os.Exit(0)
}

// PrintUsageAndExit is used to print out command line options.
func PrintUsageAndExit() {
	fmt.Printf("%s\n", usageText)
	os.Exit(0)
}

// main entry point for the application or server launch.
func main() {
	var showVersion bool
	var backendScheme string
	var backendTarget string

	flag.StringVarP(&backendScheme, "backend-scheme", "s", "", "Backend scheme ex: https")
	flag.StringVarP(&backendTarget, "backend-target", "b", "", "Backend target ex: foo.bar.com")
	flag.BoolVarP(&showVersion, "version", "v", false, "Show version.")
	flag.Parse()

	// Version flag request?
	if showVersion {
		PrintVersionAndExit()
	}

	// Check additional commands beyond the flags.
	for _, arg := range flag.Args() {
		switch strings.ToLower(arg) {
		case "version":
			PrintVersionAndExit()
		case "help":
			PrintUsageAndExit()
		}
	}

	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			director := func(req *http.Request) {
				req = r
				req.URL.Scheme = backendScheme
				req.URL.Host = backendTarget
			}
			proxy := &httputil.ReverseProxy{Director: director}
			proxy.ServeHTTP(w, r)
		})
	http.ListenAndServe(":80", nil)

	os.Exit(0)
}
