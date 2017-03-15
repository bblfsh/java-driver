package normalizer

import (
	"github.com/bblfsh/java-driver/driver/normalizer/jdt"

	. "github.com/bblfsh/sdk/uast"
	. "github.com/bblfsh/sdk/uast/ann"
)

var AnnotationRules = On(Any).Self(
	On(Not(HasInternalType(jdt.CompilationUnit))).Error("root must be CompilationUnit"),
	On(HasInternalType(jdt.CompilationUnit)).Roles(File).Descendants(
		On(HasInternalType(jdt.PackageDeclaration)).Roles(PackageDeclaration),
		On(HasInternalType(jdt.MethodDeclaration)).Roles(FunctionDeclaration),
		On(HasInternalType(jdt.ImportDeclaration)).Roles(ImportDeclaration).Children(
			On(HasInternalType(jdt.QualifiedName)).Roles(ImportPath),
		),
		On(HasInternalType(jdt.TypeDeclaration)).Roles(TypeDeclaration),
		On(HasInternalType(jdt.QualifiedName)).Roles(QualifiedIdentifier),
		On(HasInternalType(jdt.SimpleName)).Roles(SimpleIdentifier),
		On(HasInternalType(jdt.Block)).Roles(BlockScope, Block),
		On(HasInternalType(jdt.ExpressionStatement)).Roles(Statement),
		On(HasInternalType(jdt.ReturnStatement)).Roles(Return, Statement),
		On(HasInternalType(jdt.MethodInvocation)).Roles(MethodInvocation),
		On(HasInternalType(jdt.IfStatement)).Roles(If, Statement),
		On(HasInternalRole("elseStatement")).Roles(IfElse, Statement),
		On(HasInternalType(jdt.Assignment)).Roles(Assignment).Children(
			On(HasInternalRole("leftHandSide")).Roles(AssignmentVariable),
			On(HasInternalRole("rightHandSide")).Roles(AssignmentValue),
		),
		//TODO: IfBody, IfCondition
		On(HasInternalType(jdt.NullLiteral)).Roles(NullLiteral, Literal),
		On(HasInternalType(jdt.StringLiteral)).Roles(StringLiteral, Literal),
		On(HasInternalType(jdt.NumberLiteral)).Roles(NumberLiteral, Literal),
		On(HasInternalType(jdt.TypeLiteral)).Roles(TypeLiteral, Literal),
		On(HasInternalType(jdt.ThisExpression)).Roles(This, Expression),
		//TODO: synchronized
		//TODO: try-with-resources
		On(HasInternalType(jdt.Javadoc)).Roles(Documentation, Comment),
	),
)
