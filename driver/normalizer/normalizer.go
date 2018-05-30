package normalizer

import (
	"gopkg.in/bblfsh/sdk.v2/uast"
	. "gopkg.in/bblfsh/sdk.v2/uast/transformer"
)

var Preprocess []Transformer

var Normalize = Transformers([][]Transformer{
	// The main block of normalization rules.
	{Mappings(Normalizers...)},
}...)

var Normalizers = []Mapping{
	MapSemantic("", "StringLiteral", uast.String{}, nil,
		Obj{
			"escapedValue": Quote(Var("val")),
		},
		Obj{
			"Value":  Var("val"),
			"Format": String(""),
		},
	),
	MapSemantic("", "SimpleName", uast.Identifier{}, nil,
		Obj{
			"identifier": Var("name"),
		},
		Obj{
			"Name": Var("name"),
		},
	),
	MapSemantic("", "QualifiedName", uast.QualifiedIdentifier{}, nil,
		Obj{
			"name": Var("name"),
			"qualifier": Check(Has{
				uast.KeyType: String(uast.TypeOf(uast.Identifier{})),
			}, Var("par")),
		},
		Obj{
			"Names": Arr(Var("par"), Var("name")),
		},
	),
	MapSemantic("", "QualifiedName", uast.QualifiedIdentifier{}, nil,
		Obj{
			"name": Var("name"),
			"qualifier": UASTType(uast.QualifiedIdentifier{}, Obj{
				// FIXME: start position
				uast.KeyPos: AnyNode(nil),
				"Names":     Var("names"),
			}),
		},
		Obj{
			"Names": Append(Var("names"), Arr(Var("name"))),
		},
	),
	MapSemantic("", "BlockComment", uast.Comment{}, nil,
		Obj{
			"text": CommentText([2]string{"/*", "*/"}, "comm"),
		},
		CommentNode(true, "comm", nil),
	),
	MapSemantic("", "LineComment", uast.Comment{}, nil,
		Obj{
			"text": CommentText([2]string{"//", ""}, "comm"),
		},
		CommentNode(false, "comm", nil),
	),
	MapSemantic("", "Block", uast.Block{}, nil,
		Obj{
			"statements": Var("stmts"),
		},
		Obj{
			"Stmts": Var("stmts"),
		},
	),
	MapSemantic("", "ImportDeclaration", uast.Import{}, nil,
		Obj{
			"name":     Var("name"),
			"onDemand": String("true"),
			"static":   Var("static"),
		},
		Obj{
			"Path":  Var("name"),
			"All":   Bool(true),
			"Names": Arr(),
			// TODO: handle static when we have scopes
			"Scope": Obj{"static": Var("static")},
		},
	),
	MapSemantic("", "ImportDeclaration", uast.Import{}, nil,
		Obj{
			"name": Part("path", UASTType(uast.QualifiedIdentifier{}, Obj{
				"Names": Append(Var("names"), Arr(Var("name"))),
			})),
			"onDemand": String("false"),
			"static":   Var("static"),
		},
		Obj{
			"Path": Part("path", UASTType(uast.QualifiedIdentifier{}, Obj{
				"Names": Var("names"),
			})),
			"All":   Bool(false),
			"Names": Arr(Var("name")),
			// TODO: handle static when we have scopes
			"Scope": Obj{"static": Var("static")},
		},
	),

	MapSemantic("", "MethodDeclaration", uast.FunctionGroup{}, nil,
		Obj{
			"constructor":          Var("constr"),
			"extraDimensions2":     Is(nil), // TODO: find an example
			"javadoc":              Var("doc"),
			"modifiers":            Var("ann"), // TODO: it's an array, we should expand it somewhere
			"name":                 Var("name"),
			"body":                 Var("body"),
			"parameters":           Var("args"),
			"receiverQualifier":    Is(nil), // FIXME: handle receiver
			"receiverType":         Is(nil),
			"returnType2":          Var("out1"),
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
					"Obj": UASTType(uast.Function{}, Obj{
						"Type": UASTType(uast.FunctionType{}, Obj{
							"Args": Var("args"),
							"Returns": Arr(
								UASTType(uast.Argument{}, Obj{
									// TODO: can be void, so should be removed in this case
									"Type": Var("out1"),
								}),
							),
						}),
						"Body": Var("body"),
						"Recv": Is(nil),
					}),
				}),
				Obj{ // FIXME: store them as annotations at least
					"constructor":          Var("constr"),
					"thrownExceptionTypes": Var("exc"),
					"typeParameters":       Var("tmpl"),
				},
			),
		},
	),
}
