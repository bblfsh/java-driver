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
			On(HasChild(And(jdt.Modifier, jdt.KeywordPublic))).Roles(VisibleFromWorld),
			On(HasChild(And(jdt.Modifier, jdt.KeywordPrivate))).Roles(VisibleFromType),
			On(HasChild(And(jdt.Modifier, jdt.KeywordProtected))).Roles(VisibleFromSubtype),
			On(Not(HasChild(And(jdt.Modifier,
				Or(jdt.KeywordPublic, jdt.KeywordPrivate, jdt.KeywordProtected),
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
			On(jdt.PropertyExpression).Roles(CallReceiver),
			On(jdt.PropertyName).Roles(CallCallee),
			On(jdt.PropertyArguments).Roles(CallPositionalArgument),
		),
		On(jdt.IfStatement).Roles(If, Statement),
		On(jdt.PropertyElseExpression).Roles(IfElse, Statement),
		On(jdt.Assignment).Roles(Assignment).Children(
			On(jdt.PropertyLeftHandSide).Roles(AssignmentVariable),
			On(jdt.PropertyRightHandSide).Roles(AssignmentValue),
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
