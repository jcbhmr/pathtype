package filepathtype

import (
	"path/filepath"
	"unsafe"
)

type Pattern struct {
	inner string
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
