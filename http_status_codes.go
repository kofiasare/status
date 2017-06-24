package main

// HTTPStatusCodes map of all httpstatus "code"s
// from Go' stdlib
var HTTPStatusCodes = map[int]map[string]string{
	// 1xx Informational
	100: {
		"message": "Continue",
		"description": `
  The initial part of a request has been received and has not yet been rejected by the
  server. The server intends to send a final response after the request has been fully
  received and acted upon. 
  
  When the request contains an Expect header field that includes a 100-continue expectation,
  the 100 response indicates that the server wishes to receive the request payload body. 

  The client ought to continue sending the request and discard the 100 response., If the 
  request did not contain  an Expect header field containing the 100-continue expectation,
  the client can simply discard this interim response.

`,
	},

	101: {
		"message": "Switching Protocols",
		"description": `
  The server understands and is willing to comply with the client's request, via the Upgrade 
  header field, for a change in the application protocol being used on this connection. The 
  server MUST generate an Upgrade header field in the response that indicates which protocol(s)
  will be switched to immediately after the empty line that terminates the 101 response. 
  
  It is assumed that the server will only agree to switch protocols when it is advantageous to
  do so.
  
  For example, switching to a newer version of HTTP might be advantageous over older versions, 
  and switching to a real-time, synchronous protocol might be advantageous when delivering 
  resources that use such features.'
`,
	},

	102: {
		"message": "Processing",
		"description": `
  An interim response used to inform the client that the server has accepted the complete 
  request, but has not yet completed it. This status code SHOULD only be sent when the server
  has a reasonable expectation that the request will take significant time to complete.
  
  As guidance, if a method is taking longer than 20 seconds (a reasonable, but arbitrary value) 
  to process the server SHOULD return a 102 (Processing) response. The server MUST send a 
  final response after the request has been completed. Methods can potentially take a long
  period of time to process, especially methods that support the Depth header.
  
  In such cases, the client may time-out the connection while waiting for a response. To prevent
  this, the server may return a 102 Processing status code to indicate to the client that the server 
  is still processing the method.'
`,
	},

	// 2xx Success

	200: {
		"message": "OK",
		"description": `
  The request has succeeded.

  The payload sent in a 200 response depends on the request method. For the methods defined 
  by this specification, the intended meaning of the payload can be summarized as:
  
  Aside from responses to CONNECT, a 200 response always has a payload, though an origin 
  server MAY generate a payload body of zero length. If no payload is desired, an origin
  server ought to send 204 No Content instead. For CONNECT, no payload is allowed because
  the successful result is a tunnel, which begins immediately after the 200 response 
  header section.
  
  A 200 response is cacheable by default; i.e., unless otherwise indicated by the method
  definition or explicit cache controls.'
        `,
	},

	201: {
		"message": "Created",
		"description": `
  The request has been fulfilled and has resulted in one or more new resources being created.
  
  The primary resource created by the request is identified by either a Location header field
  in the response or, if no Location field is received, by the effective request URI. The 201 
  response payload typically describes and links to the resource(s) created. 

  See Section 7.2 of RFC7231 for a discussion of the meaning and purpose of validator header 
  fields, such as ETag and Last-Modified, in a 201 response.
        `,
	},

	202: {
		"message": "Accepted",
		"description": `
  The request has been accepted for processing, but the processing has not been completed.
  
  The request might or might not eventually be acted upon, as it might be disallowed when processing
  actually takes place. There is no facility in HTTP for re-sending a status code from an asynchronous
  operation.
  
  The 202 response is intentionally noncommittal. Its purpose is to allow a server to accept a
  request for some other process (perhaps a batch-oriented process that is only run once perday)
  without requiring that the user agent's connection to the server persist until the process is 
  completed. 
  
  The representation sent with this response ought to describe the request's current status and point to
  (or embed) a status monitor that can provide the user with an estimate of when the request will be
  fulfilled.'
        `,
	},

	203: {
		"message": "Non-Authoritative Information",
		"description": `
  The request was successful but the enclosed payload has been modified from that of the origin server's 
  200 OK response by a transforming proxy.
  
  This status code allows the proxy to notify recipients when a transformation has been applied, since that
  knowledge might impact later decisions regarding the content. For example, future cache validation requests
  for the content might only be applicable along the same request path (through the same proxies). The 203
  response is similar to the Warning code of 214 Transformation Applied 2, which has the advantage of
  being applicable to responses with any status code.
  
  A 203 response is cacheable by default; i.e., unless otherwise indicated by the method definition or explicit
  cache controls.
        `,
	},

	204: {
		"message": "No Content",
		"description": `
  The server has successfully fulfilled the request and that there is no additional content to send in the response
  payload body. Metadata in the response header fields refer to the target resource and its selected representation
  after the requested action was applied.
  
  For example, if a 204 status code is received in response to a PUT request and the response contains an ETag header
  field, then the PUT was successful and the ETag field-value contains the entity-tag for the new representation of that
  target resource. The 204 response allows a server to indicate that the action has been successfully applied to the 
  target resource, while implying that the user agent does not need to traverse away from its current "document
  view" (if any).
  
  The server assumes that the user agent will provide some indication of the success to its user, in accord with its own
  interface, and apply any new or updated metadata in the response to its active representation.
  
  For example, a 204 status code is commonly used with document editing interfaces corresponding to a "save" action, such
  that the document being saved remains available to the user for editing.
  
  It is also frequently used with interfaces that expect automated data transfers to be prevalent, such as within distributed
  version control systems. A 204 response is terminated by the first empty line after the header fields because it cannot contain 
  a message body.
  
  A 204 response is cacheable by default; i.e., unless otherwise indicated by the method definition or explicit cache controls.
        `,
	},

	205: {
		"message": "Reset Content",
		"description": `
  The server has fulfilled the request and desires that the user agent reset the "document view", which caused the request
  to be sent, to its original state as received from the origin server. This response is intended to support a common data
  entry use case where the user receives content that supports data entry (a form, notepad, canvas, etc.), enters or 
  manipulates data in that space, causes the entered data to be submitted in a request and then the data entry mechanism is
  reset for the next entry so that the user can easily initiate another input action.
  
  Since the 205 status code implies that no additional content will be provided, a server MUST NOT generate a payload in a
  205 response. In other words, a server MUST do one of the following for a 205 response: a) indicate a zero-length body for 
  the response by including a Content-Length header field with a value of 0; b) indicate a zero-length payload for the response
  by including a Transfer-Encoding header field with a value of chunked and a message body consisting of a single chunk of
  zero-length; or, c) close the connection immediately after sending the blank line terminating the header section.'
        `,
	},

	206: {
		"message":     "Continue",
		"description": "",
	},

	207: {
		"message":     "Continue",
		"description": "",
	},

	208: {
		"message":     "Continue",
		"description": "",
	},

	226: {
		"message":     "Continue",
		"description": "",
	},

	// 3xx Redirection
	300: {
		"message":     "Continue",
		"description": "",
	},

	301: {
		"message":     "Continue",
		"description": "",
	},

	302: {
		"message":     "Continue",
		"description": "",
	},

	303: {
		"message":     "Continue",
		"description": "",
	},

	304: {
		"message":     "Continue",
		"description": "",
	},

	305: {
		"message": "Use Proxy",
		"description": `
  Defined in a previous version of this specification and is now deprecated, due to security concerns regarding in-band 
  configuration of a proxy.
        `,
	},

	306: {
		"message":     "Continue",
		"description": "",
	},

	307: {
		"message":     "Continue",
		"description": "",
	},

	308: {
		"message":     "Continue",
		"description": "",
	},

	//4xx Client Error

	400: {
		"message":     "Continue",
		"description": "",
	},

	401: {
		"message":     "Continue",
		"description": "",
	},

	402: {
		"message":     "Continue",
		"description": "",
	},

	403: {
		"message":     "Continue",
		"description": "",
	},

	404: {
		"message":     "Continue",
		"description": "",
	},

	405: {
		"message":     "Continue",
		"description": "",
	},

	406: {
		"message":     "Continue",
		"description": "",
	},

	407: {
		"message":     "Continue",
		"description": "",
	},

	408: {
		"message":     "Continue",
		"description": "",
	},

	409: {
		"message":     "Continue",
		"description": "",
	},

	410: {
		"message":     "Continue",
		"description": "",
	},

	411: {
		"message":     "Continue",
		"description": "",
	},

	412: {
		"message":     "Continue",
		"description": "",
	},

	413: {
		"message":     "Continue",
		"description": "",
	},

	414: {
		"message":     "Continue",
		"description": "",
	},

	415: {
		"message":     "Continue",
		"description": "",
	},

	416: {
		"message":     "Continue",
		"description": "",
	},

	417: {
		"message":     "Continue",
		"description": "",
	},

	418: {
		"message":     "Continue",
		"description": "",
	},

	422: {
		"message":     "Continue",
		"description": "",
	},

	423: {
		"message":     "Continue",
		"description": "",
	},

	424: {
		"message":     "Continue",
		"description": "",
	},

	426: {
		"message":     "Continue",
		"description": "",
	},

	428: {
		"message":     "Continue",
		"description": "",
	},

	429: {
		"message":     "Continue",
		"description": "",
	},

	431: {
		"message":     "Continue",
		"description": "",
	},

	451: {
		"message":     "Continue",
		"description": "",
	},

	// 5xx Server Error

	500: {
		"message":     "Internal Server Error",
		"description": "Hello",
	},

	501: {
		"message":     "Continue",
		"description": "",
	},

	502: {
		"message": "Bad Gateway",
		"description": `
  The server, while acting as a gateway or proxy, received an invalid response from an inbound server it accessed while 
  attempting to fulfill the request.'
        `,
	},

	503: {
		"message":     "Continue",
		"description": "",
	},

	504: {
		"message":     "Continue",
		"description": "",
	},

	505: {
		"message":     "Continue",
		"description": "",
	},

	506: {
		"message":     "Continue",
		"description": "",
	},

	507: {
		"message":     "Continue",
		"description": "",
	},

	508: {
		"message":     "Continue",
		"description": "",
	},

	510: {
		"message":     "Continue",
		"description": "",
	},

	511: {
		"message":     "Continue",
		"description": "",
	},
}

func statusMessage(code int) string {
	return HTTPStatusCodes[code]["message"]
}

func statusDescription(code int) string {
	return HTTPStatusCodes[code]["description"]
}
