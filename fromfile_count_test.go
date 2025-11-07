package uniq_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/uniq"
)

func ExampleUniq_fromFile_count() {
	// uniq -c testdata/duplicates.txt
	yup.MustRun(
		Uniq(Count, yup.File("testdata/duplicates.txt")),
	)
	// Output:
	//       2 apple
	//       2 banana
	//       1 apple
}

