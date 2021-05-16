package main

import (
	"flag"
)

func main() {
	var concurrency int
	var parameters string
	var domains string
	var domain string
	var appendF bool

	flag.IntVar(&concurrency, "c", 30, "The concurrency for speed")
	flag.StringVar(&domains, "l", "", "domains file")
	flag.StringVar(&domain, "d", "", "domain")
	flag.StringVar(&parameters, "p", "", "parameters to generate wordlist")
	flag.BoolVar(&appendF, "a", true, "Append / at the end of the endpoint")
	flag.Parse()

}
