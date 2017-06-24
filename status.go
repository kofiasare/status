package main

import (
	"net/http"
	"os"

	"github.com/fatih/color"
	cecil "github.com/kofiasare/cecil/logger"
)

var l = cecil.NewLogger(os.Stdout, "", 0)

// HTTPStatusCode New
type HTTPStatusCode struct {
	c int
	m string
	d string
}

func describe(c *HTTPStatusCode) {
	prettyPrint(c)
}

func listAll() {
	for code := range HTTPStatusCodes {
		c := &HTTPStatusCode{code, statusMessage(code), statusDescription(code)}
		l.Printf("%d - %s", c.c, c.m)
	}
}

func alert(smt string) {
	color.Set(color.Bold, color.FgRed)
	defer color.Unset()
	l.Println(smt)
}

func lookUpMesg(c int, v bool) (*HTTPStatusCode, bool) {
	switch http.StatusText(c) {
	case "":
		return nil, false
	default:
		if v {
			return &HTTPStatusCode{c, statusMessage(c), statusDescription(c)}, true
		}
		return &HTTPStatusCode{c: c, m: statusMessage(c)}, true
	}
}

// preety print to stdout
// what you have
func prettyPrint(hs *HTTPStatusCode) {

	// Switch to the right formatting
	// 1xx => Yellow
	// 2xx => Green
	// 3xx => White
	// 4xx and 5xx => Red
	switch {

	// 1xx Informational
	case hs.c < 200:
		color.Set(color.Bold, color.FgYellow)
		l.Printf("  %d - %s", hs.c, hs.m)

		// 2xx Success
	case hs.c < 300:
		color.Set(color.Bold, color.FgGreen)
		l.Printf("  %d - %s", hs.c, hs.m)

		// 3xx Redirection
	case hs.c < 400:
		color.Set(color.Bold, color.FgWhite)
		l.Printf("  %d - %s", hs.c, hs.m)

		// 4xx Client Error
		// 5xx Server Error
	case hs.c < 500 || hs.c <= 600:
		color.Set(color.Bold, color.FgRed)
		l.Printf("  %d - %s", hs.c, hs.m)
	}

	if hs.d != "" {
		defer l.Println(hs.d)
	}
	defer color.Unset()
}
