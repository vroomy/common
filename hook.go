package common

// Hook is a function called after the response has been completed to the requester
type Hook func(statusCode int, ctx Context)
