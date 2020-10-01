package common

// NewResponse will return a new Response
func NewResponse(statusCode int, contentType string, value interface{}) *Response {
	r := makeResponse(statusCode, contentType, value)
	return &r
}

// Response determines how the server will respond
type Response struct {
	StatusCode  int
	ContentType string
	Value       interface{}

	// Optional fields used by a minority of responses
	Adopted  bool
	Callback string
}

func makeResponse(statusCode int, contentType string, value interface{}) (r Response) {
	r.StatusCode = statusCode
	r.ContentType = contentType
	r.Value = value
	return
}
