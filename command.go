package command

import (
	gloo "github.com/gloo-foo/framework"
	"github.com/gloo-foo/framework/patterns"
)

const (
	// pathSeparator is the byte GNU dirname treats as a path component delimiter.
	pathSeparator = '/'
	// currentDir is the result for a path with no directory part (e.g. "file"),
	// an empty path, or ".".
	currentDir = '.'
)

// Dirname returns a Command that strips the final component from each input
// path, equivalent to the Unix dirname(1) utility.
//
// It matches GNU dirname's pure string semantics: trailing separators are
// removed, the last component is dropped, and remaining trailing separators are
// removed. Interior separators are preserved and "." / ".." are not resolved, so
// "a//b//c" yields "a//b" and "/foo/.." yields "/foo".
func Dirname(_ ...any) gloo.Command[[]byte, []byte] {
	return patterns.Map(func(line []byte) ([]byte, error) {
		return dir(line), nil
	})
}

// dir computes the GNU dirname of path as a pipeline of string transforms:
// drop trailing separators, drop the last component, drop trailing separators.
func dir(path []byte) []byte {
	if len(path) == 0 {
		return []byte{currentDir}
	}
	trimmed := trimTrailingSeparators(path)
	if len(trimmed) == 0 {
		return []byte{pathSeparator}
	}
	parent := trimLastComponent(trimmed)
	if len(parent) == 0 {
		return []byte{currentDir}
	}
	cleaned := trimTrailingSeparators(parent)
	if len(cleaned) == 0 {
		return []byte{pathSeparator}
	}
	return cleaned
}

// trimTrailingSeparators returns path without any trailing separator bytes.
func trimTrailingSeparators(path []byte) []byte {
	end := len(path)
	for end > 0 && path[end-1] == pathSeparator {
		end--
	}
	return path[:end]
}

// trimLastComponent returns path without its final separator-free component,
// keeping any separators that precede that component.
func trimLastComponent(path []byte) []byte {
	end := len(path)
	for end > 0 && path[end-1] != pathSeparator {
		end--
	}
	return path[:end]
}
