package tool

import "weshop/server"

func init() {
	server.RegisterGetHandleNoUserID("/tool/writeconfig", writeconfigHandle)
}
