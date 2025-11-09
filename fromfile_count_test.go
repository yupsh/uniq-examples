package uniq_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/uniq"
)

func ExampleUniq_fromFile_count() {
	// uniq -c testdata/duplicates.txt
	gloo.MustRun(
		Uniq(Count, gloo.File("testdata/duplicates.txt")),
	)
	// Output:
	//       2 apple
	//       2 banana
	//       1 apple
}
