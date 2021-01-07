package common

// Group is a grouping interface
type Group interface {
	GET(route string, hs ...Handler)
	POST(route string, hs ...Handler)
	PUT(route string, hs ...Handler)
	DELETE(route string, hs ...Handler)
	OPTIONS(route string, hs ...Handler)

	Group(route string, hs ...Handler) Group
	Handle(method, route string, hs ...Handler)
}
