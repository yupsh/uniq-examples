package uniq_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/uniq"
)

func ExampleUniq_basic() {
	// echo "apple\napple\nbanana\nbanana\napple" | uniq
	yup.MustRun(
		Uniq(strings.NewReader("apple\napple\nbanana\nbanana\napple")),
	)
	// Output:
	// apple
	// banana
	// apple
}

