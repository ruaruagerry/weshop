package shop

import "weshop/server"

func init() {
	server.RegisterGetHandle("/shop/checkout", checkoutHandle)
	server.RegisterPostHandle("/shop/address/save", addressSaveHandle)
	server.RegisterGetHandle("/shop/address/list", addressListHandle)
	server.RegisterPostHandle("/shop/address/delete", addressDeleteHandle)
	server.RegisterPostHandle("/shop/order/create", orderCreateHandle)
	server.RegisterGetHandle("/shop/order/list", orderListHandle)
	server.RegisterPostHandle("/shop/order/detail", orderDetailHandle)
	server.RegisterPostHandle("/shop/order/cancel", orderCancelHandle)
	server.RegisterPostHandle("/shop/collect/add", collectAddHandle)
	server.RegisterPostHandle("/shop/collect/delete", collectDeleteHandle)
	server.RegisterGetHandle("/shop/collect/list", collectListHandle)
}
