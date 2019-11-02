package shop

import "weshop/server"

func init() {
	server.RegisterGetHandleNoUserID("/shop/checkout", checkoutHandle)
}
