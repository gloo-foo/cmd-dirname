package command_test

import (
	"fmt"
	"testing"

	"github.com/gloo-foo/testable"

	command "github.com/gloo-foo/cmd-dirname"
)

func TestDirname_AbsolutePath(t *testing.T) {
	lines, err := testable.TestLines(command.Dirname(), "/usr/local/bin/script.sh\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 || lines[0] != "/usr/local/bin" {
		t.Fatalf("got %q, want [/usr/local/bin]", lines)
	}
}

func TestDirname_RelativePath(t *testing.T) {
	lines, err := testable.TestLines(command.Dirname(), "relative/path/file.txt\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 || lines[0] != "relative/path" {
		t.Fatalf("got %q, want [relative/path]", lines)
	}
}

func TestDirname_FilenameOnly(t *testing.T) {
	lines, err := testable.TestLines(command.Dirname(), "filename\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 || lines[0] != "." {
		t.Fatalf("got %q, want [.]", lines)
	}
}

func TestDirname_Root(t *testing.T) {
	lines, err := testable.TestLines(command.Dirname(), "/\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 || lines[0] != "/" {
		t.Fatalf("got %q, want [/]", lines)
	}
}

func TestDirname_Dot(t *testing.T) {
	lines, err := testable.TestLines(command.Dirname(), ".\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 || lines[0] != "." {
		t.Fatalf("got %q, want [.]", lines)
	}
}

func TestDirname_DotDot(t *testing.T) {
	lines, err := testable.TestLines(command.Dirname(), "..\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 || lines[0] != "." {
		t.Fatalf("got %q, want [.]", lines)
	}
}

func TestDirname_TrailingSlash(t *testing.T) {
	lines, err := testable.TestLines(command.Dirname(), "/path/to/dir/\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 || lines[0] != "/path/to" {
		t.Fatalf("got %q, want [/path/to]", lines)
	}
}

func TestDirname_MultipleLines(t *testing.T) {
	lines, err := testable.TestLines(command.Dirname(), "/usr/bin/ls\n/etc/nginx/nginx.conf\n/var/log/app.log\n")
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"/usr/bin", "/etc/nginx", "/var/log"}
	if len(lines) != len(want) {
		t.Fatalf("got %d lines, want %d", len(lines), len(want))
	}
	for i, w := range want {
		if lines[i] != w {
			t.Errorf("line %d: got %q, want %q", i, lines[i], w)
		}
	}
}

func TestDirname_EmptyInput(t *testing.T) {
	lines, err := testable.TestLines(command.Dirname(), "")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 0 {
		t.Fatalf("got %q, want empty", lines)
	}
}

func TestDirname_DeepPath(t *testing.T) {
	lines, err := testable.TestLines(command.Dirname(), "/a/b/c/d/e/f/g/file.txt\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 || lines[0] != "/a/b/c/d/e/f/g" {
		t.Fatalf("got %q, want [/a/b/c/d/e/f/g]", lines)
	}
}

func TestDirname_HiddenFile(t *testing.T) {
	lines, err := testable.TestLines(command.Dirname(), "/home/user/.bashrc\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 || lines[0] != "/home/user" {
		t.Fatalf("got %q, want [/home/user]", lines)
	}
}

func TestDirname_Spaces(t *testing.T) {
	lines, err := testable.TestLines(command.Dirname(), "/path/to/file with spaces.txt\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 || lines[0] != "/path/to" {
		t.Fatalf("got %q, want [/path/to]", lines)
	}
}

func TestDirname_SpecialChars(t *testing.T) {
	lines, err := testable.TestLines(command.Dirname(), "/path/to/file-name_v2.txt\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 || lines[0] != "/path/to" {
		t.Fatalf("got %q, want [/path/to]", lines)
	}
}

func TestDirname_Unicode(t *testing.T) {
	lines, err := testable.TestLines(command.Dirname(), "/path/to/\u30d5\u30a1\u30a4\u30eb.txt\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 || lines[0] != "/path/to" {
		t.Fatalf("got %q, want [/path/to]", lines)
	}
}

func TestDirname_UnicodeDir(t *testing.T) {
	lines, err := testable.TestLines(command.Dirname(), "/\u8def\u5f84/\u6587\u4ef6.txt\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 || lines[0] != "/\u8def\u5f84" {
		t.Fatalf("got %q, want [/\u8def\u5f84]", lines)
	}
}

func TestDirname_PreservesInteriorSlashes(t *testing.T) {
	// GNU dirname is a pure string operation and does not collapse "//" runs the
	// way path canonicalization would: "a//b//c" yields "a//b", not "a/b".
	lines, err := testable.TestLines(command.Dirname(), "a//b//c\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 || lines[0] != "a//b" {
		t.Fatalf("got %q, want [a//b]", lines)
	}
}

func TestDirname_DoesNotResolveDotComponent(t *testing.T) {
	// GNU dirname does not resolve ".": "foo/." strips the "." component to "foo",
	// whereas canonicalization would collapse the whole path to ".".
	lines, err := testable.TestLines(command.Dirname(), "foo/.\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 || lines[0] != "foo" {
		t.Fatalf("got %q, want [foo]", lines)
	}
}

func TestDirname_DoesNotResolveParentComponent(t *testing.T) {
	// GNU dirname does not resolve "..": "/foo/.." strips the ".." component to
	// "/foo", whereas canonicalization would collapse the path to "/".
	lines, err := testable.TestLines(command.Dirname(), "/foo/..\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 || lines[0] != "/foo" {
		t.Fatalf("got %q, want [/foo]", lines)
	}
}

func TestDirname_EmptyPath(t *testing.T) {
	// A blank input line is an empty path; GNU dirname maps the empty string to
	// ".". The leading and trailing components keep the empty record in the
	// middle so it survives to the command.
	lines, err := testable.TestLines(command.Dirname(), "a/b\n\nc/d\n")
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"a", ".", "c"}
	if len(lines) != len(want) {
		t.Fatalf("got %d lines, want %d: %q", len(lines), len(want), lines)
	}
	for i, w := range want {
		if lines[i] != w {
			t.Errorf("line %d: got %q, want %q", i, lines[i], w)
		}
	}
}

func TestDirname_AllSeparators(t *testing.T) {
	// A path made only of separators has no component to strip; GNU dirname
	// returns a single separator.
	lines, err := testable.TestLines(command.Dirname(), "//\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 || lines[0] != "/" {
		t.Fatalf("got %q, want [/]", lines)
	}
}

func TestDirname_TrailingSeparatorsThenRoot(t *testing.T) {
	// "/usr/" strips trailing separators to "/usr", drops the "usr" component to
	// "/", which is the GNU result.
	lines, err := testable.TestLines(command.Dirname(), "/usr/\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 || lines[0] != "/" {
		t.Fatalf("got %q, want [/]", lines)
	}
}

func ExampleDirname() {
	lines, _ := testable.TestLines(command.Dirname(), "/usr/local/bin/script.sh\n/var/log/app.log\n")
	for _, line := range lines {
		fmt.Println(line)
	}
	// Output:
	// /usr/local/bin
	// /var/log
}

func ExampleDirname_filenameOnly() {
	lines, _ := testable.TestLines(command.Dirname(), "filename\n")
	for _, line := range lines {
		fmt.Println(line)
	}
	// Output:
	// .
}
