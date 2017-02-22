package java

import (
	"github.com/bblfsh/sdk/uast"
)

// NewOriginalToNoder creates a new uast.OriginalToNoder to convert
// Java ASTs to UAST.
func NewOriginalToNoder() uast.OriginalToNoder {
	return &uast.BaseOriginalToNoder{
		InternalTypeKey: "internalClass",
		LineKey:         "line",
		OffsetKey:       "startPosition",
	}
}

var typeToRoleTable map[string][]uast.Role = map[string][]uast.Role{
	"PackageDeclaration": []uast.Role{uast.PackageDeclaration},
	"MethodDeclaration":  []uast.Role{uast.FunctionDeclaration},
}

// Annotate annotates the given UAST.
func Annotate(n *uast.Node) error {
	return uast.PreOrderVisit(n, annotate)
}

func annotate(n *uast.Node) error {
	roles, ok := typeToRoleTable[n.InternalType]
	if ok {
		n.Roles = append(n.Roles, roles...)
	}

	return nil
}
