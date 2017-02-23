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
		},
		//TODO: add names of children (e.g. elseStatement) as
		//      children node properties.
	}
}

var annotationRules uast.Rule = uast.Rules(
	uast.OnInternalType("PackageDeclaration").Role(uast.PackageDeclaration),
	uast.OnInternalType("MethodDeclaration").Role(uast.FunctionDeclaration),
	uast.OnInternalType("ImportDeclaration").Role(uast.ImportDeclaration),
	uast.OnInternalType("ImportDeclaration", "QualifiedName").Role(uast.ImportPath),
	uast.OnInternalType("QualifiedName").Role(uast.QualifiedIdentifier),
	uast.OnInternalType("SimpleName").Role(uast.SimpleIdentifier),
	uast.OnInternalType("IfStatement").Role(uast.If, uast.Statement),
	uast.OnInternalRole("elseStatement").Role(uast.IfElse, uast.Statement),
	//TODO: IfBody, IfCondition
)

// Annotate annotates the given UAST.
func Annotate(n *uast.Node) error {
	return uast.PreOrderVisit(n, annotationRules)
}
