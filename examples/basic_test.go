package dirname_test

import (
	"fmt"

	command "github.com/gloo-foo/cmd-dirname"
	"github.com/gloo-foo/testable"
)

func ExampleDirname_basic() {
	// echo "/path/to/file.txt" | dirname
	output, _ := testable.Test(command.Dirname(), "/path/to/file.txt\n")
	fmt.Print(output)
	// Output:
	// /path/to
}
