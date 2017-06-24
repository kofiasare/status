package main

import "flag"

func main() {
	// Lets get code to lookup for with some
	// other options that were passed to us
	// verbose(v) or list all (l)
	// list all lists all the status codes available
	// verbose is used along with options c for verbose
	// documentation
	c := flag.Int("c", 0, "the status code you want to lookup")
	l := flag.Bool("l", false, "list all HTTP status codes")
	v := flag.Bool("v", false, "verbose ")
	flag.Parse()

	// Lets lookup for code by passing our cli args to
	//  this lookUp() function.
	if *c > 0 {
		hsc, found := lookUpMesg(*c, *v)

		// let us know the search outcome
		// when found describe the status
		// code to us otherwise say to us
		// "unknown status"
		if found {
			describe(hsc)
			return
		}
		alert("  unknown status")
		return
	}

	// With our refrecence, lets lookup for code
	// passing our cli args to the lookUp() mtd
	// of hscExplorer
	if *l {
		listAll()
		return
	}

	// If no option is passed print help message
	flag.Usage()
}
