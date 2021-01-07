package common

// Plugins represents the plugins library interface. This allows for plugins and libraries
// to not be affected by vroomy/plugins version changes.
type Plugins interface {
	Backend(key string, val interface{}) error
}
