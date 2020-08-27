package common

import (
	"io"
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

	// AddHook will add a hook function to be ran after the context has completed
	AddHook(func(statusCode int, storage map[string]string))

	// GetRequest will return http.Request
	GetRequest() (req *http.Request)

	// GetWriter will return response writer.
	GetWriter() (writer http.ResponseWriter)
}

// Handler is the HTTP handler type
type Handler func(ctx *Context) Response

// Response is the http response handlers return
type Response interface {
	StatusCode() (code int)
	ContentType() (contentType string)
	WriteTo(w io.Writer) (n int64, err error)
}
