package shop

import "weshop/server"

func init() {
	server.RegisterGetHandleNoUserID("/shop/checkout", checkoutHandle)
	server.RegisterPostHandleNoUserID("/shop/address/save", addressSaveHandle)
	server.RegisterGetHandleNoUserID("/shop/address/list", addressListHandle)
	server.RegisterPostHandleNoUserID("/shop/order/create", orderCreateHandle)
}
