package catalog

import "weshop/server"

func init() {
	// todo test
	server.RegisterGetHandle("/catalog/index", indexHandle)
	server.RegisterPostHandle("/catalog/current", currentHandle)
}
