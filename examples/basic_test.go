package dirname_test

import (
	"fmt"

	"github.com/gloo-foo/testable"

	command "github.com/gloo-foo/cmd-dirname"
)

func ExampleDirname_basic() {
	// echo "/path/to/file.txt" | dirname
	output, _ := testable.Test(command.Dirname(), "/path/to/file.txt\n")
	fmt.Print(output)
	// Output:
	// /path/to
}
