package catagory

import "weshop/server"

func init() {
	server.RegisterGetHandle("/catagory/index", indexHandle)
	server.RegisterPostHandle("/catagory/current", currentHandle)
}
