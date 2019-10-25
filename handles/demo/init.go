package demo

import "weshop/server"

func init() {
	server.RegisterGetHandle("/demon/hello", helloHandle)
}
