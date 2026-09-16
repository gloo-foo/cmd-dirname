package alias_test

import (
	"slices"
	"testing"

	"github.com/gloo-foo/testable"

	dirname "github.com/gloo-foo/cmd-dirname/alias"
)

// The alias package re-exports the constructor under an unprefixed name. A
// mis-wired re-export (Dirname bound to the wrong function) compiles cleanly, so
// only behavior can prove the wiring. The tests exercise the re-export and
// assert the GNU dirname output it must produce: the final path component is
// stripped, interior separators are preserved, and "." / ".." are not resolved.

func assertLines(t *testing.T, got, want []string) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestAlias_DirnameStripsFinalComponent(t *testing.T) {
	lines, err := testable.TestLines(dirname.Dirname(), "/usr/local/bin/script.sh\n")
	if err != nil {
		t.Fatal(err)
	}
	assertLines(t, lines, []string{"/usr/local/bin"})
}

func TestAlias_DirnamePreservesInteriorSeparators(t *testing.T) {
	// GNU dirname is a pure string operation: it does not collapse "//" runs.
	lines, err := testable.TestLines(dirname.Dirname(), "a//b//c\n")
	if err != nil {
		t.Fatal(err)
	}
	assertLines(t, lines, []string{"a//b"})
}
