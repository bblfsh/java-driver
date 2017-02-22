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
		TokenKeys: map[string]bool{
			"identifier":        true, // SimpleName
			"escapedValue":      true, // StringLiteral
			"keyword":           true, // Modifier
			"primitiveTypeCode": true, // ?
		},
		SyntheticTokens: map[string]string{
			"PackageDeclaration": "package",
		},
	}
}

var annotationRules uast.Rule = uast.Rules(
	uast.OnInternalType("PackageDeclaration").Role(uast.PackageDeclaration),
	uast.OnInternalType("MethodDeclaration").Role(uast.FunctionDeclaration),
	uast.OnInternalType("ImportDeclaration").Role(uast.ImportDeclaration),
	uast.OnInternalType("ImportDeclaration", "QualifiedName").Role(uast.ImportPath),
	uast.OnInternalType("QualifiedName").Role(uast.QualifiedIdentifier),
	uast.OnInternalType("SimpleName").Role(uast.SimpleIdentifier),
	uast.OnInternalType("IfStatement").Role(uast.IfStatement),
)

// Annotate annotates the given UAST.
func Annotate(n *uast.Node) error {
	return uast.PreOrderVisit(n, annotationRules)
}
