package main

import (
	"errors"

	"github.com/fatih/color"
)

var testData = []*HTTPStatusCode{
	{
		100,
		"Continue",
		``,
		color.FgYellow,
		nil,
	},
	{
		900,
		"",
		``,
		color.FgRed,
		&AssociateError{errors.New("unknown status")},
	},
	{
		200,
		"OK",
		``,
		color.FgGreen, nil,
	},
	{
		305,
		"Use Proxy",
		`Defined in a previous version of this specification and is now deprecated, due to security concerns regarding in-band 
		 configuration of a proxy.`,
		color.FgWhite, nil,
	},
	{
		400,
		"Bad Request",
		``,
		color.FgRed,
		nil,
	},
	{
		502,
		"Bad Gateway",
		`The server, while acting as a gateway or proxy, received an invalid response from an inbound server it accessed while 
		 attempting to fulfill the request.`,
		color.FgRed,
		nil,
	},
}
