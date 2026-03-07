package filepath_test

import (
	"fmt"
	"iter"
	"runtime"
	"testing"

	"go.jcbhmr.com/pathtype/filepath"
)

func TestFilePath_stringConversion(t *testing.T) {
	var pathAsString string = "./foo/bar/baz.txt"
	var pathAsFilePath filepath.FilePath = filepath.FilePath(pathAsString)
	_ = pathAsFilePath
}

func TestFilePath_notStringlike(t *testing.T) {
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

	var path filepath.FilePath
	if _, ok := any(path).(Stringlike[filepath.FilePath]); ok {
		t.Errorf("filepath.FilePath should not implement Stringlike, but it does")
	}
}

func TestFilePath_IsAbs(t *testing.T) {
	var absString string
	if runtime.GOOS == "windows" {
		absString = `C:\foo\bar\baz.txt`
	} else {
		absString = `/foo/bar/baz.txt`
	}
	var absPath filepath.FilePath = filepath.FilePath(absString)
	if want, got := true, absPath.IsAbs(); want != got {
		t.Errorf("want %#+v, got %#+v", want, got)
	}
}

func TestFilePath_Join(t *testing.T) {
	var path1, path2, path3 filepath.FilePath = "foo", "bar", "baz.txt"
	if want, got := filepath.FilePath("foo/bar/baz.txt"), path1.Join(path2, path3); want != got {
		t.Errorf("want %#+v, got %#+v", want, got)
	}
}
