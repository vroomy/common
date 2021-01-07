package common

// Plugins represents the plugins library interface. This allows for plugins and libraries
// to not be affected by vroomy/plugins version changes.
type Plugins interface {
	Backend(key string, val interface{}) error
}

// Handler is the HTTP handler type
type Handler func(ctx Context) *Response

// Storage is used as a basic form of Param storage for a Context
// TODO: Determine with team if it seems valuable to change this to map[string]interface{}.
// I'd prefer if we can keep it as-is, due to the fact that map[string]string has much less
// GC overhead. Additionally, avoiding type assertion would be fantastic.
type Storage map[string]string
