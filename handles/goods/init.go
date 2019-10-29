package goods

import "weshop/server"

func init() {
	server.RegisterGetHandleNoUserID("/goods/count", countHandle)
	server.RegisterPostHandleNoUserID("/goods/list", listHandle)
	server.RegisterPostHandleNoUserID("/goods/detail", detailHandle)
}
