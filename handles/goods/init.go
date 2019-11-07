package goods

import "weshop/server"

func init() {
	server.RegisterGetHandle("/goods/count", countHandle)
	server.RegisterPostHandle("/goods/list", listHandle)
	server.RegisterPostHandle("/goods/detail", detailHandle)
}
