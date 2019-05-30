package fixtures

import (
	"path/filepath"
	"testing"

	"github.com/bblfsh/java-driver/driver/normalizer"
	"github.com/bblfsh/sdk/v3/driver"
	"github.com/bblfsh/sdk/v3/driver/fixtures"
	"github.com/bblfsh/sdk/v3/driver/native"
	"github.com/bblfsh/sdk/v3/uast/transformer/positioner"
)

const projectRoot = "../../"

var Suite = &fixtures.Suite{
	Lang: "java",
	Ext:  ".java",
	Path: filepath.Join(projectRoot, fixtures.Dir),
	NewDriver: func() driver.Native {
		return native.NewDriverAt(filepath.Join(projectRoot, "build/bin/native"), native.UTF8)
	},
	Transforms: normalizer.Transforms,
	BenchName:  "wildcard_type", // TODO: find a really large java file
	Semantic: fixtures.SemanticConfig{
		BlacklistTypes: []string{
			"StringLiteral",
			"SimpleName",
			"QualifiedName",
			"BlockComment",
			"LineComment",
			"Block",
			"ImportDeclaration",
			"MethodDeclaration",
		},
	},
	VerifyTokens: []positioner.VerifyToken{
		{Types: []string{
			"SimpleName",
			"StringLiteral",
		}},
	},
}

func TestJavaDriver(t *testing.T) {
	Suite.RunTests(t)
}

func BenchmarkJavaDriver(b *testing.B) {
	Suite.RunBenchmarks(b)
}
