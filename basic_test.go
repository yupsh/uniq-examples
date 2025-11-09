package uniq_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/uniq"
)

func ExampleUniq_basic() {
	// echo "apple\napple\nbanana\nbanana\napple" | uniq
	gloo.MustRun(
		Uniq(strings.NewReader("apple\napple\nbanana\nbanana\napple")),
	)
	// Output:
	// apple
	// banana
	// apple
}
