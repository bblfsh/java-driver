package normalizer

import (
	"github.com/bblfsh/sdk/protocol/driver"
	"github.com/bblfsh/sdk/protocol/native"
)

var ToNoder = &native.ObjectToNoder{
	InternalTypeKey: "internalClass",
	LineKey:         "line",
	OffsetKey:       "startPosition",
	//TODO: Should this be part of the UAST rules?
	TokenKeys: map[string]bool{
		"identifier":        true, // SimpleName
		"escapedValue":      true, // StringLiteral
		"keyword":           true, // Modifier
		"primitiveTypeCode": true, // ?
	},
	SyntheticTokens: map[string]string{
		"PackageDeclaration": "package",
		"IfStatement":        "if",
		"NullLiteral":        "null",
	},
	//TODO: add names of children (e.g. elseStatement) as
	//      children node properties.
}

func transformationParser(opts driver.UASTParserOptions) (tr driver.UASTParser, err error) {
	parser, err := native.ExecParser(ToNoder, opts.NativeBin)
	if err != nil {
		return tr, err
	}

	tr = &driver.TransformationUASTParser{
		UASTParser: parser,
		Transformation: driver.FillLineColFromOffset,
	}

	return tr, nil
}

// UASTParserBuilder creates a parser that transform source code files into *uast.Node.
func UASTParserBuilder(opts driver.UASTParserOptions) (driver.UASTParser, error) {
	parser, err := transformationParser(opts)
	if err != nil {
		return nil, err
	}

	return parser, nil
}
