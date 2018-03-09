package normalizer

import (
	"strings"

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

	Modifier: func(n map[string]interface{}) error {
		// Remove //, /*...*/ and /**..*/ from comment nodes
		if t, ok := n["internalClass"]; ok {
			switch t {
			case "LineComment":
				if text, ok := n["text"].(string); ok && strings.HasPrefix(text, "//") {
					n["text"] = text[2:]
				}
			case "BlockComment":
				if text, ok := n["text"].(string); ok && strings.HasPrefix(text, "/*") {
					n["text"] = text[2 : len(text)-2]
				}
			case "Javadoc":
				if text, ok := n["text"].(string); ok && strings.HasPrefix(text, "/**") {
					n["text"] = text[3 : len(text)-2]
				}
			}
		}

		return nil
	},
	//TODO: add names of children (e.g. elseStatement) as
	//      children node properties.
}
