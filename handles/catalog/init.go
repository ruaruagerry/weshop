package catalog

import "weshop/server"

func init() {
	// todo test
	server.RegisterGetHandleNoUserID("/catalog/index", indexHandle)
	server.RegisterPostHandleNoUserID("/catalog/current", currentHandle)
}
