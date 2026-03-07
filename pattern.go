package pathtype

import (
	"iter"
	"path"
)

// Pattern is a valid pattern for [path.Match]
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

func (p Pattern) RunesSeq() iter.Seq2[int, rune] {
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

// Compare compares p and o according to their inner strings.
//
//   - p < o: -1
//   - p == o: 0
//   - p > o: 1
func (p Pattern) Compare(o Pattern) int {
	if p.inner < o.inner {
		return -1
	} else if p.inner > o.inner {
		return 1
	} else {
		return 0
	}
}

func NewPattern(pattern string) (Pattern, error) {
	_, err := path.Match(pattern, "")
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

func (p Pattern) Match(name Path) bool {
	matched, err := path.Match(p.inner, string(name))
	if err != nil {
		panic(err)
	}
	return matched
}
