package filepath

import (
	"iter"
	"path/filepath"
	"unsafe"
)

// Pattern is a wrapper around a validated [string] that conforms to the pattern syntax described by [filepath.Match].
type Pattern struct {
	inner string
}

func (p Pattern) String() string {
	return p.inner
}

func (p Pattern) Bytes() []byte {
	return []byte(p.inner)
}

func (p Pattern) Runes() []rune {
	return []rune(p.inner)
}

func (p Pattern) All() iter.Seq2[int, rune] {
	return func(yield func(int, rune) bool) {
		for i, r := range p.inner {
			if !yield(i, r) {
				return
			}
		}
	}
}

func (p Pattern) At(i int) byte {
	return p.inner[i]
}

func (p Pattern) Compare(o Pattern) int {
	if p.inner < o.inner {
		return -1
	} else if p.inner > o.inner {
		return 1
	} else {
		return 0
	}
}

func (p Pattern) Equal(o Pattern) bool {
	return p.inner == o.inner
}

func NewPattern(pattern string) (Pattern, error) {
	_, err := filepath.Match(pattern, "")
	if err != nil {
		return Pattern{}, err
	}
	return Pattern{pattern}, nil
}

func MustNewPattern(pattern string) Pattern {
	p, err := NewPattern(pattern)
	if err != nil {
		panic(err)
	}
	return p
}

func (p Pattern) Match(name FilePath) bool {
	matched, err := filepath.Match(p.inner, string(name))
	if err != nil {
		panic(err)
	}
	return matched
}

func (p Pattern) Glob() []FilePath {
	matchesPlain, err := filepath.Glob(p.inner)
	if err != nil {
		panic(err)
	}
	// SAFETY: FilePath is a newtype of string, so this conversion is safe.
	matches := *(*[]FilePath)(unsafe.Pointer(&matchesPlain))
	return matches
}
