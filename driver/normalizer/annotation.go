package normalizer

import (
	"github.com/bblfsh/sdk/v3/uast"
	"github.com/bblfsh/sdk/v3/uast/nodes"
	"github.com/bblfsh/sdk/v3/uast/role"
	. "github.com/bblfsh/sdk/v3/uast/transformer"
)

// Native is the of list `transformer.Transformer` to apply to a native AST.
// To learn more about the Transformers and the available ones take a look to:
// https://godoc.org/github.com/bblfsh/sdk/v3/uast/transformer
var Native = Transformers([][]Transformer{
	{
		// ResponseMetadata is a transform that trims response metadata from AST.
		//
		// https://godoc.org/github.com/bblfsh/sdk/v3/uast#ResponseMetadata
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

var (
	modifierRoles = StringToRolesMap(map[string][]role.Role{
		"public":    {role.Visibility, role.World},
		"protected": {role.Visibility, role.Subtype},
		"private":   {role.Visibility, role.Instance},

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
		"volatile":  {role.Incomplete},
		// method
		"synchronized": {role.Incomplete},
		"native":       {role.Incomplete},
		// interface method
		"default": {role.Incomplete},
	})
	infixRoles = StringToRolesMap(map[string][]role.Role{
		"+":   {role.Arithmetic, role.Add},
		"-":   {role.Arithmetic, role.Substract},
		"*":   {role.Arithmetic, role.Multiply},
		"/":   {role.Arithmetic, role.Divide},
		"%":   {role.Arithmetic, role.Modulo},
		"<<":  {role.Bitwise, role.LeftShift},
		">>":  {role.Bitwise, role.RightShift},
		">>>": {role.Bitwise, role.RightShift, role.Unsigned},
		"<":   {role.LessThan, role.Relational},
		">":   {role.GreaterThan, role.Relational},
		"<=":  {role.LessThanOrEqual, role.Relational},
		">=":  {role.GreaterThanOrEqual, role.Relational},
		"==":  {role.Equal, role.Relational},
		"!=":  {role.Equal, role.Not, role.Relational},
		"&":   {role.Bitwise, role.And},
		"|":   {role.Bitwise, role.Or},
		"&&":  {role.Boolean, role.And},
		"||":  {role.Boolean, role.Or},
		"^":   {role.Boolean, role.Xor},
	})
	postfixRoles = StringToRolesMap(map[string][]role.Role{
		"++": {role.Arithmetic, role.Increment},
		"--": {role.Arithmetic, role.Decrement},
	})
	prefixRoles = StringToRolesMap(map[string][]role.Role{
		"++": {role.Arithmetic, role.Increment},
		"--": {role.Arithmetic, role.Decrement},
		"+":  {role.Arithmetic, role.Positive},
		"-":  {role.Arithmetic, role.Negative},
		"~":  {role.Bitwise, role.Not},
		"!":  {role.Boolean, role.Not},
	})
	assignRoles = StringToRolesMap(map[string][]role.Role{
		"=":    {},
		"+=":   {role.Arithmetic, role.Add},
		"-=":   {role.Arithmetic, role.Substract},
		"*=":   {role.Arithmetic, role.Multiply},
		"/=":   {role.Arithmetic, role.Divide},
		"%=":   {role.Arithmetic, role.Modulo},
		"&=":   {role.Bitwise, role.And},
		"|=":   {role.Bitwise, role.Or},
		"^=":   {role.Bitwise, role.Xor},
		"<<=":  {role.Bitwise, role.LeftShift},
		">>=":  {role.Bitwise, role.RightShift},
		">>>=": {role.Bitwise, role.RightShift, role.Unsigned},
	})
	variadicRoles = StringToRolesMap(map[string][]role.Role{
		"true":  {role.ArgsList},
		"false": {},
	})
	primitiveRoles = StringToRolesMap(map[string][]role.Role{
		"boolean": {role.Boolean},
		"byte":    {role.Byte},
		"char":    {role.Character},
		"short":   {role.Number},
		"int":     {role.Number},
		"long":    {role.Number},
		"float":   {role.Number},
		"double":  {role.Number},
		"void":    {},
	})
)

func annotateModifiers(typ string, mod string, roles ...role.Role) Mapping {
	c := AnyElem(Has{
		uast.KeyToken: String(mod),
	})
	if mod == "" {
		c = All(Not(Has{
			uast.KeyToken: In(
				nodes.String("public"),
				nodes.String("private"),
				nodes.String("protected"),
			),
		}))
	}
	return AnnotateType(typ, ObjMap{
		"modifiers": Map(
			Check(c, Var("mod")),
			Var("mod"),
		),
	}, roles...)
}

// Annotations is a list of individual transformations to annotate a native AST with roles.
var Annotations = []Mapping{
	AnnotateType("CompilationUnit", nil, role.File),

	// Names
	AnnotateType("QualifiedName", nil, role.Expression, role.Identifier, role.Qualified),
	AnnotateType("SimpleName",
		FieldRoles{
			"identifier": {Rename: uast.KeyToken},
		},
		role.Expression, role.Identifier,
	),
	AnnotateTypeCustom("PrimitiveType",
		FieldRoles{
			"primitiveTypeCode": {Rename: uast.KeyToken, Op: Var("typ")},
		},
		LookupArrOpVar("typ", primitiveRoles),
		role.Type, role.Primitive,
	),

	// Visibility
	AnnotateType("TypeDeclaration",
		ObjMap{"modifiers": Is(nil)},
		role.Visibility, role.Package,
	),
	AnnotateType("MethodDeclaration",
		ObjMap{"modifiers": Is(nil)},
		role.Visibility, role.Package,
	),

	annotateModifiers("TypeDeclaration", "", role.Visibility, role.Package),
	annotateModifiers("TypeDeclaration", "public", role.Visibility, role.World),
	annotateModifiers("TypeDeclaration", "private", role.Visibility, role.Type),
	annotateModifiers("TypeDeclaration", "protected", role.Visibility, role.Subtype),

	annotateModifiers("MethodDeclaration", "", role.Visibility, role.Package),
	annotateModifiers("MethodDeclaration", "public", role.Visibility, role.World),
	annotateModifiers("MethodDeclaration", "private", role.Visibility, role.Type),
	annotateModifiers("MethodDeclaration", "protected", role.Visibility, role.Subtype),

	// Package and imports
	AnnotateType("ImportDeclaration", nil, role.Declaration, role.Import),
	AnnotateType("ImportDeclaration", FieldRoles{
		"name": {Sub: ObjMap{
			uast.KeyType: String("QualifiedName"),
		}, Roles: role.Roles{role.Pathname, role.Import}},
	}),
	AnnotateType("ImportDeclaration", FieldRoles{
		"qualifier": {Sub: ObjMap{
			uast.KeyType: String("QualifiedName"),
		}, Roles: role.Roles{role.Pathname, role.Import}},
	}),

	AnnotateType("PackageDeclaration",
		FieldRoles{
			uast.KeyToken: {Add: true, Op: String("package")},
		},
		role.Declaration, role.Package,
	),

	// Type declarations
	AnnotateType("AnonymousClassDeclaration",
		FieldRoles{
			"bodyDeclarations": {Arr: true, Roles: role.Roles{role.Body}},
		},
		role.Expression, role.Declaration, role.Type, role.Anonymous,
	),

	AnnotateType("AnnotationTypeDeclaration",
		FieldRoles{
			"bodyDeclarations": {Arr: true, Roles: role.Roles{role.Body}},
		},
		role.Declaration, role.Type, role.Annotation,
	),
	AnnotateType("EnumDeclaration", nil, role.Declaration, role.Type, role.Enumeration),

	// ClassDeclaration | InterfaceDeclaration
	AnnotateType("TypeDeclaration", nil, role.Declaration, role.Type),
	// Local (TypeDeclaration | EnumDeclaration)
	AnnotateType("TypeDeclarationStatement", nil, role.Statement, role.Declaration, role.Type),

	// Method declarations
	AnnotateType("MethodDeclaration", MapObj(Fields{
		{Name: "name", Op: ObjectRoles("name")},
		{Name: "body", Op: OptObjectRoles("body")},
		{Name: "parameters", Op: Each("param", ObjectRolesCustom("p", Obj{
			"name":    ObjectRoles("pname"),
			"varargs": Var("variadic"),
		}))},
	}, Fields{ // ->
		{Name: "name", Op: ObjectRoles("name", role.Function, role.Name)},
		{Name: "body", Op: OptObjectRoles("body", role.Function, role.Body)},
		{Name: "parameters", Op: Each("param", ObjectRolesCustomOp("p", Obj{
			"name":    ObjectRoles("pname", role.Function, role.Name, role.Argument),
			"varargs": Var("variadic"),
		}, LookupArrOpVar("variadic", variadicRoles), role.Function, role.Argument))},
	}), role.Declaration, role.Function),

	AnnotateType("LambdaExpression", FieldRoles{
		"body": {Roles: role.Roles{role.Function, role.Body}},
		"parameters": {Arr: true, Sub: FieldRoles{
			"name": {Roles: role.Roles{role.Function, role.Name, role.Argument}},
		}, Roles: role.Roles{role.Function, role.Argument}},
	}, role.Declaration, role.Function, role.Anonymous),

	AnnotateType("TypeMethodReference",
		FieldRoles{
			"name":          {Roles: role.Roles{role.Function, role.Name}},
			"type":          {Roles: role.Roles{role.Function, role.Return}},
			"typeArguments": {Arr: true, Roles: role.Roles{role.Function, role.Argument}},
		},
		role.Declaration, role.Function,
	),

	// Other declarations
	AnnotateType("AnnotationTypeMemberDeclaration", nil, role.Declaration, role.Type, role.Annotation),
	AnnotateType("EnumConstantDeclaration", nil, role.Declaration, role.Enumeration),
	AnnotateType("FieldDeclaration", nil, role.Declaration, role.Variable),
	// TODO: differentiate between static (class) and instance initialization
	AnnotateType("Initializer", nil, role.Initialization, role.Block, role.Incomplete),
	AnnotateType("SingleVariableDeclaration", nil, role.Declaration, role.Variable),
	AnnotateType("VariableDeclarationExpression", nil, role.Expression, role.Declaration, role.Variable),
	AnnotateType("VariableDeclarationFragment", nil, role.Declaration, role.Variable),
	AnnotateType("VariableDeclarationStatement", nil, role.Statement, role.Declaration, role.Variable),

	// Literals
	AnnotateType("BooleanLiteral", nil, role.Expression, role.Literal, role.Boolean),
	AnnotateType("TypeLiteral", nil, role.Expression, role.Literal, role.Type),
	AnnotateType("NumberLiteral", nil, role.Expression, role.Literal, role.Number),

	AnnotateType("NullLiteral",
		FieldRoles{
			uast.KeyToken: {Add: true, Op: String("null")},
		},
		role.Expression, role.Literal, role.Null,
	),
	AnnotateType("StringLiteral",
		FieldRoles{
			"escapedValue": {Rename: uast.KeyToken},
		},
		role.Expression, role.Literal, role.String,
	),
	AnnotateType("CharacterLiteral",
		FieldRoles{
			"escapedValue": {Rename: uast.KeyToken},
		},
		role.Expression, role.Literal, role.Character,
	),

	// Calls
	AnnotateType("ClassInstanceCreation",
		FieldRoles{
			"type":      {Roles: role.Roles{role.Call, role.Callee}},
			"arguments": {Arr: true, Roles: role.Roles{role.Call, role.Argument, role.Positional}},
		},
		role.Expression, role.Call, role.Instance,
	),
	AnnotateType("ConstructorInvocation",
		FieldRoles{
			//"type": ObjectRoles("type", role.Call, role.Callee),
			"arguments": {Arr: true, Roles: role.Roles{role.Call, role.Argument, role.Positional}},
		},
		role.Statement, role.Call, role.Incomplete,
	),
	AnnotateType("MethodInvocation",
		FieldRoles{
			"expression": {Opt: true, Roles: role.Roles{role.Call, role.Receiver}},
			"name":       {Opt: true, Roles: role.Roles{role.Call, role.Callee}},
			"arguments":  {Arr: true, Roles: role.Roles{role.Call, role.Argument, role.Positional}},
		},
		role.Expression, role.Call,
	),
	AnnotateType("SuperConstructorInvocation",
		ObjRoles{
			"expression": {role.Call, role.Receiver},
			"arguments":  {role.Call, role.Argument, role.Positional},
		},
		role.Statement, role.Call, role.Base, role.Incomplete,
	),
	AnnotateType("SuperMethodInvocation",
		FieldRoles{
			"qualifier": {Opt: true, Roles: role.Roles{role.Call, role.Callee}},
			"name":      {Opt: true, Roles: role.Roles{role.Call, role.Callee}},
			"arguments": {Arr: true, Roles: role.Roles{role.Call, role.Argument, role.Positional}},
		},
		role.Expression, role.Call, role.Base,
	),

	// Conditionals
	AnnotateType("IfStatement",
		FieldRoles{
			uast.KeyToken:   {Add: true, Op: String("if")},
			"expression":    {Roles: role.Roles{role.If, role.Condition}},
			"thenStatement": {Roles: role.Roles{role.If, role.Then, role.Body}},
			"elseStatement": {Opt: true, Roles: role.Roles{role.If, role.Else, role.Body}},
		},
		role.Statement, role.If,
	),
	AnnotateType("ConditionalExpression",
		ObjRoles{
			"expression":     {role.If, role.Condition},
			"thenExpression": {role.If, role.Then},
			"elseExpression": {role.If, role.Else},
		},
		role.Expression, role.If,
	),

	AnnotateType("SwitchStatement", MapObj(Obj{
		"expression": ObjectRoles("expr"),
		"statements": Var("stmts"),
	}, Obj{
		"expression": ObjectRoles("expr", role.Switch, role.Expression),
		"statements": opSwitchStmtGroup{vr: "stmts"}, // will add "body" field to SwitchCase
	}), role.Statement, role.Switch),

	// Loops
	AnnotateType("EnhancedForStatement",
		ObjRoles{
			"parameter":  {role.For, role.Iterator},
			"expression": {role.Expression, role.For},
			"body":       {role.For, role.Body},
		},
		role.Statement, role.For, role.Iterator,
	),
	AnnotateType("ForStatement",
		FieldRoles{
			"initializers": {Arr: true, Roles: role.Roles{role.For, role.Initialization}},
			"expression":   {Roles: role.Roles{role.Expression, role.For, role.Condition}},
			"updaters":     {Arr: true, Roles: role.Roles{role.For, role.Update}},
			"body":         {Roles: role.Roles{role.For, role.Body}},
		},
		role.Statement, role.For,
	),
	AnnotateType("WhileStatement",
		ObjRoles{
			"expression": {role.Expression, role.While, role.Condition},
			"body":       {role.While, role.Body},
		},
		role.Statement, role.While,
	),
	AnnotateType("DoStatement",
		ObjRoles{
			"expression": {role.DoWhile, role.Condition},
			"body":       {role.DoWhile, role.Body},
		},
		role.Statement, role.DoWhile,
	),

	// Operators
	AnnotateTypeCustom("InfixExpression",
		MapObj(Obj{
			"leftOperand":  ObjectRoles("left"),
			"rightOperand": ObjectRoles("right"),
			"operator":     Var("op"),
		}, Fields{ // ->
			{Name: "operator", Op: Operator("op", infixRoles, role.Binary)},
			{Name: "leftOperand", Op: ObjectRoles("left", role.Expression, role.Binary, role.Left)},
			{Name: "rightOperand", Op: ObjectRoles("right", role.Expression, role.Binary, role.Right)},
		}),
		LookupArrOpVar("op", infixRoles), role.Expression, role.Binary, role.Operator,
	),

	AnnotateTypeCustom("PostfixExpression",
		MapObj(Obj{
			"operator": Var("op"),
		}, Fields{ // ->
			{Name: "operator", Op: Operator("op", postfixRoles, role.Unary, role.Postfix)},
		}),
		LookupArrOpVar("op", postfixRoles), role.Expression, role.Operator, role.Unary, role.Postfix,
	),

	AnnotateTypeCustom("PrefixExpression",
		MapObj(Obj{
			"operator": Var("op"),
		}, Fields{ // ->
			{Name: "operator", Op: Operator("op", prefixRoles, role.Unary)},
		}),
		LookupArrOpVar("op", prefixRoles), role.Expression, role.Operator, role.Unary,
	),

	AnnotateTypeCustom("Assignment",
		MapObj(Obj{
			"leftHandSide":  ObjectRoles("left"),
			"rightHandSide": ObjectRoles("right"),
			"operator":      Var("op"),
		}, Fields{ // ->
			{Name: "operator", Op: Operator("op", assignRoles, role.Assignment, role.Binary)},
			{Name: "leftHandSide", Op: ObjectRoles("left", role.Assignment, role.Binary, role.Left)},
			{Name: "rightHandSide", Op: ObjectRoles("right", role.Assignment, role.Binary, role.Right)},
		}),
		LookupArrOpVar("op", assignRoles), role.Expression, role.Assignment, role.Operator, role.Binary,
	),

	// Types
	AnnotateType("ArrayType", nil, role.Type, role.Primitive, role.List),
	AnnotateType("IntersectionType", nil, role.Type, role.And),
	AnnotateType("NameQualifiedType", nil, role.Type, role.Name, role.Qualified),
	AnnotateType("ParameterizedType", nil, role.Type, role.Incomplete),
	AnnotateType("QualifiedType", nil, role.Type, role.Qualified),
	AnnotateType("SimpleType", nil, role.Type),
	AnnotateType("UnionType", nil, role.Type, role.Or),
	AnnotateType("WildcardType", nil, role.Type, role.Incomplete),

	// Modifiers
	AnnotateTypeCustom("Modifier",
		FieldRoles{
			"keyword": {Rename: uast.KeyToken, Op: Var("mod")},
		},
		LookupArrOpVar("mod", modifierRoles),
	),

	// Exceptions
	AnnotateType("TryStatement",
		FieldRoles{
			"resources":    {Opt: true, Roles: role.Roles{role.Try}},
			"body":         {Roles: role.Roles{role.Try, role.Body}},
			"catchClauses": {Arr: true, Roles: role.Roles{role.Try, role.Catch}},
			"finally":      {Opt: true, Roles: role.Roles{role.Try, role.Finally}},
		},
		role.Statement, role.Try,
	),
	AnnotateType("ThrowStatement", nil, role.Statement, role.Throw),
	AnnotateType("AssertStatement", nil, role.Statement, role.Assert),

	// Annotations
	AnnotateType("MarkerAnnotation", nil, role.Annotation, role.Incomplete),
	AnnotateType("NormalAnnotation", nil, role.Annotation, role.Incomplete),
	AnnotateType("SingleMemberAnnotation", nil, role.Annotation, role.Incomplete),
	AnnotateType("MemberValuePair", nil, role.Annotation, role.Incomplete),

	// Comments
	AnnotateType("BlockComment", MapObj(Obj{
		"text": UncommentCLike("text"),
	}, Obj{
		uast.KeyToken: Var("text"),
	}), role.Comment),

	AnnotateType("Javadoc", MapObj(Fields{
		{Name: "text", Op: UncommentCLike("text"), Optional: "txt"},
	}, Fields{
		{Name: uast.KeyToken, Op: Var("text"), Optional: "txt"},
	}), role.Documentation, role.Comment),

	AnnotateType("LineComment", MapObj(Obj{
		"text": UncommentCLike("text"),
	}, Obj{
		uast.KeyToken: Var("text"),
	}), role.Comment),

	// Javadoc tags
	AnnotateType("MemberRef", nil, role.Documentation, role.Variable, role.Incomplete),
	AnnotateType("MethodRef", nil, role.Documentation, role.Function, role.Incomplete),
	AnnotateType("MethodRefParameter", nil, role.Documentation, role.Function, role.Incomplete),
	AnnotateType("TagElement", nil, role.Comment, role.Documentation, role.Incomplete),
	AnnotateType("TextElement", nil, role.Comment, role.Documentation, role.Incomplete),

	// Other expressions
	AnnotateType("ArrayAccess", nil, role.Expression, role.Incomplete),
	AnnotateType("ArrayCreation", nil, role.Expression, role.Incomplete),
	AnnotateType("CastExpression", nil, role.Expression, role.Incomplete),
	AnnotateType("CreationReference", nil, role.Expression, role.Incomplete),
	AnnotateType("ExpressionMethodReference", nil, role.Expression, role.Incomplete),
	AnnotateType("ParenthesizedExpression", nil, role.Expression, role.Incomplete),
	AnnotateType("SuperMethodReference", nil, role.Expression, role.Incomplete),
	AnnotateType("ThisExpression", nil, role.Expression, role.This),

	// Other statements
	AnnotateType("Block", nil, role.Statement, role.Block, role.Scope),
	AnnotateType("BreakStatement", nil, role.Statement, role.Break),
	AnnotateType("EmptyStatement", nil, role.Statement),
	AnnotateType("ExpressionStatement", nil, role.Statement),
	AnnotateType("LabeledStatement", nil, role.Statement, role.Incomplete),
	AnnotateType("ReturnStatement", nil, role.Statement, role.Return),
	AnnotateType("SynchronizedStatement", nil, role.Statement, role.Incomplete),

	// Others
	AnnotateType("ArrayInitializer", nil, role.Expression, role.List, role.Literal),
	AnnotateType("Dimension", nil, role.Type, role.Incomplete),
	AnnotateType("TypeParameter", nil, role.Type, role.Incomplete),
}

var (
	caseMap = AnnotateType("SwitchCase",
		FieldRoles{
			"expression": {Roles: role.Roles{role.Expression, role.Switch, role.Case, role.Condition}},
			"body":       {Arr: true, Roles: role.Roles{role.Switch, role.Case, role.Body}},
		},
		role.Statement, role.Switch, role.Case,
	)
	defaultMap = AnnotateType("SwitchCase",
		FieldRoles{
			"expression": {Op: Is(nil)},
			"body":       {Arr: true, Roles: role.Roles{role.Switch, role.Case, role.Body}},
		},
		role.Statement, role.Switch, role.Default,
	)
)

var (
	casesMapping    = Mappings(caseMap, defaultMap)
	casesRevMapping = Mappings(Reverse(caseMap), Reverse(defaultMap))
)

var _ Op = opSwitchStmtGroup{}

type opSwitchStmtGroup struct {
	vr string
}

func (opSwitchStmtGroup) Kinds() nodes.Kind {
	return nodes.KindArray
}

func (op opSwitchStmtGroup) Check(st *State, n nodes.Node) (bool, error) {
	cases, ok := n.(nodes.Array)
	if !ok && n != nil {
		return false, nil
	}
	var out nodes.Array
	for _, c := range cases {
		co, ok := c.(nodes.Object)
		if !ok || uast.TypeOf(co) != "SwitchCase" {
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
		stmts, ok := so.(nodes.Array)
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

func (op opSwitchStmtGroup) Construct(st *State, _ nodes.Node) (nodes.Node, error) {
	o, err := st.MustGetVar(op.vr)
	if err != nil {
		return nil, err
	}
	stmts, ok := o.(nodes.Array)
	if !ok && o != nil {
		return nil, ErrExpectedList.New(o)
	}
	var out nodes.Array
	var (
		ccase nodes.Object
		cur   nodes.Array
	)
	for _, s := range stmts {
		so, ok := s.(nodes.Object)
		if !ok {
			return nil, ErrExpectedObject.New(s)
		}
		if uast.TypeOf(so) == "SwitchCase" {
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
