package perl_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/perl"
)

func ExamplePerl_basic() {
	// echo "hello world" | perl -pe 's/world/universe/'
	gloo.MustRun(
		Perl(PrintLoop, Code("s/world/universe/"), strings.NewReader("hello world")),
	)
	// Output:
	// hello universe
}
