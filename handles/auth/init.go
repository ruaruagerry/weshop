package auth

import "weshop/server"

func init() {
	server.RegisterGetHandle("/auth/login", loginHandle)
}
