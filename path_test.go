package path_test

import (
	"fmt"
	"iter"
	"testing"

	path "go.jcbhmr.com/pathtype"
)

func TestPath_stringConversion(t *testing.T) {
	var pathAsString string = "./foo/bar/baz.txt"
	var pathAsPath path.Path = path.Path(pathAsString)
	_ = pathAsPath
}

func TestPath_notStringlike(t *testing.T) {
	type Stringlike[T any] interface {
		// string(v)
		fmt.Stringer
		// []byte(v)
		Bytes() []byte
		// []rune(v)
		Runes() []rune
		// for i, r := range v
		All() iter.Seq2[int, rune]
		// v[i]
		At(int) byte
		// v < w, v == w, v > w
		Compare(T) int
		// v == w, v != w
		Equal(T) bool
	}

	var p path.Path
	if _, ok := any(p).(Stringlike[path.Path]); ok {
		t.Errorf("filepath.FilePath should not implement Stringlike, but it does")
	}
}

func TestPath_IsAbs(t *testing.T) {
	var absPath path.Path = path.Path(`/foo/bar/baz.txt`)
	if want, got := true, absPath.IsAbs(); want != got {
		t.Errorf("want %#+v, got %#+v", want, got)
	}
}

func TestPath_Join(t *testing.T) {
	var path1, path2, path3 path.Path = "foo", "bar", "baz.txt"
	if want, got := path.Path("foo/bar/baz.txt"), path1.Join(path2, path3); want != got {
		t.Errorf("want %#+v, got %#+v", want, got)
	}
}
