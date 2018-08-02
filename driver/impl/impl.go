package impl

import (
	"gopkg.in/bblfsh/sdk.v2/driver/native"
	"gopkg.in/bblfsh/sdk.v2/driver/server"
)

func init() {
	// Can be overridden when native Go implementation is available
	server.DefaultDriver = native.NewDriver(native.UTF8)
}
