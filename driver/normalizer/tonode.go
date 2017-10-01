package normalizer

import (
	"gopkg.in/bblfsh/sdk.v1/uast"
)

// ToNode is an instance of `uast.ObjectToNode`, defining how to transform an
// into a UAST (`uast.Node`).
//
// https://godoc.org/gopkg.in/bblfsh/sdk.v1/uast#ObjectToNode
var ToNode = &uast.ObjectToNode{
	InternalTypeKey: "internalClass",
	LineKey:         "startLine",
	ColumnKey:       "startColumn",
	OffsetKey:       "startPosition",
	EndLineKey:      "endLine",
	EndColumnKey:    "endColumn",
	EndOffsetKey:    "endPosition",

	//TODO: Should this be part of the UAST rules?
	TokenKeys: map[string]bool{
		"identifier":        true, // SimpleName
		"escapedValue":      true, // StringLiteral
		"keyword":           true, // Modifier
		"primitiveTypeCode": true, // ?
	},
	SyntheticTokens: map[string]string{
		"PackageDeclaration": "package",
		"IfStatement":        "if",
		"NullLiteral":        "null",
	},
	//TODO: add names of children (e.g. elseStatement) as
	//      children node properties.
}
