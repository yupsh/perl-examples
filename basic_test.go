package perl_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/perl"
)

func ExamplePerl_basic() {
	// echo "hello world" | perl -pe 's/world/universe/'
	yup.MustRun(
		Perl(PrintLoop, Code("s/world/universe/"), strings.NewReader("hello world")),
	)
	// Output:
	// hello universe
}

