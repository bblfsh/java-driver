package fixtures

import (
	"path/filepath"
	"testing"

	"github.com/bblfsh/java-driver/driver/normalizer"
	"gopkg.in/bblfsh/sdk.v2/sdk/driver"
	"gopkg.in/bblfsh/sdk.v2/sdk/driver/fixtures"
)

const projectRoot = "../../"

var Suite = &fixtures.Suite{
	Lang: "java",
	Ext:  ".java",
	Path: filepath.Join(projectRoot, fixtures.Dir),
	NewDriver: func() driver.BaseDriver {
		return driver.NewExecDriverAt(filepath.Join(projectRoot, "build/bin/native"))
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
	Docker: fixtures.DockerConfig{
		Image: "openjdk:8",
	},
}

func TestJavaDriver(t *testing.T) {
	Suite.RunTests(t)
}

func BenchmarkJavaDriver(b *testing.B) {
	Suite.RunBenchmarks(b)
}
