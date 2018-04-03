package impl

import "gopkg.in/bblfsh/sdk.v1/sdk/driver"

func init() {
	// Can be overridden when native Go implementation is available
	driver.DefaultDriver = driver.NewExecDriver()
}
