package impl

import (
	"github.com/bblfsh/sdk/v3/driver/native"
	"github.com/bblfsh/sdk/v3/driver/server"
)

func init() {
	// Can be overridden when native Go implementation is available
	server.DefaultDriver = native.NewDriver(native.UTF8)
}
