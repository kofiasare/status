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
	err = errors.New("unknown status")
)

// HTTPStatusCodes all http status codes
type HTTPStatusCodes struct{}

// HTTPStatusCode  Represents an HTTPStatusCode
type HTTPStatusCode struct {
	statusCode  int
	message     string
	description string
	color       color.Attribute
	*AssociateError
}

// LookUp returns an HTTPStatusCode with
// any lookup error, nil if the lookup
// found something, "unknown" if nothing
// found
func (h *HTTPStatusCodes) LookUp(c int, v bool) (*HTTPStatusCode, *AssociateError) {
	switch http.StatusText(c) {
	case "":
		return &HTTPStatusCode{color: color.FgRed}, &AssociateError{err}
	default:
		if v {
			return &HTTPStatusCode{c, statusMessage(c), statusDescription(c), statusColor(c), nil}, nil
		}
		return &HTTPStatusCode{statusCode: c, message: statusMessage(c), color: statusColor(c)}, nil
	}
}

// Describe describes or explains an HTTPStatusCode
func (h *HTTPStatusCodes) Describe(sc *HTTPStatusCode) {
	color.Set(color.Bold, sc.color)
	l.Printf("  %d - %s", sc.statusCode, sc.message)
	if sc.description != "" {
		defer l.Println(sc.description)
		defer color.Unset()
	}
}

// ListAll list all HTTPStatusCodes
func (h *HTTPStatusCodes) ListAll() {
	for code := range StatusCodes {
		c := &HTTPStatusCode{statusCode: code, message: statusMessage(code)}
		l.Printf("%d - %s", c.statusCode, c.message)
	}

}

// AssociateError HTTPStatusCode lookup error
type AssociateError struct {
	err error
}

// DescribeErr describes or explains an error
// always "unkown status"
func (ae *AssociateError) DescribeErr(hsc *HTTPStatusCode) {
	color.Set(color.Bold, hsc.color)
	defer color.Unset()
	l.Println(ae.errorMessage())
}

func (ae *AssociateError) errorMessage() string {
	return ae.err.Error()
}
