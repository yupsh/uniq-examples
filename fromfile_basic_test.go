package uniq_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/uniq"
)

// This example demonstrates reading from a file instead of strings.NewReader
func ExampleUniq_fromFile_basic() {
	// uniq testdata/duplicates.txt
	gloo.MustRun(
		Uniq(gloo.File("testdata/duplicates.txt")),
	)
	// Output:
	// apple
	// banana
	// apple
}
