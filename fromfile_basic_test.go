package perl_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/perl"
)

// This example demonstrates reading from a file instead of strings.NewReader
func ExamplePerl_fromFile_basic() {
	// cat testdata/text.txt | perl -pe 's/world/universe/'
	gloo.MustRun(
		Perl(PrintLoop, Code("s/world/universe/"), gloo.File("testdata/text.txt")),
	)
	// Output:
	// hello universe
}
