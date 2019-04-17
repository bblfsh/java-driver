package normalizer

import (
	"github.com/bblfsh/sdk/v3/uast"
	. "github.com/bblfsh/sdk/v3/uast/transformer"
)

var Preprocess []Transformer

// PreprocessCode is a preprocessor stage that can use the source code to
// fix tokens and positional information.
// Java already provides all the information we need.
var PreprocessCode []CodeTransformer

var Normalize = Transformers([][]Transformer{
	// The main block of normalization rules.
	{Mappings(Normalizers...)},
}...)

var Normalizers = []Mapping{
	MapSemantic("StringLiteral", uast.String{}, MapObj(
		Fields{
			{Name: "unescapedValue", Op: Var("val")},
			{Name: "escapedValue", Drop: true, Op: Any()}, // only used in Annotated
		},
		Obj{
			"Value":  Var("val"),
			"Format": String(""),
		},
	)),
	MapSemantic("SimpleName", uast.Identifier{}, MapObj(
		Obj{
			"identifier": Var("name"),
		},
		Obj{
			"Name": Var("name"),
		},
	)),
	MapSemantic("QualifiedName", uast.QualifiedIdentifier{}, MapObj(
		CasesObj("case",
			// common
			Obj{"name": Var("name")},
			Objs{
				// the last name = identifier
				{
					"qualifier": Check(Has{
						uast.KeyType: String(uast.TypeOf(uast.Identifier{})),
					}, Var("par")),
				},
				// linked list
				{
					"qualifier": UASTType(uast.QualifiedIdentifier{}, Obj{
						// FIXME: start position
						uast.KeyPos: AnyNode(nil),
						"Names":     Var("names"),
					}),
				},
			},
		),
		CasesObj("case", nil,
			Objs{
				// the last name = identifier
				{
					"Names": Arr(Var("par"), Var("name")),
				},
				// linked list
				{
					"Names": Append(Var("names"), Arr(Var("name"))),
				},
			},
		),
	)),
	MapSemantic("BlockComment", uast.Comment{}, MapObj(
		Obj{
			"text": CommentText([2]string{"/*", "*/"}, "comm"),
		},
		CommentNode(true, "comm", nil),
	)),
	MapSemantic("LineComment", uast.Comment{}, MapObj(
		Obj{
			"text": CommentText([2]string{"//", ""}, "comm"),
		},
		CommentNode(false, "comm", nil),
	)),
	MapSemantic("Block", uast.Block{}, MapObj(
		Obj{
			"statements": Var("stmts"),
		},
		Obj{
			"Statements": Var("stmts"),
		},
	)),
	MapSemantic("ImportDeclaration", uast.Import{}, MapObj(
		CasesObj("case",
			// common
			Obj{
				"static": Var("static"),
			},
			Objs{
				// star import (on demand)
				{
					"name":     Var("name"),
					"onDemand": String("true"),
				},
				// normal import
				{
					"name": Part("path", UASTType(uast.QualifiedIdentifier{}, Obj{
						"Names": Append(Var("names"), Arr(Var("name"))),
					})),
					"onDemand": String("false"),
				},
			},
		),
		CasesObj("case",
			// common
			Obj{
				// TODO: handle static when we have scopes
				"Target": Obj{"static": Var("static")},
			},
			Objs{
				// star import (on demand)
				{
					"Path":  Var("name"),
					"All":   Bool(true),
					"Names": Arr(),
				},
				{
					"Path": Part("path", UASTType(uast.QualifiedIdentifier{}, Obj{
						"Names": Var("names"),
					})),
					"All":   Bool(false),
					"Names": Arr(Var("name")),
				},
			},
		),
	)),

	MapSemantic("MethodDeclaration", uast.FunctionGroup{}, MapObj(
		Obj{
			"constructor":      Var("constr"),
			"extraDimensions2": Is(nil), // TODO: find an example
			"javadoc":          Var("doc"),
			"modifiers":        Var("ann"), // TODO: it's an array, we should expand it somewhere
			"name":             Var("name"),
			"body":             Var("body"),
			"parameters": Each("args", Obj{
				uast.KeyType:         String("SingleVariableDeclaration"),
				uast.KeyPos:          Var("apos"),
				"extraDimensions2":   Is(nil),
				"initializer":        Var("ainit"),
				"modifiers":          Is(nil),
				"name":               Var("aname"),
				"type":               Var("atype"),
				"varargs":            Cases("varg", String("false"), String("true")),
				"varargsAnnotations": Is(nil),
			}),
			"receiverQualifier": Var("recv_name"),
			"receiverType": Cases("recv",
				Is(nil),
				Check(NotNil(), Var("recv_type")),
			),
			"returnType2": Cases("out_case",
				// no return type (constructor)
				Is(nil),
				// void
				Obj{
					uast.KeyType:        String("PrimitiveType"),
					uast.KeyPos:         AnyNode(nil),
					"annotations":       Is(nil),
					"primitiveTypeCode": String("void"),
				},
				// any other type
				Check(
					Not(And(
						Is(nil),
						Has{
							uast.KeyType:        String("PrimitiveType"),
							"annotations":       Is(nil),
							"primitiveTypeCode": String("void"),
						},
					)),
					Var("out1"),
				),
			),
			"thrownExceptionTypes": Var("exc"),
			"typeParameters":       Var("tmpl"),
		},
		Obj{
			"Nodes": Arr(
				Var("doc"),
				Var("ann"),
				UASTType(uast.Alias{}, Obj{
					// FIXME: add position
					"Name": Var("name"),
					"Node": UASTType(uast.Function{}, Obj{
						"Type": UASTType(uast.FunctionType{}, Obj{
							"Arguments": Cases("recv",
								// default receiver
								argsPart,
								// additional receiver
								PrependOne(
									UASTType(uast.Argument{}, Obj{
										"Name":     Var("recv_name"),
										"Type":     Var("recv_type"),
										"Receiver": Bool(true),
									}),
									argsPart,
								),
							),
							"Returns": Cases("out_case",
								// no return (constructor)
								Is(nil),
								// void return
								Is(nil),
								// normal return type
								Arr(
									UASTType(uast.Argument{}, Obj{
										"Type": Var("out1"),
									}),
								),
							),
						}),
						"Body": Var("body"),
					}),
				}),
				Obj{ // FIXME: store them as annotations at least
					"constructor":          Var("constr"),
					"thrownExceptionTypes": Var("exc"),
					"typeParameters":       Var("tmpl"),
				},
			),
		},
	)),

	MapSemantic("Javadoc", uast.Comment{}, MapObj(
		Obj{
			"text": CommentText([2]string{"/**", "*/"}, "comm"),
		},
		CommentNode(true, "comm", nil),
	)),

	MapSemantic("Javadoc", uast.Group{}, MapObj(
		Obj{
			"tags": Var("parts"),
		},
		Obj{
			"Nodes": Var("parts"),
		},
	)),
	MapSemantic("TagElement", uast.Group{}, MapObj(
		Obj{
			"fragments": Var("parts"),
			"tagName":   Is(nil),
		},
		Obj{
			"Nodes": Var("parts"),
		},
	)),
	MapSemantic("TextElement", uast.Comment{}, MapObj(
		Obj{
			"text": CommentText([2]string{"", ""}, "comm"),
		},
		CommentNode(false, "comm", nil),
	)),
}

var argsPart = Each("args", UASTType(uast.Argument{}, Obj{
	uast.KeyPos: Var("apos"),
	"Name":      Var("aname"),
	"Type":      Var("atype"),
	"Init":      Var("ainit"),
	"Variadic":  Cases("varg", Bool(false), Bool(true)),
}))
