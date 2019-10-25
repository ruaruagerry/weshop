package goods

import "weshop/server"

func init() {
	server.RegisterGetHandleNoUserID("/goods/count", countHandle)
}
