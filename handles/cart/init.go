package cart

import "weshop/server"

func init() {
	// todo test
	server.RegisterGetHandleNoUserID("/cart/index", indexHandle)
	server.RegisterPostHandleNoUserID("/cart/add", addHandle)
	server.RegisterGetHandleNoUserID("/cart/count", countHandle)
	server.RegisterPostHandleNoUserID("/cart/checked", checkedHandle)
	server.RegisterPostHandleNoUserID("/cart/checkedall", checkedallHandle)
	server.RegisterPostHandleNoUserID("/cart/update", updateHandle)
	server.RegisterPostHandleNoUserID("/cart/delete", deleteHandle)
}
