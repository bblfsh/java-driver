package normalizer

import (
	"github.com/bblfsh/java-driver/driver/normalizer/jdt"

	. "github.com/bblfsh/sdk/uast"
	. "github.com/bblfsh/sdk/uast/ann"
)

var AnnotationRules = On(jdt.CompilationUnit).Roles(File).Descendants(
	// Names
	On(jdt.QualifiedName).Roles(QualifiedIdentifier, Expression),
	On(jdt.SimpleName).Roles(SimpleIdentifier, Expression),

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
	On(jdt.AnonymousClassDeclaration).Roles(TypeDeclaration, Expression, Incomplete).Children(
		On(jdt.PropertyBodyDeclarations).Roles(TypeDeclarationBody),
	),
	On(jdt.AnnotationTypeDeclaration).Roles(TypeDeclaration, Incomplete).Children(
		On(jdt.PropertyBodyDeclarations).Roles(TypeDeclarationBody),
	),
	On(jdt.EnumDeclaration).Roles(TypeDeclaration, Incomplete),

	// ClassDeclaration | InterfaceDeclaration
	On(jdt.TypeDeclaration).Roles(TypeDeclaration),
	// Local (TypeDeclaration | EnumDeclaration)
	On(jdt.TypeDeclarationStatement).Roles(TypeDeclaration, Incomplete),

	// Method declarations
	On(jdt.MethodDeclaration).Roles(FunctionDeclaration).Children(
		On(jdt.PropertyName).Roles(FunctionDeclarationName),
		On(jdt.PropertyBody).Roles(FunctionDeclarationBody),
		On(jdt.PropertyParameters).Roles(FunctionDeclarationArgument).Self(
			On(HasProperty("varargs", "true")).Roles(FunctionDeclarationVarArgsList),
		).Children(
			On(jdt.PropertyName).Roles(FunctionDeclarationArgumentName),
		),
	),
	// FIXME: A lambda expression is not really a function declaration
	// but current UAST doesn't provide anything else for function definitions
	// so I'm considering a lambda expression a function declaration for now
	On(jdt.LambdaExpression).Roles(FunctionDeclaration, Incomplete).Children(
		On(jdt.PropertyBody).Roles(FunctionDeclarationBody),
		On(jdt.PropertyParameters).Roles(FunctionDeclarationArgument).Self(
			On(HasProperty("varargs", "true")).Roles(FunctionDeclarationVarArgsList),
		).Children(
			On(jdt.PropertyName).Roles(FunctionDeclarationArgumentName),
		),
	),

	// Other declarations
	On(jdt.AnnotationTypeMemberDeclaration).Roles(Incomplete),
	On(jdt.EnumConstantDeclaration).Roles(Incomplete),
	On(jdt.FieldDeclaration).Roles(Incomplete),
	On(jdt.Initializer).Roles(Incomplete),
	On(jdt.SingleVariableDeclaration).Roles(Incomplete),
	On(jdt.VariableDeclarationExpression).Roles(Expression, Incomplete),
	On(jdt.VariableDeclarationFragment).Roles(Incomplete),
	On(jdt.VariableDeclarationStatement).Roles(Statement, Incomplete),

	// Literals
	On(jdt.BooleanLiteral).Roles(BooleanLiteral, Expression),
	On(jdt.CharacterLiteral).Roles(CharacterLiteral, Expression),
	On(jdt.NullLiteral).Roles(NullLiteral, Expression),
	On(jdt.NumberLiteral).Roles(NumberLiteral, Expression),
	On(jdt.StringLiteral).Roles(StringLiteral, Expression),
	On(jdt.TypeLiteral).Roles(TypeLiteral, Expression),

	// Calls
	On(jdt.ClassInstanceCreation).Roles(Call, Expression, Incomplete).Children(
		On(jdt.PropertyType).Roles(CallCallee),
		On(jdt.PropertyArguments).Roles(CallPositionalArgument),
	),
	On(jdt.ConstructorInvocation).Roles(Call, Statement, Incomplete).Children(
		On(jdt.PropertyType).Roles(CallCallee),
		On(jdt.PropertyArguments).Roles(CallPositionalArgument),
	),
	On(jdt.MethodInvocation).Roles(Call, Expression).Children(
		On(jdt.PropertyExpression).Roles(CallReceiver),
		On(jdt.PropertyName).Roles(CallCallee),
		On(jdt.PropertyArguments).Roles(CallPositionalArgument),
	),
	On(jdt.SuperConstructorInvocation).Roles(Call, Statement, Incomplete).Children(
		On(jdt.PropertyExpression).Roles(CallReceiver),
		On(jdt.PropertyArguments).Roles(CallPositionalArgument),
	),
	On(jdt.SuperMethodInvocation).Roles(Call, Expression, Incomplete).Children(
		On(jdt.PropertyQualifier).Roles(CallCallee),
		On(jdt.PropertyName).Roles(CallCallee),
		On(jdt.PropertyArguments).Roles(CallPositionalArgument),
	),

	// Conditionals
	On(jdt.IfStatement).Roles(If, Statement).Children(
		On(jdt.PropertyExpression).Roles(IfCondition),
		On(jdt.PropertyThenStatement).Roles(IfBody),
		On(jdt.PropertyElseStatement).Roles(IfElse),
	),

	On(jdt.SwitchStatement).Roles(Switch, Statement).Children(
		//TODO: On(jdt.PropertyExpression).Roles(SwitchExpression),
		On(jdt.SwitchCase).Roles(Statement).Self(
			On(HasChild(Any)).Roles(SwitchCase).Children(
				On(jdt.PropertyExpression).Roles(SwitchCaseCondition),
			),
			On(Not(HasChild(Any))).Roles(SwitchDefault),
		),
		// FIXME: Switch case bodies are not enclosed in a block, thus it may
		// contain an arbitrary number of statements (of any kind). So this
		// is just an initial approach.
		On(jdt.ExpressionStatement).Roles(SwitchCaseBody),
	),

	// Loops
	On(jdt.EnhancedForStatement).Roles(ForEach, Statement).Children(
		On(jdt.PropertyParameter).Roles(ForInit, ForUpdate),
		On(jdt.PropertyExpression).Roles(ForExpression),
		On(jdt.PropertyBody).Roles(ForBody),
	),

	On(jdt.ForStatement).Roles(For, Statement).Children(
		On(jdt.PropertyInitializers).Roles(ForInit),
		On(jdt.PropertyExpression).Roles(ForExpression),
		On(jdt.PropertyUpdaters).Roles(ForUpdate),
		On(jdt.PropertyBody).Roles(ForBody),
	),

	On(jdt.WhileStatement).Roles(While, Statement).Children(
		On(jdt.PropertyExpression).Roles(WhileCondition),
		On(jdt.PropertyBody).Roles(WhileBody),
	),

	On(jdt.DoStatement).Roles(DoWhile, Statement).Children(
		On(jdt.PropertyExpression).Roles(DoWhileCondition),
		On(jdt.PropertyBody).Roles(DoWhileBody),
	),

	// Operators
	On(jdt.InfixExpression).Roles(BinaryExpression, BinaryExpressionOp, Expression).Self(
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
	).Children(
		On(jdt.PropertyLeftOperand).Roles(BinaryExpressionLeft),
		On(jdt.PropertyRightOperand).Roles(BinaryExpressionRight),
	),

	On(jdt.PostfixExpression).Roles(Expression).Self(
		On(HasProperty("operator", "++")).Roles(OpPostIncrement),
		On(HasProperty("operator", "--")).Roles(OpPostDecrement),
	),

	On(jdt.PrefixExpression).Roles(Expression).Self(
		On(HasProperty("operator", "++")).Roles(OpPreIncrement),
		On(HasProperty("operator", "--")).Roles(OpPreDecrement),
		On(HasProperty("operator", "+")).Roles(OpPositive),
		On(HasProperty("operator", "-")).Roles(OpNegative),
		On(HasProperty("operator", "~")).Roles(OpBitwiseComplement),
		On(HasProperty("operator", "!")).Roles(OpBooleanNot),
	),

	On(jdt.Assignment).Roles(Assignment, Expression).Children(
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

	// Types
	On(jdt.ArrayType).Roles(Incomplete),
	On(jdt.IntersectionType).Roles(Incomplete),
	On(jdt.NameQualifiedType).Roles(Incomplete),
	On(jdt.ParameterizedType).Roles(Incomplete),
	On(jdt.PrimitiveType).Roles(Incomplete),
	On(jdt.QualifiedType).Roles(Incomplete),
	On(jdt.SimpleType).Roles(Incomplete),
	On(jdt.UnionType).Roles(Incomplete),
	On(jdt.WildcardType).Roles(Incomplete),

	// Modifiers
	On(jdt.Modifier).Self(
		On(jdt.KeywordPublic).Roles(VisibleFromWorld),
		On(jdt.KeywordProtected).Roles(VisibleFromSubtype),
		On(jdt.KeywordPrivate).Roles(VisibleFromInstance),

		// class | method | interface
		On(jdt.KeywordAbstract).Roles(Incomplete),
		// class | field | method | interface
		On(jdt.KeywordStatic).Roles(Incomplete),
		// class | field | method
		On(jdt.KeywordFinal).Roles(Incomplete),
		// class | method | interface
		On(jdt.KeywordStrictfp).Roles(Incomplete),
		// field
		On(jdt.KeywordTransient).Roles(Incomplete),
		On(jdt.KeywordVolatile).Roles(Incomplete),
		// method
		On(jdt.KeywordSynchronized).Roles(Incomplete),
		On(jdt.KeywordNative).Roles(Incomplete),
	),

	// Exceptions
	On(jdt.TryStatement).Roles(Try, Statement).Children(
		// TODO: TryWithResourcesStatement
		On(jdt.PropertyBody).Roles(TryBody),
		On(jdt.PropertyCatchClauses).Roles(TryCatch),
		On(jdt.PropertyFinally).Roles(TryFinally),
	),

	On(jdt.ThrowStatement).Roles(Throw, Statement),

	On(jdt.AssertStatement).Roles(Assert, Statement),

	// Annotations
	On(jdt.MarkerAnnotation).Roles(Incomplete),
	On(jdt.MemberRef).Roles(Incomplete),
	On(jdt.MemberValuePair).Roles(Incomplete),
	On(jdt.MethodRef).Roles(Incomplete),
	On(jdt.MethodRefParameter).Roles(Incomplete),
	On(jdt.NormalAnnotation).Roles(Incomplete),
	On(jdt.SingleMemberAnnotation).Roles(Incomplete),
	On(jdt.TagElement).Roles(Incomplete),
	On(jdt.TextElement).Roles(Incomplete),

	// Comments
	On(jdt.BlockComment).Roles(Comment),
	On(jdt.Javadoc).Roles(Documentation, Comment),
	On(jdt.LineComment).Roles(Comment),

	// Other expressions
	On(jdt.ArrayAccess).Roles(Expression, Incomplete),
	On(jdt.ArrayCreation).Roles(Expression, Incomplete),
	On(jdt.CastExpression).Roles(Expression, Incomplete),
	On(jdt.CreationReference).Roles(Expression, Incomplete),
	On(jdt.ExpressionMethodReference).Roles(Expression, Incomplete),
	On(jdt.ParenthesizedExpression).Roles(Expression, Incomplete),
	On(jdt.SuperMethodReference).Roles(Expression, Incomplete),
	On(jdt.ThisExpression).Roles(This, Expression),

	// Other statements
	On(jdt.Block).Roles(BlockScope, Block, Statement),
	On(jdt.BreakStatement).Roles(Break, Statement),
	On(jdt.EmptyStatement).Roles(Statement),
	On(jdt.ExpressionStatement).Roles(Statement),
	On(jdt.LabeledStatement).Roles(Statement, Incomplete),
	On(jdt.ReturnStatement).Roles(Return, Statement),
	On(jdt.SynchronizedStatement).Roles(Statement, Incomplete),

	// Others
	On(jdt.ArrayInitializer).Roles(Incomplete),
	On(jdt.Dimension).Roles(Incomplete),
	On(jdt.TypeParameter).Roles(Incomplete),
)
