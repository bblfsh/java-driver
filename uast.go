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
			"NullLiteral":        "null",
		},
		//TODO: add names of children (e.g. elseStatement) as
		//      children node properties.
	}
}

// AnnotationRules for Java UAST.
var AnnotationRules uast.Rule = uast.Rules(
	uast.OnInternalType("CompilationUnit").Role(uast.File),
	uast.OnInternalType("PackageDeclaration").Role(uast.PackageDeclaration),
	uast.OnInternalType("MethodDeclaration").Role(uast.FunctionDeclaration),
	uast.OnInternalType("ImportDeclaration").Role(uast.ImportDeclaration),
	uast.OnInternalType("TypeDeclaration").Role(uast.TypeDeclaration),
	uast.OnInternalType("ImportDeclaration", "QualifiedName").Role(uast.ImportPath),
	uast.OnInternalType("QualifiedName").Role(uast.QualifiedIdentifier),
	uast.OnInternalType("SimpleName").Role(uast.SimpleIdentifier),
	uast.OnInternalType("Block").Role(uast.BlockScope, uast.Block),
	uast.OnInternalType("ExpressionStatement").Role(uast.Statement),
	uast.OnInternalType("ReturnStatement").Role(uast.Return, uast.Statement),
	uast.OnInternalType("MethodInvocation").Role(uast.MethodInvocation),
	uast.OnInternalType("IfStatement").Role(uast.If, uast.Statement),
	uast.OnInternalRole("elseStatement").Role(uast.IfElse, uast.Statement),
	uast.OnPath(uast.OnInternalType("Assignment")).Role(uast.Assignment),
	uast.OnPath(uast.OnInternalType("Assignment"), uast.OnInternalRole("leftHandSide")).Role(uast.AssignmentVariable),
	uast.OnPath(uast.OnInternalType("Assignment"), uast.OnInternalRole("rightHandSide")).Role(uast.AssignmentValue),
	//TODO: IfBody, IfCondition
	uast.OnInternalType("NullLiteral").Role(uast.NullLiteral, uast.Literal),
	uast.OnInternalType("StringLiteral").Role(uast.StringLiteral, uast.Literal),
	uast.OnInternalType("NumberLiteral").Role(uast.NumberLiteral, uast.Literal),
	uast.OnInternalType("TypeLiteral").Role(uast.TypeLiteral, uast.Literal),
	uast.OnInternalType("ThisExpression").Role(uast.This, uast.Expression),
	//TODO: synchronized
	//TODO: try-with-resources
	uast.OnInternalType("Javadoc").Role(uast.Documentation, uast.Comment),
)

// Annotate annotates the given Java UAST.
func Annotate(n *uast.Node) error {
	return uast.PreOrderVisit(n, AnnotationRules)
}
