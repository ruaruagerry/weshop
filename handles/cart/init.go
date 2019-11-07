package cart

import "weshop/server"

func init() {
	// todo test
	server.RegisterGetHandle("/cart/index", indexHandle)
	server.RegisterPostHandle("/cart/add", addHandle)
	server.RegisterGetHandle("/cart/count", countHandle)
	server.RegisterPostHandle("/cart/checked", checkedHandle)
	server.RegisterPostHandle("/cart/checkedall", checkedallHandle)
	server.RegisterPostHandle("/cart/update", updateHandle)
	server.RegisterPostHandle("/cart/delete", deleteHandle)
}
