package common

import (
	"net/http"
)

// Plugins represents the plugins library interface. This allows for plugins and libraries
// to not be affected by vroomy/plugins version changes.
type Plugins interface {
	Backend(key string, val interface{}) error
}

// Context is the http context plugins need to support handlers
type Context interface {
	// Write will write a byteslice
	Write(bs []byte) (n int, err error)

	// WriteString will write a string
	WriteString(str string) (n int, err error)

	// Param will return the associated parameter value with the provided key
	Param(key string) (value string)

	// Get will retrieve a value for a provided key from the Context's internal storage
	Get(key string) (value string)

	// Put will set a value for a provided key into the Context's internal storage
	Put(key, value string)

	// BindJSON is a helper function which binds the request body to a provided value to be parsed as JSON
	BindJSON(value interface{}) (err error)

	// GetRequest will return http.Request
	GetRequest() (req *http.Request)

	// GetWriter will return response writer.
	GetWriter() (writer http.ResponseWriter)
}

// Handler is the HTTP handler type
type Handler func(ctx Context) *Response

// Storage is used as a basic form of Param storage for a Context
// TODO: Determine with team if it seems valuable to change this to map[string]interface{}.
// I'd prefer if we can keep it as-is, due to the fact that map[string]string has much less
// GC overhead. Additionally, avoiding type assertion would be fantastic.
type Storage map[string]string

// Hook is a function called after the response has been completed to the requester
type Hook func(statusCode int, storage map[string]string)
