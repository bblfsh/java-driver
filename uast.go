package java

import (
	"github.com/bblfsh/sdk/uast"
)

func NewOriginalToNoder() uast.OriginalToNoder {
	return &uast.BaseOriginalToNoder{
		InternalTypeKey: "internalClass",
	}
}
