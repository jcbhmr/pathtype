package path_test

import (
	"fmt"
	"iter"
	"testing"

	path "go.jcbhmr.com/pathtype"
)

func TestPattern_isStringlike(t *testing.T) {
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

	var zero path.Pattern
	var _ Stringlike[path.Pattern] = zero
}

func ExamplePattern() {
	pattern := path.MustNewPattern("a*/?")

	pathValue := path.Path("abcd/e")
	result := pattern.Match(pathValue)
	fmt.Printf("%q matched %q is %t", pathValue, pattern, result)

	// Output: "abcd/e" matched "a*/?" is true
}
