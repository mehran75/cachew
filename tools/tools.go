//go:build tools

package tools

// This file contains imports for tools. The tools would otherwise be considered unnecessary dependencies by the
// `go mod` tool
import (
	_ "gotest.tools/gotestsum/cmd"      // test output formatting command
	_ "gotest.tools/gotestsum/testjson" // test output formatting tool
	_ "honnef.co/go/tools/config"       // staticcheck config dependency
	_ "honnef.co/go/tools/staticcheck"  // static source code analysis tool
)
