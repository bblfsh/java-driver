package normalizer

import (
	"github.com/bblfsh/java-driver/driver/normalizer/jdt"

	. "github.com/bblfsh/sdk/uast"
	. "github.com/bblfsh/sdk/uast/ann"
)

var AnnotationRules = On(Any).Self(
	On(Not(jdt.CompilationUnit)).Error("root must be CompilationUnit"),
	On(jdt.CompilationUnit).Roles(File).Descendants(
		// Names
		On(jdt.QualifiedName).Roles(QualifiedIdentifier),
		On(jdt.SimpleName).Roles(SimpleIdentifier),

		// Visibility
		On(Or(jdt.MethodDeclaration, jdt.TypeDeclaration)).Self(
			On(HasChild(And(jdt.Modifier, jdt.KeywordPublic))).Roles(VisibleFromWorld),
			On(HasChild(And(jdt.Modifier, jdt.KeywordPrivate))).Roles(VisibleFromType),
			On(HasChild(And(jdt.Modifier, jdt.KeywordProtected))).Roles(VisibleFromSubtype),
			On(Not(HasChild(And(jdt.Modifier,
				Or(jdt.KeywordPublic, jdt.KeywordPrivate, jdt.KeywordProtected),
			)))).Roles(VisibleFromPackage),
		),

		// Package and imports
		On(jdt.PackageDeclaration).Roles(PackageDeclaration),
		On(jdt.ImportDeclaration).Roles(ImportDeclaration).Children(
			On(jdt.QualifiedName).Roles(ImportPath),
		),

		// Type declarations
		On(jdt.TypeDeclaration).Roles(TypeDeclaration),

		// Method declarations
		On(jdt.MethodDeclaration).Roles(FunctionDeclaration).Children(
			//TODO: On(jdt.PropertyTypeParameters).Roles(FunctionDeclarationTypeParameter),
			//TODO: On(jdt.PropertyReturnType2).Roles(FunctionDeclarationReturnType)
			On(jdt.PropertyName).Roles(FunctionDeclarationName),
			On(jdt.PropertyBody).Roles(FunctionDeclarationBody),
			On(jdt.PropertyParameters).Roles(FunctionDeclarationArgument).Self(
				On(HasProperty("varargs", "true")).Roles(FunctionDeclarationVarArgsList),
			).Children(
				//TODO: On(jdt.PropertyType).Roles(FunctionDeclarationArgumentType),
				On(jdt.PropertyName).Roles(FunctionDeclarationArgumentName),
			),
		),

		// Literals
		On(jdt.BooleanLiteral).Roles(BooleanLiteral),
		On(jdt.CharacterLiteral).Roles(CharacterLiteral),
		On(jdt.NullLiteral).Roles(NullLiteral),
		On(jdt.NumberLiteral).Roles(NumberLiteral),
		On(jdt.StringLiteral).Roles(StringLiteral),
		On(jdt.TypeLiteral).Roles(TypeLiteral),

		// Calls
		On(jdt.MethodInvocation).Roles(Call).Children(
			On(jdt.PropertyExpression).Roles(CallReceiver),
			On(jdt.PropertyName).Roles(CallCallee),
			On(jdt.PropertyArguments).Roles(CallPositionalArgument),
		),

		// Conditionals
		On(jdt.IfStatement).Roles(If, Statement).Children(
			On(jdt.PropertyExpression).Roles(IfCondition),
			On(jdt.PropertyThenStatement).Roles(IfBody),
			On(jdt.PropertyElseExpression).Roles(IfElse),
		),

		On(jdt.SwitchStatement).Roles(Switch, Statement).Children(
			//TODO: On(jdt.PropertyExpression).Roles(SwitchExpression),
			On(jdt.SwitchCase).Roles(SwitchCase),
		),

		On(jdt.InfixExpression).Self(
			On(HasProperty("operator", "+")).Roles(OpAdd),
			On(HasProperty("operator", "-")).Roles(OpSubstract),
			On(HasProperty("operator", "*")).Roles(OpMultiply),
			On(HasProperty("operator", "/")).Roles(OpDivide),
			On(HasProperty("operator", "%")).Roles(OpMod),
			On(HasProperty("operator", "<<")).Roles(OpBitwiseLeftShift),
			On(HasProperty("operator", ">>")).Roles(OpBitwiseRightShift),
			On(HasProperty("operator", ">>>")).Roles(OpBitwiseUnsignedRightShift),
			On(HasProperty("operator", "&")).Roles(OpBitwiseAnd),
			On(HasProperty("operator", "|")).Roles(OpBitwiseOr),
			On(HasProperty("operator", "&&")).Roles(OpBooleanAnd),
			On(HasProperty("operator", "||")).Roles(OpBooleanOr),
			On(HasProperty("operator", "^")).Roles(OpBooleanXor),
		),

		On(jdt.PostfixExpression).Self(
			On(HasProperty("operator", "++")).Roles(OpPostIncrement),
			On(HasProperty("operator", "--")).Roles(OpPostDecrement),
		),

		On(jdt.PrefixExpression).Self(
			On(HasProperty("operator", "++")).Roles(OpPreIncrement),
			On(HasProperty("operator", "--")).Roles(OpPreDecrement),
			On(HasProperty("operator", "+")).Roles(OpPositive),
			On(HasProperty("operator", "-")).Roles(OpNegative),
			On(HasProperty("operator", "~")).Roles(OpBitwiseComplement),
			On(HasProperty("operator", "!")).Roles(OpBooleanNot),
		),

		On(jdt.Assignment).Roles(Assignment).Children(
			On(jdt.PropertyLeftHandSide).Roles(AssignmentVariable),
			On(jdt.PropertyRightHandSide).Roles(AssignmentValue),
		).Self(
			On(Not(HasProperty("operator", "="))).Roles(AugmentedAssignmentOperator, AugmentedAssignment).Self(
				On(HasProperty("operator", "+=")).Roles(OpAdd),
				On(HasProperty("operator", "-=")).Roles(OpSubstract),
				On(HasProperty("operator", "*=")).Roles(OpMultiply),
				On(HasProperty("operator", "/=")).Roles(OpDivide),
				On(HasProperty("operator", "%=")).Roles(OpMod),
				On(HasProperty("operator", "&=")).Roles(OpBitwiseAnd),
				On(HasProperty("operator", "|=")).Roles(OpBitwiseOr),
				On(HasProperty("operator", "^=")).Roles(OpBooleanXor),
				On(HasProperty("operator", "<<=")).Roles(OpBitwiseLeftShift),
				On(HasProperty("operator", ">>=")).Roles(OpBitwiseRightShift),
				On(HasProperty("operator", ">>>=")).Roles(OpBitwiseUnsignedRightShift),
			),
		),

		// Others
		On(jdt.Block).Roles(BlockScope, Block),
		On(jdt.ExpressionStatement).Roles(Statement),
		On(jdt.ReturnStatement).Roles(Return, Statement),

		On(jdt.ThisExpression).Roles(This, Expression),
		//TODO: synchronized
		//TODO: try-with-resources
		On(jdt.Javadoc).Roles(Documentation, Comment),
	),
)
