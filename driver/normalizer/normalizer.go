package normalizer

import (
	"github.com/bblfsh/java-driver/driver/normalizer/jdt"

	. "github.com/bblfsh/sdk/uast"
	. "github.com/bblfsh/sdk/uast/ann"
)

var AnnotationRules = On(Any).Self(
	On(Not(jdt.CompilationUnit)).Error("root must be CompilationUnit"),
	On(jdt.CompilationUnit).Roles(File).Descendants(
		On(Or(jdt.MethodDeclaration, jdt.TypeDeclaration)).Self(
			On(HasChild(And(jdt.Modifier, HasToken("public")))).Roles(VisibleFromWorld),
			On(HasChild(And(jdt.Modifier, HasToken("private")))).Roles(VisibleFromType),
			On(HasChild(And(jdt.Modifier, HasToken("protected")))).Roles(VisibleFromSubtype),
			On(Not(HasChild(And(jdt.Modifier,
				Or(HasToken("public"), HasToken("private"), HasToken("protected")),
			)))).Roles(VisibleFromPackage),
		),
		On(jdt.PackageDeclaration).Roles(PackageDeclaration),
		On(jdt.ImportDeclaration).Roles(ImportDeclaration).Children(
			On(jdt.QualifiedName).Roles(ImportPath),
		),
		On(jdt.TypeDeclaration).Roles(TypeDeclaration),
		On(jdt.QualifiedName).Roles(QualifiedIdentifier),
		On(jdt.SimpleName).Roles(SimpleIdentifier),
		On(jdt.Block).Roles(BlockScope, Block),
		On(jdt.ExpressionStatement).Roles(Statement),
		On(jdt.ReturnStatement).Roles(Return, Statement),
		On(jdt.MethodInvocation).Roles(Call).Children(
			On(HasInternalRole("expression")).Roles(CallReceiver),
			On(HasInternalRole("name")).Roles(CallCallee),
			On(HasInternalRole("arguments")).Roles(CallPositionalArgument),
		),
		On(jdt.IfStatement).Roles(If, Statement),
		On(HasInternalRole("elseStatement")).Roles(IfElse, Statement),
		On(jdt.Assignment).Roles(Assignment).Children(
			On(HasInternalRole("leftHandSide")).Roles(AssignmentVariable),
			On(HasInternalRole("rightHandSide")).Roles(AssignmentValue),
		),
		//TODO: IfBody, IfCondition
		On(jdt.NullLiteral).Roles(NullLiteral, Literal),
		On(jdt.StringLiteral).Roles(StringLiteral, Literal),
		On(jdt.NumberLiteral).Roles(NumberLiteral, Literal),
		On(jdt.TypeLiteral).Roles(TypeLiteral, Literal),
		On(jdt.ThisExpression).Roles(This, Expression),
		//TODO: synchronized
		//TODO: try-with-resources
		On(jdt.Javadoc).Roles(Documentation, Comment),
	),
)
