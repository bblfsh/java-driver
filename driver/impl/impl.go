package impl

import "gopkg.in/bblfsh/sdk.v2/sdk/driver"

func init() {
	// Can be overridden when native Go implementation is available
	driver.DefaultDriver = driver.NewExecDriver()
}
