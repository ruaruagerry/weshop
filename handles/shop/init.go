package shop

import "weshop/server"

func init() {
	server.RegisterGetHandleNoUserID("/shop/checkout", checkoutHandle)
	server.RegisterPostHandleNoUserID("/shop/address/save", addressSaveHandle)
	server.RegisterGetHandleNoUserID("/shop/address/list", addressListHandle)
	server.RegisterPostHandleNoUserID("/shop/address/delete", addressDeleteHandle)
	server.RegisterPostHandleNoUserID("/shop/order/create", orderCreateHandle)
	server.RegisterGetHandleNoUserID("/shop/order/list", orderListHandle)
	server.RegisterPostHandleNoUserID("/shop/order/detail", orderDetailHandle)
	server.RegisterPostHandleNoUserID("/shop/order/cancel", orderCancelHandle)
	server.RegisterPostHandleNoUserID("/shop/collect/add", collectAddHandle)
	server.RegisterPostHandleNoUserID("/shop/collect/delete", collectDeleteHandle)
	server.RegisterGetHandleNoUserID("/shop/collect/list", collectListHandle)
}
