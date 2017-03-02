package normalizer

import (
	"github.com/bblfsh/java-driver/driver/normalizer/jdt"

	"github.com/bblfsh/sdk/uast"
)

// NewToNoder creates a new uast.ToNoder to convert
// Java ASTs to UAST.
func NewToNoder() uast.ToNoder {
	return &uast.BaseToNoder{
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
	uast.OnInternalType(jdt.CompilationUnit).Role(uast.File),
	uast.OnInternalType(jdt.PackageDeclaration).Role(uast.PackageDeclaration),
	uast.OnInternalType(jdt.MethodDeclaration).Role(uast.FunctionDeclaration),
	uast.OnInternalType(jdt.ImportDeclaration).Role(uast.ImportDeclaration),
	uast.OnInternalType(jdt.TypeDeclaration).Role(uast.TypeDeclaration),
	uast.OnInternalType(jdt.ImportDeclaration, jdt.QualifiedName).Role(uast.ImportPath),
	uast.OnInternalType(jdt.QualifiedName).Role(uast.QualifiedIdentifier),
	uast.OnInternalType(jdt.SimpleName).Role(uast.SimpleIdentifier),
	uast.OnInternalType(jdt.Block).Role(uast.BlockScope, uast.Block),
	uast.OnInternalType(jdt.ExpressionStatement).Role(uast.Statement),
	uast.OnInternalType(jdt.ReturnStatement).Role(uast.Return, uast.Statement),
	uast.OnInternalType(jdt.MethodInvocation).Role(uast.MethodInvocation),
	uast.OnInternalType(jdt.IfStatement).Role(uast.If, uast.Statement),
	uast.OnInternalRole("elseStatement").Role(uast.IfElse, uast.Statement),
	uast.OnPath(uast.OnInternalType(jdt.Assignment)).Role(uast.Assignment),
	uast.OnPath(uast.OnInternalType(jdt.Assignment), uast.OnInternalRole("leftHandSide")).Role(uast.AssignmentVariable),
	uast.OnPath(uast.OnInternalType(jdt.Assignment), uast.OnInternalRole("rightHandSide")).Role(uast.AssignmentValue),
	//TODO: IfBody, IfCondition
	uast.OnInternalType(jdt.NullLiteral).Role(uast.NullLiteral, uast.Literal),
	uast.OnInternalType(jdt.StringLiteral).Role(uast.StringLiteral, uast.Literal),
	uast.OnInternalType(jdt.NumberLiteral).Role(uast.NumberLiteral, uast.Literal),
	uast.OnInternalType(jdt.TypeLiteral).Role(uast.TypeLiteral, uast.Literal),
	uast.OnInternalType(jdt.ThisExpression).Role(uast.This, uast.Expression),
	//TODO: synchronized
	//TODO: try-with-resources
	uast.OnInternalType(jdt.Javadoc).Role(uast.Documentation, uast.Comment),
)

// Annotate annotates the given Java UAST.
func Annotate(n *uast.Node) error {
	return uast.PreOrderVisit(n, AnnotationRules)
}
