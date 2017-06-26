package main

import "flag"
import "fmt"

func main() {
	// Lets get code to lookup for with some
	// other options that were passed to us
	// verbose(v) or list all (l)
	// list all lists all the status codes available
	// verbose is used along with options c for verbose
	// documentation
	c := flag.Int("c", 0, "the status code you want to lookup")
	a := flag.Bool("a", false, "list all HTTP status codes")
	v := flag.Bool("v", false, "verbose ")
	ver := flag.Bool("version", false, "show version")
	flag.Parse()

	// make reference to HTTPStatusCode
	hsc := &HTTPStatusCodes{}

	// Lets lookup for code by passing our cli args to
	//  this lookUp() function.
	if *c > 0 {
		sc, err := hsc.LookUp(*c, *v)

		// let us know the search outcome
		// when found describe the status
		// code to us otherwise say to us
		// "unknown status"
		if err != nil {
			err.DescribeErrFrom(sc)
			return
		}
		hsc.Describe(sc)
		return
	}

	// With our refrecence, lets lookup for code
	// passing our cli args to the lookUp() mtd
	// of hscExplorer
	if *a {
		hsc.ListAll()
		return
	}

	if *ver {
		fmt.Println(version)
		return
	}

	// If no option is passed print help message
	flag.Usage()
}
