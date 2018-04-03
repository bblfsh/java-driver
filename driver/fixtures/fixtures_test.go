package fixtures

import (
	"testing"
	"path/filepath"

	"github.com/bblfsh/go-driver/driver/normalizer"
	"gopkg.in/bblfsh/sdk.v1/sdk/driver"
	"gopkg.in/bblfsh/sdk.v1/sdk/driver/fixtures"
)

const projectRoot = "../../"

var Suite = &fixtures.Suite{
	Lang: "java",
	Ext:  ".java",
	Path: filepath.Join(projectRoot, fixtures.Dir),
	NewDriver: func() driver.BaseDriver {
		return driver.NewExecDriverAt(filepath.Join(projectRoot, "build/bin/native"))
	},
	Transforms: driver.Transforms{
		Native: normalizer.Native,
		Code:   normalizer.Code,
	},
	BenchName: "wildcard_type", // TODO: find a really large java file
}

func TestJavaDriver(t *testing.T) {
	Suite.RunTests(t)
}

func BenchmarkJavaDriver(b *testing.B) {
	Suite.RunBenchmarks(b)
}
