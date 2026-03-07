package filepath_test

import (
	"fmt"
	"iter"
	"testing"

	"go.jcbhmr.com/pathtype/filepath"
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

	var zero filepath.Pattern
	var _ Stringlike[filepath.Pattern] = zero
}
