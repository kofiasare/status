package main

import (
	"errors"
	"net/http"
	"os"

	"github.com/fatih/color"
	cecil "github.com/kofiasare/cecil/logger"
)

const (
	version = " version 0.1"
)

var (
	l   = cecil.NewLogger(os.Stdout, "", 0)
	err = errors.New("unknown status")
)

type HTTPStatusCodes struct {
}

// HTTPStatusCode New
type HTTPStatusCode struct {
	statusCode  int
	message     string
	description string
	color       color.Attribute
	*AssociateError
}

// LookUp returns an HTTPStatusCode with
func (h *HTTPStatusCodes) LookUp(c int, v bool) (*HTTPStatusCode, *AssociateError) {
	switch http.StatusText(c) {
	case "":
		return &HTTPStatusCode{color: color.FgRed}, &AssociateError{err}
	default:
		if v {
			return &HTTPStatusCode{
				statusCode:  c,
				message:     statusMessage(c),
				description: statusDescription(c),
				color:       statusColor(c),
			}, nil
		}

		return &HTTPStatusCode{
			statusCode: c,
			message:    statusMessage(c),
			color:      statusColor(c),
		}, nil

	}
}

// Describe describes or explains an HTTPStatusCode
func (h *HTTPStatusCodes) Describe(sc *HTTPStatusCode) {
	prettyPrint(sc)
}

// preety print to stdout
// what you have
func prettyPrint(hs *HTTPStatusCode) {
	color.Set(color.Bold, hs.color)
	l.Printf("  %d - %s", hs.statusCode, hs.message)
	if hs.description != "" {
		defer l.Println(hs.description)
	}
	defer color.Unset()
}

// ListAll list all HTTPStatusCodes
func (h *HTTPStatusCodes) ListAll() {
	for code := range StatusCodes {
		c := &HTTPStatusCode{
			statusCode:  code,
			message:     statusMessage(code),
			description: statusDescription(code),
		}
		l.Printf("%d - %s", c.statusCode, c.message)
	}

}

// AssociateError HTTPStatusCode Error
type AssociateError struct {
	err error
}

// DescribeErrFrom describes or explains an error
// always "unkown status"
func (ae *AssociateError) DescribeErrFrom(hsc *HTTPStatusCode) {
	ae.printError(ae, hsc)
}

func (ae *AssociateError) printError(err *AssociateError, hsc *HTTPStatusCode) {
	defer color.Unset()
	color.Set(color.Bold, hsc.color)
	l.Println(err.errorMessage())
}

func (ae *AssociateError) errorMessage() string {
	return ae.err.Error()
}
