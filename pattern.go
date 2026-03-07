package pathtype

import "path"

// Pattern is
type Pattern struct {
	inner string
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
