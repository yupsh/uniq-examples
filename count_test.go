package uniq_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/uniq"
)

func ExampleUniq_count() {
	// echo "apple\napple\nbanana\napple" | uniq -c
	gloo.MustRun(
		Uniq(Count, strings.NewReader("apple\napple\nbanana\napple")),
	)
	// Output:
	//       2 apple
	//       1 banana
	//       1 apple
}
