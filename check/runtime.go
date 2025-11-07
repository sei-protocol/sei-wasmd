// Package check provides utilities to inspect properties of the running Go
// binary, with a focus on detecting whether certain build-time flags were used
// during compilation.
//
// For example, using -trimpath can remove absolute file paths from the compiled
// executable, which affects stack traces and error reporting. This package helps
// detect that condition at runtime, which can be useful for debugging, logging,
// or reproducibility checks that otherwise can result in chain halt.
package check

import (
	"runtime"
	"strings"
)

// RuntimePathTrimmed reports whether the Go binary was likely built with
// the `-trimpath` flag by inspecting the file path of the current call stack.
//
// It checks if the file path of the caller frame contains the original module
// path ("github.com/CosmWasm/wasmd"). If the path contains this prefix, it
// likely means that `-trimpath` was used to remove or rewrite file system paths
// in the compiled executable.
//
// Note that building with `-trimpath` can remove absolute paths from stack
// traces, which may affect reproducibility of error reporting.
//
// Returns:
//   - trimmed: true if the file path indicates that the module path was trimmed.
//   - ok: true if runtime caller information could be retrieved successfully.
func RuntimePathTrimmed() (trimmed, ok bool) {
	switch _, file, _, success := runtime.Caller(0); {
	case !success:
		return false, false
	default:
		const modulePathPrefix = "github.com/CosmWasm/wasmd"
		return strings.HasPrefix(file, modulePathPrefix), true
	}
}
