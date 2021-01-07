package common

import "net/http"

// Context is the http context plugins need to support handlers
type Context interface {
	// Bind is a helper function which binds the request body to a provided value to be parsed as JSON
	Bind(value interface{}) (err error)
	// Param will return the associated parameter value with the provided key
	Param(key string) (value string)
	// AddHook will append a new hook to the Context
	AddHook(hook Hook)

	// Get will retrieve a value for a provided key from the Context's internal storage
	Get(key string) (value string)
	// Put will set a value for a provided key into the Context's internal storage
	Put(key, value string)

	// WriteString will write as a string value
	WriteString(status int, contentType, str string)
	// WriteBytes will write as a byteslice value
	WriteBytes(status int, contentType string, str []byte)
	// WriteJSON will write as JSON
	WriteJSON(status int, value interface{})

	// Request will return the underlying http.Request
	Request() (req *http.Request)
	// Writer will return the underlying response writer.
	Writer() (writer http.ResponseWriter)
}
