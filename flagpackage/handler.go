package flagpackage

import (
	"flag"
	"fmt"
	"os"
)

var urls URLList

func checkURL(url string) {
	fmt.Println("Checking URL:", url)
}

// FlagsHandler обработчик
func FlagsHandler() {
	var url string
	flag.StringVar(&url, "url", "", "URL to check")

	var retries int
	flag.IntVar(&retries, "retries", 3, "Number of retries for a failed request")

	var threshold float64
	flag.Float64Var(&threshold, "threshold", 0.5, "Threshold value for considering a response to be too slow (in seconds)")

	var silent, verbose bool
	flag.BoolVar(&silent, "silent", false, "Run in silent mode without standard output")
	flag.BoolVar(&verbose, "verbose", false, "Run in verbose mode (overrides silent mode)")

	flag.Var(&urls, "urls", "Comma-separated list of URLs to check")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "This tool performs health checks on specified URLs.\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if url != "" {
		checkURL(url)
	} else if len(urls) > 0 {
		for _, u := range urls {
			checkURL(u)
		}
	} else {
		flag.Usage()
		//os.Exit(1)
	}
}
