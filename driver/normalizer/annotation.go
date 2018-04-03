package normalizer


import (
	"gopkg.in/bblfsh/sdk.v1/uast/role"
	. "gopkg.in/bblfsh/sdk.v1/uast/transformer"
)

// Native is the of list `transformer.Transformer` to apply to a native AST.
// To learn more about the Transformers and the available ones take a look to:
// https://godoc.org/gopkg.in/bblfsh/sdk.v1/uast/transformer
var Native = Transformers([][]Transformer{
	{
		// ResponseMetadata is a transform that trims response metadata from AST.
		//
		// https://godoc.org/gopkg.in/bblfsh/sdk.v1/uast#ResponseMetadata
		ResponseMetadata{
			TopLevelIsRootNode: false,
		},
	},
	// The main block of transformation rules.
	{Mappings(Annotations...)},
	{
		// RolesDedup is used to remove duplicate roles assigned by multiple
		// transformation rules.
		RolesDedup(),
	},
}...)

// Code is a special block of transformations that are applied at the end
// and can access original source code file. It can be used to improve or
// fix positional information.
//
// https://godoc.org/gopkg.in/bblfsh/sdk.v1/uast/transformer/positioner
var Code []CodeTransformer // Java already provides all the information we need

// mapAST is a helper for describing a single AST transformation for a given node type.
func mapAST(typ string, ast, norm ObjectOp, roles ...role.Role) Mapping {
	return ASTMap(typ,
		ASTObjectLeft(typ, ast),
		ASTObjectRight(typ, norm, nil, roles...),
	)
}

// Annotations is a list of individual transformations to annotate a native AST with roles.
var Annotations = []Mapping{
}

/*
var AnnotationRules = On(jdt.CompilationUnit).Roles(uast.File).Descendants(
	// Names
	On(jdt.QualifiedName).Roles(uast.Expression, uast.Identifier, uast.Qualified),
	On(jdt.SimpleName).Roles(uast.Expression, uast.Identifier),

	// Visibility
	On(Or(jdt.MethodDeclaration, jdt.TypeDeclaration)).Self(
		On(HasChild(And(jdt.Modifier, jdt.KeywordPublic))).Roles(uast.Visibility, uast.World),
		On(HasChild(And(jdt.Modifier, jdt.KeywordPrivate))).Roles(uast.Visibility, uast.Type),
		On(HasChild(And(jdt.Modifier, jdt.KeywordProtected))).Roles(uast.Visibility, uast.Subtype),
		On(Not(HasChild(And(jdt.Modifier,
			Or(jdt.KeywordPublic, jdt.KeywordPrivate, jdt.KeywordProtected),
		)))).Roles(uast.Visibility, uast.Package),
	),

	// Package and imports
	On(jdt.PackageDeclaration).Roles(uast.Declaration, uast.Package),
	On(jdt.ImportDeclaration).Roles(uast.Declaration, uast.Import).Children(
		On(jdt.QualifiedName).Roles(uast.Pathname, uast.Import),
	),

	// Type declarations
	On(jdt.AnonymousClassDeclaration).Roles(uast.Expression, uast.Declaration, uast.Type, uast.Anonymous).Children(
		On(jdt.PropertyBodyDeclarations).Roles(uast.Body),
	),
	On(jdt.AnnotationTypeDeclaration).Roles(uast.Declaration, uast.Type, uast.Annotation).Children(
		On(jdt.PropertyBodyDeclarations).Roles(uast.Body),
	),
	On(jdt.EnumDeclaration).Roles(uast.Declaration, uast.Type, uast.Enumeration),

	// ClassDeclaration | InterfaceDeclaration
	On(jdt.TypeDeclaration).Roles(uast.Declaration, uast.Type),
	// Local (TypeDeclaration | EnumDeclaration)
	On(jdt.TypeDeclarationStatement).Roles(uast.Statement, uast.Declaration, uast.Type),

	// Method declarations
	On(jdt.MethodDeclaration).Roles(uast.Declaration, uast.Function).Children(
		On(jdt.PropertyName).Roles(uast.Function, uast.Name),
		On(jdt.PropertyBody).Roles(uast.Function, uast.Body),
		On(jdt.PropertyParameters).Roles(uast.Function, uast.Argument).Self(
			On(HasProperty("varargs", "true")).Roles(uast.Function, uast.ArgsList),
		).Children(
			On(jdt.PropertyName).Roles(uast.Function, uast.Name, uast.Argument),
		),
	),
	On(jdt.LambdaExpression).Roles(uast.Declaration, uast.Function, uast.Anonymous).Children(
		On(jdt.PropertyBody).Roles(uast.Function, uast.Body),
		On(jdt.PropertyParameters).Roles(uast.Function, uast.Argument).Self(
			On(HasProperty("varargs", "true")).Roles(uast.Function, uast.ArgsList),
		).Children(
			On(jdt.PropertyName).Roles(uast.Function, uast.Name, uast.Argument),
		),
	),
	On(jdt.TypeMethodReference).Roles(uast.Declaration, uast.Function).Children(
		On(jdt.PropertyName).Roles(uast.Function, uast.Name),
		On(jdt.PropertyTypeArguments).Roles(uast.Function, uast.Argument),
		On(jdt.PropertyType).Roles(uast.Function, uast.Return),
	),

	// Other declarations
	On(jdt.AnnotationTypeMemberDeclaration).Roles(uast.Declaration, uast.Type, uast.Annotation),
	On(jdt.EnumConstantDeclaration).Roles(uast.Declaration, uast.Enumeration),
	On(jdt.FieldDeclaration).Roles(uast.Declaration, uast.Variable),
	// TODO: differentiate between static (class) and instance initialization
	On(jdt.Initializer).Roles(uast.Initialization, uast.Block, uast.Incomplete),
	On(jdt.SingleVariableDeclaration).Roles(uast.Declaration, uast.Variable),
	On(jdt.VariableDeclarationExpression).Roles(uast.Expression, uast.Declaration, uast.Variable),
	On(jdt.VariableDeclarationFragment).Roles(uast.Declaration, uast.Variable),
	On(jdt.VariableDeclarationStatement).Roles(uast.Statement, uast.Declaration, uast.Variable),

	// Literals
	On(jdt.BooleanLiteral).Roles(uast.Expression, uast.Literal, uast.Boolean),
	On(jdt.CharacterLiteral).Roles(uast.Expression, uast.Literal, uast.Character),
	On(jdt.NullLiteral).Roles(uast.Expression, uast.Literal, uast.Null),
	On(jdt.NumberLiteral).Roles(uast.Expression, uast.Literal, uast.Number),
	On(jdt.StringLiteral).Roles(uast.Expression, uast.Literal, uast.String),
	On(jdt.TypeLiteral).Roles(uast.Expression, uast.Literal, uast.Type),

	// Calls
	On(jdt.ClassInstanceCreation).Roles(uast.Expression, uast.Call, uast.Instance).Children(
		On(jdt.PropertyType).Roles(uast.Call, uast.Callee),
		On(jdt.PropertyArguments).Roles(uast.Call, uast.Argument, uast.Positional),
	),
	On(jdt.ConstructorInvocation).Roles(uast.Statement, uast.Call, uast.Incomplete).Children(
		On(jdt.PropertyType).Roles(uast.Call, uast.Callee),
		On(jdt.PropertyArguments).Roles(uast.Call, uast.Argument, uast.Positional),
	),
	On(jdt.MethodInvocation).Roles(uast.Expression, uast.Call).Children(
		On(jdt.PropertyExpression).Roles(uast.Call, uast.Receiver),
		On(jdt.PropertyName).Roles(uast.Call, uast.Callee),
		On(jdt.PropertyArguments).Roles(uast.Call, uast.Argument, uast.Positional),
	),
	On(jdt.SuperConstructorInvocation).Roles(uast.Statement, uast.Call, uast.Base, uast.Incomplete).Children(
		On(jdt.PropertyExpression).Roles(uast.Call, uast.Receiver),
		On(jdt.PropertyArguments).Roles(uast.Call, uast.Argument, uast.Positional),
	),
	On(jdt.SuperMethodInvocation).Roles(uast.Expression, uast.Call, uast.Base).Children(
		On(jdt.PropertyQualifier).Roles(uast.Call, uast.Callee),
		On(jdt.PropertyName).Roles(uast.Call, uast.Callee),
		On(jdt.PropertyArguments).Roles(uast.Call, uast.Argument, uast.Positional),
	),

	// Conditionals
	On(jdt.IfStatement).Roles(uast.Statement, uast.If).Children(
		On(jdt.PropertyExpression).Roles(uast.If, uast.Condition),
		On(jdt.PropertyThenStatement).Roles(uast.If, uast.Then, uast.Body),
		On(jdt.PropertyElseStatement).Roles(uast.If, uast.Else, uast.Body),
	),
	On(jdt.ConditionalExpression).Roles(uast.Expression, uast.If).Children(
		On(jdt.PropertyExpression).Roles(uast.If, uast.Condition),
		On(jdt.PropertyThenExpression).Roles(uast.If, uast.Then),
		On(jdt.PropertyElseExpression).Roles(uast.If, uast.Else),
	),

	On(jdt.SwitchStatement).Roles(uast.Statement, uast.Switch).Children(
		On(jdt.PropertyExpression).Roles(uast.Expression, uast.Switch),
		On(jdt.SwitchCase).Roles(uast.Statement, uast.Switch).Self(
			On(HasChild(Any)).Roles(uast.Case).Children(
				On(jdt.PropertyExpression).Roles(uast.Expression, uast.Switch, uast.Case, uast.Condition),
			),
			On(Not(HasChild(Any))).Roles(uast.Default),
		),
		// FIXME: Switch case bodies are not enclosed in a block, thus it may
		// contain an arbitrary number of statements (of any kind). So this
		// is just an initial approach.
		On(jdt.ExpressionStatement).Roles(uast.Switch, uast.Case, uast.Body),
	),

	// Loops
	On(jdt.EnhancedForStatement).Roles(uast.Statement, uast.For, uast.Iterator).Children(
		On(jdt.PropertyParameter).Roles(uast.For, uast.Iterator),
		On(jdt.PropertyExpression).Roles(uast.Expression, uast.For),
		On(jdt.PropertyBody).Roles(uast.For, uast.Body),
	),

	On(jdt.ForStatement).Roles(uast.Statement, uast.For).Children(
		On(jdt.PropertyInitializers).Roles(uast.For, uast.Initialization),
		On(jdt.PropertyExpression).Roles(uast.Expression, uast.For, uast.Condition),
		On(jdt.PropertyUpdaters).Roles(uast.For, uast.Update),
		On(jdt.PropertyBody).Roles(uast.For, uast.Body),
	),

	On(jdt.WhileStatement).Roles(uast.Statement, uast.While).Children(
		On(jdt.PropertyExpression).Roles(uast.Expression, uast.While, uast.Condition),
		On(jdt.PropertyBody).Roles(uast.While, uast.Body),
	),

	On(jdt.DoStatement).Roles(uast.Statement, uast.DoWhile).Children(
		On(jdt.PropertyExpression).Roles(uast.DoWhile, uast.Condition),
		On(jdt.PropertyBody).Roles(uast.DoWhile, uast.Body),
	),

	// Operators
	On(jdt.InfixExpression).Roles(uast.Expression, uast.Binary, uast.Operator).Self(
		On(HasProperty("operator", "+")).Roles(uast.Arithmetic, uast.Add),
		On(HasProperty("operator", "-")).Roles(uast.Arithmetic, uast.Substract),
		On(HasProperty("operator", "*")).Roles(uast.Arithmetic, uast.Multiply),
		On(HasProperty("operator", "/")).Roles(uast.Arithmetic, uast.Divide),
		On(HasProperty("operator", "%")).Roles(uast.Arithmetic, uast.Modulo),
		On(HasProperty("operator", "<<")).Roles(uast.Bitwise, uast.LeftShift),
		On(HasProperty("operator", ">>")).Roles(uast.Bitwise, uast.RightShift),
		On(HasProperty("operator", ">>>")).Roles(uast.Bitwise, uast.RightShift, uast.Unsigned),
		On(HasProperty("operator", "<")).Roles(uast.LessThan, uast.Relational),
		On(HasProperty("operator", ">")).Roles(uast.GreaterThan, uast.Relational),
		On(HasProperty("operator", "<=")).Roles(uast.LessThanOrEqual, uast.Relational),
		On(HasProperty("operator", ">=")).Roles(uast.GreaterThanOrEqual, uast.Relational),
		On(HasProperty("operator", "==")).Roles(uast.Equal, uast.Relational),
		On(HasProperty("operator", "!=")).Roles(uast.Equal, uast.Not, uast.Relational),
		On(HasProperty("operator", "&")).Roles(uast.Bitwise, uast.And),
		On(HasProperty("operator", "|")).Roles(uast.Bitwise, uast.Or),
		On(HasProperty("operator", "&&")).Roles(uast.Boolean, uast.And),
		On(HasProperty("operator", "||")).Roles(uast.Boolean, uast.Or),
		On(HasProperty("operator", "^")).Roles(uast.Boolean, uast.Xor),
	).Children(
		On(jdt.PropertyLeftOperand).Roles(uast.Expression, uast.Binary, uast.Left),
		On(jdt.PropertyRightOperand).Roles(uast.Expression, uast.Binary, uast.Right),
	),

	On(jdt.PostfixExpression).Roles(uast.Expression, uast.Operator, uast.Unary, uast.Postfix).Self(
		On(HasProperty("operator", "++")).Roles(uast.Arithmetic, uast.Increment),
		On(HasProperty("operator", "--")).Roles(uast.Arithmetic, uast.Increment),
	),

	On(jdt.PrefixExpression).Roles(uast.Expression, uast.Operator, uast.Unary).Self(
		On(HasProperty("operator", "++")).Roles(uast.Arithmetic, uast.Increment),
		On(HasProperty("operator", "--")).Roles(uast.Arithmetic, uast.Decrement),
		On(HasProperty("operator", "+")).Roles(uast.Arithmetic, uast.Positive),
		On(HasProperty("operator", "-")).Roles(uast.Arithmetic, uast.Negative),
		On(HasProperty("operator", "~")).Roles(uast.Bitwise, uast.Not),
		On(HasProperty("operator", "!")).Roles(uast.Boolean, uast.Not),
	),

	On(jdt.Assignment).Roles(uast.Expression, uast.Assignment, uast.Operator, uast.Binary).Children(
		On(jdt.PropertyLeftHandSide).Roles(uast.Assignment, uast.Binary, uast.Left),
		On(jdt.PropertyRightHandSide).Roles(uast.Assignment, uast.Binary, uast.Right),
	).Self(
		On(HasProperty("operator", "+=")).Roles(uast.Arithmetic, uast.Add),
		On(HasProperty("operator", "-=")).Roles(uast.Arithmetic, uast.Substract),
		On(HasProperty("operator", "*=")).Roles(uast.Arithmetic, uast.Multiply),
		On(HasProperty("operator", "/=")).Roles(uast.Arithmetic, uast.Divide),
		On(HasProperty("operator", "%=")).Roles(uast.Arithmetic, uast.Modulo),
		On(HasProperty("operator", "&=")).Roles(uast.Bitwise, uast.And),
		On(HasProperty("operator", "|=")).Roles(uast.Bitwise, uast.Or),
		On(HasProperty("operator", "^=")).Roles(uast.Bitwise, uast.Xor),
		On(HasProperty("operator", "<<=")).Roles(uast.Bitwise, uast.LeftShift),
		On(HasProperty("operator", ">>=")).Roles(uast.Bitwise, uast.RightShift),
		On(HasProperty("operator", ">>>=")).Roles(uast.Bitwise, uast.RightShift, uast.Unsigned),
	),

	// Types
	On(jdt.ArrayType).Roles(uast.Type, uast.Primitive, uast.List),
	On(jdt.IntersectionType).Roles(uast.Type, uast.And),
	On(jdt.NameQualifiedType).Roles(uast.Type, uast.Name, uast.Qualified),
	On(jdt.ParameterizedType).Roles(uast.Type, uast.Incomplete),
	On(jdt.PrimitiveType).Roles(uast.Type, uast.Primitive),
	On(jdt.QualifiedType).Roles(uast.Type, uast.Qualified),
	On(jdt.SimpleType).Roles(uast.Type),
	On(jdt.UnionType).Roles(uast.Type, uast.Or),
	On(jdt.WildcardType).Roles(uast.Type, uast.Incomplete),

	// Modifiers
	On(jdt.Modifier).Self(
		On(jdt.KeywordPublic).Roles(uast.Visibility, uast.World),
		On(jdt.KeywordProtected).Roles(uast.Visibility, uast.Subtype),
		On(jdt.KeywordPrivate).Roles(uast.Visibility, uast.Instance),

		// class | method | interface
		On(jdt.KeywordAbstract).Roles(uast.Incomplete),
		// class | field | method | interface
		On(jdt.KeywordStatic).Roles(uast.Incomplete),
		// class | field | method
		On(jdt.KeywordFinal).Roles(uast.Incomplete),
		// class | method | interface
		On(jdt.KeywordStrictfp).Roles(uast.Incomplete),
		// field
		On(jdt.KeywordTransient).Roles(uast.Incomplete),
		On(jdt.KeywordVolatile).Roles(uast.Incomplete),
		// method
		On(jdt.KeywordSynchronized).Roles(uast.Incomplete),
		On(jdt.KeywordNative).Roles(uast.Incomplete),
	),

	// Exceptions
	On(jdt.TryStatement).Roles(uast.Statement, uast.Try).Children(
		On(jdt.PropertyResources).Roles(uast.Try),
		On(jdt.PropertyBody).Roles(uast.Try, uast.Body),
		On(jdt.PropertyCatchClauses).Roles(uast.Try, uast.Catch),
		On(jdt.PropertyFinally).Roles(uast.Try, uast.Finally),
	),

	On(jdt.ThrowStatement).Roles(uast.Statement, uast.Throw),

	On(jdt.AssertStatement).Roles(uast.Statement, uast.Assert),

	// Annotations
	On(jdt.MarkerAnnotation).Roles(uast.Annotation, uast.Incomplete),
	On(jdt.NormalAnnotation).Roles(uast.Annotation, uast.Incomplete),
	On(jdt.SingleMemberAnnotation).Roles(uast.Annotation, uast.Incomplete),
	On(jdt.MemberValuePair).Roles(uast.Annotation, uast.Incomplete),

	// Comments
	On(jdt.BlockComment).Roles(uast.Comment),
	On(jdt.Javadoc).Roles(uast.Documentation, uast.Comment),
	On(jdt.LineComment).Roles(uast.Comment),

	// Javadoc tags
	On(jdt.MemberRef).Roles(uast.Documentation, uast.Variable, uast.Incomplete),
	On(jdt.MethodRef).Roles(uast.Documentation, uast.Function, uast.Incomplete),
	On(jdt.MethodRefParameter).Roles(uast.Documentation, uast.Function, uast.Incomplete),
	On(jdt.TagElement).Roles(uast.Documentation, uast.Incomplete),
	On(jdt.TextElement).Roles(uast.Documentation, uast.Incomplete),

	// Other expressions
	On(jdt.ArrayAccess).Roles(uast.Expression, uast.Incomplete),
	On(jdt.ArrayCreation).Roles(uast.Expression, uast.Incomplete),
	On(jdt.CastExpression).Roles(uast.Expression, uast.Incomplete),
	On(jdt.CreationReference).Roles(uast.Expression, uast.Incomplete),
	On(jdt.ExpressionMethodReference).Roles(uast.Expression, uast.Incomplete),
	On(jdt.ParenthesizedExpression).Roles(uast.Expression, uast.Incomplete),
	On(jdt.SuperMethodReference).Roles(uast.Expression, uast.Incomplete),
	On(jdt.ThisExpression).Roles(uast.Expression, uast.This),

	// Other statements
	On(jdt.Block).Roles(uast.Statement, uast.Block, uast.Scope),
	On(jdt.BreakStatement).Roles(uast.Statement, uast.Break),
	On(jdt.EmptyStatement).Roles(uast.Statement),
	On(jdt.ExpressionStatement).Roles(uast.Statement),
	On(jdt.LabeledStatement).Roles(uast.Statement, uast.Incomplete),
	On(jdt.ReturnStatement).Roles(uast.Statement, uast.Return),
	On(jdt.SynchronizedStatement).Roles(uast.Statement, uast.Incomplete),

	// Others
	On(jdt.ArrayInitializer).Roles(uast.Expression, uast.List, uast.Literal),
	On(jdt.Dimension).Roles(uast.Type, uast.Incomplete),
	On(jdt.TypeParameter).Roles(uast.Type, uast.Incomplete),
)
*/