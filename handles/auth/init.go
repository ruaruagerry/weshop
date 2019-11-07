package auth

import "weshop/server"

func init() {
	server.RegisterPostHandleNoUserID("/auth/login", loginHandle)
}
