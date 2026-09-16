// Package alias provides an unprefixed name for the dirname command.
//
//	import dirname "github.com/gloo-foo/cmd-dirname/alias"
//	dirname.Dirname()
package alias

import (
	gloo "github.com/gloo-foo/framework"

	command "github.com/gloo-foo/cmd-dirname"
)

// Dirname strips the final component from each input path; see the command
// package for the semantics.
func Dirname(opts ...any) gloo.Command[[]byte, []byte] { return command.Dirname(opts...) }
