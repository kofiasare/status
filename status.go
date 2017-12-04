package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
)

const (
	version = " version 0.1"
)

var (
	l   = log.New(os.Stdout, "", 0)
	err = errors.New("HTTP status code not found")
)

// HTTPStatusCodes all http status codes
type HTTPStatusCodes struct{}

// HTTPStatusCode  Represents an HTTPStatusCode
type HTTPStatusCode struct {
	statusCode  int
	message     string
	description string
	color       color.Attribute
	err         error
}

// LookUp returns an HTTPStatusCode with
// any lookup error, nil if the lookup
// found something, "unknown" if nothing
// found
func (HTTPStatusCodes) LookUp(code int, verbose bool) (*HTTPStatusCode, error) {
	switch http.StatusText(code) {
	case "":
		return &HTTPStatusCode{color: color.FgRed}, err
	default:
		if verbose {
			return &HTTPStatusCode{
				statusCode:  code,
				message:     statusMessage(code),
				description: statusDescription(code),
				color:       statusColor(code),
			}, nil
		}

		return &HTTPStatusCode{
			statusCode: code,
			message:    statusMessage(code),
			color:      statusColor(code),
		}, nil
	}
}

// Describe gives description for some HTTPStatusCode
func (HTTPStatusCodes) Describe(code *HTTPStatusCode) {

	color.Set(color.Bold, code.color)
	defer color.Unset()

	l.Printf("  %d - %s", code.statusCode, code.message)
	if code.description != "" {
		l.Println(code.description)
	}
}

// ListAll all HTTPStatusCodes
func (HTTPStatusCodes) ListAll() {
	for code := range StatusCodes {
		c := HTTPStatusCode{statusCode: code, message: statusMessage(code)}
		l.Printf("%d - %s", c.statusCode, c.message)
	}

}

// DescribeErr lookup err
func (HTTPStatusCodes) DescribeErr(hsc *HTTPStatusCode) {
	color.Set(color.Bold, hsc.color)
	defer color.Unset()
	l.Println(hsc.err)
}
