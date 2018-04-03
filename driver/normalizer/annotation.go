package normalizer

import (
	"gopkg.in/bblfsh/sdk.v1/uast/role"
	. "gopkg.in/bblfsh/sdk.v1/uast/transformer"
	"gopkg.in/bblfsh/sdk.v1/uast"
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
	return mapASTCustom(typ, ast, norm, nil, roles...)
}

func mapASTCustom(typ string, ast, norm ObjectOp, rop ArrayOp,  roles ...role.Role) Mapping {
	return ASTMap(typ,
		ASTObjectLeft(typ, ast),
		ASTObjectRight(typ, norm, rop, roles...),
	)
}

var (
	modifierRoles = make(map[uast.Value]ArrayOp)
	infixRoles = make(map[uast.Value]ArrayOp)
	postfixRoles = make(map[uast.Value]ArrayOp)
	prefixRoles = make(map[uast.Value]ArrayOp)
	assignRoles = make(map[uast.Value]ArrayOp)
	variadicRoles = make(map[uast.Value]ArrayOp)
)

func fillTokenToRolesMap(dst map[uast.Value]ArrayOp, src map[string][]role.Role) {
	for tok, roles := range src {
		dst[uast.String(tok)] = Roles(roles...)
	}
}

func init() {
	fillTokenToRolesMap(modifierRoles, map[string][]role.Role{
		"public": {role.Visibility, role.World},
		"protected": {role.Visibility, role.Subtype},
		"private": {role.Visibility, role.Instance},

		// class | method | interface
		"abstract": {role.Incomplete},
			// class | field | method | interface
			"static": {role.Incomplete},
			// class | field | method
			"final": {role.Incomplete},
			// class | method | interface
			"strictfp": {role.Incomplete},
			// field
			"transient": {role.Incomplete},
			"volatile": {role.Incomplete},
			// method
			"synchronized": {role.Incomplete},
			"native": {role.Incomplete},
	})
	fillTokenToRolesMap(infixRoles, map[string][]role.Role{
		"+": {role.Arithmetic, role.Add},
		"-": {role.Arithmetic, role.Substract},
		"*": {role.Arithmetic, role.Multiply},
		"/": {role.Arithmetic, role.Divide},
		"%": {role.Arithmetic, role.Modulo},
		"<<": {role.Bitwise, role.LeftShift},
		">>": {role.Bitwise, role.RightShift},
		">>>": {role.Bitwise, role.RightShift, role.Unsigned},
		"<": {role.LessThan, role.Relational},
		">": {role.GreaterThan, role.Relational},
		"<=": {role.LessThanOrEqual, role.Relational},
		">=": {role.GreaterThanOrEqual, role.Relational},
		"==": {role.Equal, role.Relational},
		"!=": {role.Equal, role.Not, role.Relational},
		"&": {role.Bitwise, role.And},
		"|": {role.Bitwise, role.Or},
		"&&": {role.Boolean, role.And},
		"||": {role.Boolean, role.Or},
		"^": {role.Boolean, role.Xor},
	})
	fillTokenToRolesMap(postfixRoles, map[string][]role.Role{
		"++": {role.Arithmetic, role.Increment},
		"--": {role.Arithmetic, role.Decrement},
	})
	fillTokenToRolesMap(prefixRoles, map[string][]role.Role{
		"++": {role.Arithmetic, role.Increment},
		"--": {role.Arithmetic, role.Decrement},
		"+": {role.Arithmetic, role.Positive},
		"-": {role.Arithmetic, role.Negative},
		"~": {role.Bitwise, role.Not},
		"!": {role.Boolean, role.Not},
	})
	fillTokenToRolesMap(assignRoles, map[string][]role.Role{
		"=": {},
		"+=": {role.Arithmetic, role.Add},
		"-=": {role.Arithmetic, role.Substract},
		"*=": {role.Arithmetic, role.Multiply},
		"/=": {role.Arithmetic, role.Divide},
		"%=": {role.Arithmetic, role.Modulo},
		"&=": {role.Bitwise, role.And},
		"|=": {role.Bitwise, role.Or},
		"^=": {role.Bitwise, role.Xor},
		"<<=": {role.Bitwise, role.LeftShift},
		">>=": {role.Bitwise, role.RightShift},
		">>>=": {role.Bitwise, role.RightShift, role.Unsigned},
	})
	fillTokenToRolesMap(variadicRoles, map[string][]role.Role{
		"true": {role.ArgsList},
		"false": {},
	})
}


type objRoles map[string][]role.Role

func annotateType(typ string, fields objRoles, roles ...role.Role) Mapping {
	left := make(Obj, len(fields))
	right := make(Obj, len(fields))
	for name, roles := range fields {
		left[name] = OptObjectRoles(name + "_var")
		right[name] = OptObjectRoles(name+"_var", roles...)
	}
	return mapAST(typ, left, right, roles...)
}

func annotateModifiers(typ string, mod string, roles ...role.Role) Mapping {
	c := Any(Has{
		uast.KeyToken: String(mod),
	})
	if mod == "" {
		c = All(Not(Has{
			uast.KeyToken: In(
				uast.String("public"),
				uast.String("private"),
				uast.String("protected"),
			),
		}))
	}
	return mapAST(typ, Obj{
		"modifiers": Check(c,Var("mod")),
	}, Obj{
		"modifiers": Var("mod"),
	}, roles...)
}

// Annotations is a list of individual transformations to annotate a native AST with roles.
var Annotations = []Mapping{
	annotateType("CompilationUnit", nil, role.File),

	// Names
	annotateType("QualifiedName", nil, role.Expression, role.Identifier, role.Qualified),
	mapAST("SimpleName", Obj{
		"identifier": Var("name"),
	}, Obj{
		uast.KeyToken: Var("name"),
	}, role.Expression, role.Identifier),
	
	mapAST("PrimitiveType", Obj{
		"primitiveTypeCode": Var("type"),
	}, Obj{
		uast.KeyToken: Var("type"),
	}, role.Type, role.Primitive),

	// Visibility
	mapAST("TypeDeclaration", Obj{
		"modifiers": Is(nil),
	}, Obj{
		"modifiers": Is(nil),
	}, role.Visibility, role.Package),

	mapAST("MethodDeclaration", Obj{
		"modifiers": Is(nil),
	}, Obj{
		"modifiers": Is(nil),
	}, role.Visibility, role.Package),

	annotateModifiers("TypeDeclaration", "", role.Visibility, role.Package),
	annotateModifiers("TypeDeclaration", "public", role.Visibility, role.World),
	annotateModifiers("TypeDeclaration", "private", role.Visibility, role.Type),
	annotateModifiers("TypeDeclaration", "protected", role.Visibility, role.Subtype),

	annotateModifiers("MethodDeclaration", "", role.Visibility, role.Package),
	annotateModifiers("MethodDeclaration", "public", role.Visibility, role.World),
	annotateModifiers("MethodDeclaration", "private", role.Visibility, role.Type),
	annotateModifiers("MethodDeclaration", "protected", role.Visibility, role.Subtype),

	// Package and imports
	annotateType("ImportDeclaration", nil, role.Declaration, role.Import),
	mapAST("ImportDeclaration", Obj{
		"name": ObjectRolesCustom("q", Obj{
			uast.KeyType: String("QualifiedName"),
		}),
	}, Obj{
		"name": ObjectRolesCustom("q", Obj{
			uast.KeyType: String("QualifiedName"),
		}, role.Pathname, role.Import),
	}),
	mapAST("ImportDeclaration", Obj{
		"qualifier": ObjectRolesCustom("q", Obj{
			uast.KeyType: String("QualifiedName"),
		}),
	}, Obj{
		"qualifier": ObjectRolesCustom("q", Obj{
			uast.KeyType: String("QualifiedName"),
		}, role.Pathname, role.Import),
	}),

	mapAST("PackageDeclaration", Obj{}, Obj{
		uast.KeyToken: String("package"),
	}, role.Declaration, role.Package),

	// Type declarations
	mapAST("AnonymousClassDeclaration", Obj{
		"bodyDeclarations": Each("decls", ObjectRoles("decl")),
	}, Obj{
		"bodyDeclarations": Each("decls", ObjectRoles("decl", role.Body)),
	},role.Expression, role.Declaration, role.Type, role.Anonymous),

	mapAST("AnnotationTypeDeclaration", Obj{
		"bodyDeclarations": Each("decls", ObjectRoles("decl")),
	}, Obj{
		"bodyDeclarations": Each("decls", ObjectRoles("decl", role.Body)),
	},role.Declaration, role.Type, role.Annotation),
	annotateType("EnumDeclaration", nil, role.Declaration, role.Type, role.Enumeration),

	// ClassDeclaration | InterfaceDeclaration
	annotateType("TypeDeclaration", nil, role.Declaration, role.Type),
	// Local (TypeDeclaration | EnumDeclaration)
	annotateType("TypeDeclarationStatement", nil, role.Statement, role.Declaration, role.Type),

	// Method declarations
	mapAST("MethodDeclaration", Fields{
		{Name: "name", Op: ObjectRoles("name")},
		{Name: "body", Op: OptObjectRoles("body")},
		{Name: "parameters", Op: Each("param", ObjectRolesCustom("p", Obj{
			"name": ObjectRoles("pname"),
			"varargs": Var("variadic"),
		}))},
	}, Fields{ // ->
		{Name: "name", Op: ObjectRoles("name", role.Function, role.Name)},
		{Name: "body", Op: OptObjectRoles("body", role.Function, role.Body)},
		{Name: "parameters", Op: Each("param", ObjectRolesCustomOp("p", Obj{
			"name": ObjectRoles("pname",role.Function, role.Name, role.Argument),
			"varargs": Var("variadic"),
		}, LookupArrOpVar("variadic", variadicRoles), role.Function, role.Argument))},
	}, role.Declaration, role.Function),

	mapAST("LambdaExpression", Fields{
		{Name: "body", Op: ObjectRoles("body")},
		{Name: "parameters", Op: Each("param", ObjectRolesCustom("p", Obj{
			"name": ObjectRoles("pname"),
		}))},
	}, Fields{ // ->
		{Name: "body", Op: ObjectRoles("body", role.Function, role.Body)},
		{Name: "parameters", Op: Each("param", ObjectRolesCustom("p", Obj{
			"name": ObjectRoles("pname",role.Function, role.Name, role.Argument ),
		}, role.Function, role.Argument))},
	}, role.Declaration, role.Function, role.Anonymous),

	mapAST("TypeMethodReference", Obj{
		"name": ObjectRoles("name"),
		"type": ObjectRoles("type"),
		"typeArguments": Each("targs", ObjectRoles("targ")),
	}, Obj{
		"name": ObjectRoles("name", role.Function, role.Name),
		"type": ObjectRoles("type", role.Function, role.Return),
		"typeArguments": Each("targs", ObjectRoles("targ", role.Function, role.Argument)),
	}, role.Declaration, role.Function),

	// Other declarations
	annotateType("AnnotationTypeMemberDeclaration", nil, role.Declaration, role.Type, role.Annotation),
	annotateType("EnumConstantDeclaration", nil, role.Declaration, role.Enumeration),
	annotateType("FieldDeclaration", nil, role.Declaration, role.Variable),
	// TODO: differentiate between static (class) and instance initialization
	annotateType("Initializer", nil, role.Initialization, role.Block, role.Incomplete),
	annotateType("SingleVariableDeclaration", nil, role.Declaration, role.Variable),
	annotateType("VariableDeclarationExpression", nil, role.Expression, role.Declaration, role.Variable),
	annotateType("VariableDeclarationFragment", nil, role.Declaration, role.Variable),
	annotateType("VariableDeclarationStatement", nil, role.Statement, role.Declaration, role.Variable),

	// Literals
	annotateType("BooleanLiteral", nil, role.Expression, role.Literal, role.Boolean),
	annotateType("TypeLiteral", nil, role.Expression, role.Literal, role.Type),
	annotateType("NumberLiteral", nil, role.Expression, role.Literal, role.Number),

	mapAST("NullLiteral", Obj{}, Obj{
		uast.KeyToken: String("null"),
	}, role.Expression, role.Literal, role.Null),


	mapAST("StringLiteral", Obj{
		"escapedValue": Var("v"),
	}, Obj{
		uast.KeyToken: Var("v"),
	}, role.Expression, role.Literal, role.String),


	mapAST("CharacterLiteral", Obj{
		"escapedValue": Var("v"),
	}, Obj{
		uast.KeyToken: Var("v"),
	}, role.Expression, role.Literal, role.Character),

	// Calls
	mapAST("ClassInstanceCreation", Obj{
		"type": ObjectRoles("type"),
		"arguments": Each("args", ObjectRoles("arg")),
	},Obj{
		"type": ObjectRoles("type", role.Call, role.Callee),
		"arguments": Each("args", ObjectRoles("arg", role.Call, role.Argument, role.Positional)),
	}, role.Expression, role.Call, role.Instance),

	mapAST("ConstructorInvocation", Obj{
		//"type": ObjectRoles("type"),
		"arguments": Each("args", ObjectRoles("arg")),
	},Obj{
		//"type": ObjectRoles("type", role.Call, role.Callee),
		"arguments": Each("args", ObjectRoles("arg", role.Call, role.Argument, role.Positional)),
	}, role.Statement, role.Call, role.Incomplete),

	mapAST("MethodInvocation", Obj{
		"expression":  OptObjectRoles("expr"),
		"name": OptObjectRoles("name"),
		"arguments": Each("args", ObjectRoles("arg")),
	}, Obj{
		"expression":  OptObjectRoles("expr", role.Call, role.Receiver),
		"name": OptObjectRoles("name", role.Call, role.Callee),
		"arguments": Each("args", ObjectRoles("arg", role.Call, role.Argument, role.Positional)),
	}, role.Expression, role.Call),

	annotateType("SuperConstructorInvocation", objRoles{
		"expression": {role.Call, role.Receiver},
		"arguments": {role.Call, role.Argument, role.Positional},
	}, role.Statement, role.Call, role.Base, role.Incomplete),

	mapAST("SuperMethodInvocation", Obj{
		"qualifier": OptObjectRoles("qname"),
		"name": OptObjectRoles("name"),
		"arguments": Each("args", ObjectRoles("arg")),
	}, Obj{
		"qualifier": OptObjectRoles("qname", role.Call, role.Callee),
		"name": OptObjectRoles("name", role.Call, role.Callee),
		"arguments": Each("args", ObjectRoles("arg", role.Call, role.Argument, role.Positional)),
	}, role.Expression, role.Call, role.Base),

	// Conditionals

	mapAST("IfStatement", Obj{
		"expression": ObjectRoles("expr"),
		"thenStatement": ObjectRoles("then"),
		"elseStatement": OptObjectRoles("else"),
	}, Obj{
		uast.KeyToken: String("if"),
		"expression": ObjectRoles("expr", role.If, role.Condition),
		"thenStatement": ObjectRoles("then", role.If, role.Then, role.Body),
		"elseStatement": OptObjectRoles("else", role.If, role.Else, role.Body),
	},role.Statement, role.If),
	
	annotateType("ConditionalExpression", objRoles{
		"expression": {role.If, role.Condition},
		"thenExpression": {role.If, role.Then},
		"elseExpression": {role.If, role.Else},
	},role.Expression, role.If),


	mapAST("SwitchStatement", Obj{
		"expression": ObjectRoles("expr"),
		"statements": Var("stmts"),
	}, Obj{
		"expression": ObjectRoles("expr", role.Switch, role.Expression),
		"statements": opSwitchStmtGroup{vr:"stmts"}, // will add "body" field to SwitchCase
	},role.Statement, role.Switch),

	// Loops
	annotateType("EnhancedForStatement", objRoles{
		"parameter":  {role.For, role.Iterator},
		"expression": {role.Expression, role.For},
		"body":       {role.For, role.Body},
	}, role.Statement, role.For, role.Iterator),

	mapAST("ForStatement", Obj{
		"initializers": Each("inits", ObjectRoles("init",)),
		"expression":   ObjectRoles("expr"),
		"updaters":     Each("upds", ObjectRoles("upd")),
		"body":         ObjectRoles("body"),
	}, Obj{
		"initializers": Each("inits", ObjectRoles("init", role.For, role.Initialization)),
		"expression":   ObjectRoles("expr", role.Expression, role.For, role.Condition),
		"updaters":     Each("upds", ObjectRoles("upd",role.For, role.Update)),
		"body":         ObjectRoles("body",role.For, role.Body),
	}, role.Statement, role.For),

	annotateType("WhileStatement", objRoles{
		"expression": {role.Expression, role.While, role.Condition},
		"body":       {role.While, role.Body},
	}, role.Statement, role.While),

	annotateType("DoStatement", objRoles{
		"expression": {role.DoWhile, role.Condition},
		"body":       {role.DoWhile, role.Body},
	}, role.Statement, role.DoWhile),

	// Operators
	mapASTCustom("InfixExpression", Obj{
		"leftOperand": ObjectRoles("left"),
		"rightOperand":ObjectRoles("right"),
		"operator": Var("op"),
	}, Fields{ // ->
		{Name:"operator", Op: Operator("op", infixRoles, role.Binary)},
		{Name: "leftOperand", Op: ObjectRoles("left", role.Expression, role.Binary, role.Left)},
		{Name: "rightOperand", Op: ObjectRoles("right", role.Expression, role.Binary, role.Right)},
	}, LookupArrOpVar("op", infixRoles), role.Expression, role.Binary, role.Operator),

	mapASTCustom("PostfixExpression", Obj{
		"operator": Var("op"),
	}, Fields{ // ->
		{Name:"operator", Op: Operator("op", postfixRoles, role.Unary, role.Postfix)},
	}, LookupArrOpVar("op", postfixRoles), role.Expression, role.Operator, role.Unary, role.Postfix),

	mapASTCustom("PrefixExpression", Obj{
		"operator": Var("op"),
	}, Fields{ // ->
		{Name:"operator", Op: Operator("op", prefixRoles, role.Unary)},
	}, LookupArrOpVar("op", prefixRoles), role.Expression, role.Operator, role.Unary),

	mapASTCustom("Assignment", Obj{
		"leftHandSide": ObjectRoles("left"),
		"rightHandSide":ObjectRoles("right"),
		"operator": Var("op"),
	}, Fields{ // ->
		{Name:"operator", Op: Operator("op", assignRoles, role.Assignment, role.Binary)},
		{Name: "leftHandSide", Op: ObjectRoles("left", role.Assignment, role.Binary, role.Left)},
		{Name: "rightHandSide", Op: ObjectRoles("right", role.Assignment, role.Binary, role.Right)},
	}, LookupArrOpVar("op", assignRoles), role.Expression, role.Assignment, role.Operator, role.Binary),

	// Types
	annotateType("ArrayType", nil, role.Type, role.Primitive, role.List),
	annotateType("IntersectionType", nil, role.Type, role.And),
	annotateType("NameQualifiedType", nil, role.Type, role.Name, role.Qualified),
	annotateType("ParameterizedType", nil, role.Type, role.Incomplete),
	annotateType("QualifiedType", nil, role.Type, role.Qualified),
	annotateType("SimpleType", nil, role.Type),
	annotateType("UnionType", nil, role.Type, role.Or),
	annotateType("WildcardType", nil, role.Type, role.Incomplete),


	// Modifiers
	mapASTCustom("Modifier", Obj{
		"keyword": Var("mod"),
	}, Obj{
		uast.KeyToken: Var("mod"),
	}, LookupArrOpVar("mod", modifierRoles)),

	// Exceptions
	mapAST("TryStatement", Obj{
		"resources": OptObjectRoles("res"),
		"body": ObjectRoles("body"),
		"catchClauses": Each("catches", ObjectRoles("catch")),
		"finally": OptObjectRoles("finally"),
	}, Obj{ // ->
		"resources": OptObjectRoles("res", role.Try),
		"body": ObjectRoles("body", role.Try, role.Body),
		"catchClauses": Each("catches", ObjectRoles("catch", role.Try, role.Catch)),
		"finally": OptObjectRoles("finally", role.Try, role.Finally),
	}, role.Statement, role.Try),

	annotateType("ThrowStatement", nil, role.Statement, role.Throw),

	annotateType("AssertStatement", nil, role.Statement, role.Assert),

	// Annotations
	annotateType("MarkerAnnotation", nil, role.Annotation, role.Incomplete),
	annotateType("NormalAnnotation", nil, role.Annotation, role.Incomplete),
	annotateType("SingleMemberAnnotation", nil, role.Annotation, role.Incomplete),
	annotateType("MemberValuePair", nil, role.Annotation, role.Incomplete),

	// Comments
	mapAST("BlockComment", Obj{
		"text": UncommentCLike("text"),
	}, Obj{
		uast.KeyToken: Var("text"),
	}, role.Comment),
	mapAST("Javadoc", Fields{
		{Name:"text", Op: UncommentCLike("text"), Optional:"txt"},
	}, Fields{
		{Name:uast.KeyToken, Op: Var("text"), Optional:"txt"},
	}, role.Documentation, role.Comment),
	mapAST("LineComment", Obj{
		"text": UncommentCLike("text"),
	}, Obj{
		uast.KeyToken: Var("text"),
	}, role.Comment),

	// Javadoc tags
	annotateType("MemberRef", nil, role.Documentation, role.Variable, role.Incomplete),
	annotateType("MethodRef", nil, role.Documentation, role.Function, role.Incomplete),
	annotateType("MethodRefParameter", nil, role.Documentation, role.Function, role.Incomplete),
	annotateType("TagElement", nil, role.Documentation, role.Incomplete),
	annotateType("TextElement", nil, role.Documentation, role.Incomplete),

	// Other expressions
	annotateType("ArrayAccess", nil, role.Expression, role.Incomplete),
	annotateType("ArrayCreation", nil, role.Expression, role.Incomplete),
	annotateType("CastExpression", nil, role.Expression, role.Incomplete),
	annotateType("CreationReference", nil, role.Expression, role.Incomplete),
	annotateType("ExpressionMethodReference", nil, role.Expression, role.Incomplete),
	annotateType("ParenthesizedExpression", nil, role.Expression, role.Incomplete),
	annotateType("SuperMethodReference", nil, role.Expression, role.Incomplete),
	annotateType("ThisExpression", nil, role.Expression, role.This),

	// Other statements
	annotateType("Block", nil, role.Statement, role.Block, role.Scope),
	annotateType("BreakStatement", nil, role.Statement, role.Break),
	annotateType("EmptyStatement", nil, role.Statement),
	annotateType("ExpressionStatement", nil, role.Statement),
	annotateType("LabeledStatement", nil, role.Statement, role.Incomplete),
	annotateType("ReturnStatement", nil, role.Statement, role.Return),
	annotateType("SynchronizedStatement", nil, role.Statement, role.Incomplete),

	// Others
	annotateType("ArrayInitializer", nil, role.Expression, role.List, role.Literal),
	annotateType("Dimension", nil, role.Type, role.Incomplete),
	annotateType("TypeParameter", nil, role.Type, role.Incomplete),
}

var (
	caseMap = mapAST("SwitchCase", Obj{
		"expression": ObjectRoles("expr"),
		"body": Each("stmts", ObjectRoles("stmt")),
	}, Obj{
		"expression": ObjectRoles("expr", role.Expression, role.Switch, role.Case, role.Condition),
		"body": Each("stmts", ObjectRoles("stmt", role.Switch, role.Case, role.Body)),
	}, role.Statement, role.Switch, role.Case)

	defaultMap = mapAST("SwitchCase", Obj{
		"expression": Is(nil),
		"body": Each("stmts", ObjectRoles("stmt")),
	}, Obj{
		"expression": Is(nil),
		"body": Each("stmts", ObjectRoles("stmt", role.Switch, role.Case, role.Body)),
	}, role.Statement, role.Switch, role.Default)
)

var (
	casesMapping = Mappings(caseMap, defaultMap)
	casesRevMapping = Mappings(caseMap.Reverse(), defaultMap.Reverse())
)

var _ Op = opSwitchStmtGroup{}

type opSwitchStmtGroup struct {
	vr string
}

func (op opSwitchStmtGroup) Check(st *State, n uast.Node) (bool, error) {
	cases, ok := n.(uast.List)
	if !ok {
		return false, nil
	}
	var out uast.List
	for _, c := range cases {
		co, ok := c.(uast.Object)
		if !ok || co.Type() != "SwitchCase" {
			return false, nil
		}
		so, ok := co["body"]
		if !ok {
			return false, nil
		}
		cs := co.CloneObject()
		delete(cs, "body")
		out = append(out, cs)
		if so == nil {
			continue
		}
		stmts, ok := so.(uast.List)
		if !ok {
			return false, nil
		}
		out = append(out, stmts...)
	}
	for i, c := range out {
		cn, err := casesRevMapping.Do(c)
		if err != nil {
			return false, err
		}
		out[i] = cn
	}
	if err := st.SetVar(op.vr, out); err != nil {
		return false, err
	}
	return true, nil
}

func (op opSwitchStmtGroup) Construct(st *State, _ uast.Node) (uast.Node, error) {
	o, err :=  st.MustGetVar(op.vr)
	if err != nil {
		return nil, err
	}
	stmts, ok := o.(uast.List)
	if !ok {
		return nil, ErrExpectedList.New(o)
	}
	var out uast.List
	var (
		ccase uast.Object
		cur uast.List
	)
	for _, s := range stmts {
		so, ok := s.(uast.Object)
		if !ok {
			return nil, ErrExpectedObject.New(s)
		}
		if so.Type() == "SwitchCase" {
			if ccase != nil {
				ccase["body"] = cur
				out = append(out, ccase)
			}
			ccase = so
			cur = nil
		} else {
			cur = append(cur, s)
		}
	}
	if ccase != nil {
		ccase["body"] = cur
		out = append(out, ccase)
	}
	for i, c := range out {
		cn, err := casesMapping.Do(c)
		if err != nil {
			return out, err
		}
		out[i] = cn
	}
	return out, nil
}
