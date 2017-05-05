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

// ASTParserBuilder creates a parser that transform source code files into *uast.Node.
func ASTParserBuilder(opts driver.ASTParserOptions) (driver.ASTParser, error) {
	parser, err := native.ExecParser(ToNoder, opts.NativeBin)
	if err != nil {
		return nil, err
	}

	return parser, nil
}
