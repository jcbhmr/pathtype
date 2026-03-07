package pathtype

import (
	"path"
)

type Abs struct {
	inner Path
}

var _ interface {
	Bytes() []byte
} = Abs{}

func (p Abs) assertNotZero() {
	if p == (Abs{}) {
		panic("zero value of Abs is not valid")
	}
}

func AbsFrom(p Path) (Abs, bool) {
	if p.IsAbs() {
		return Abs{inner: p}, true
	} else {
		return Abs{}, false
	}
}

func (p Abs) Path() Path {
	p.assertNotZero()
	return p.inner
}

func (p Abs) Bytes() []byte {
	p.assertNotZero()
	return p.Path().Bytes()
}

func (p Abs) Base() string {
	p.assertNotZero()
	return p.Path().Base()
}

func (p Abs) Clean() Abs {
	p.assertNotZero()
	return Abs{inner: p.Path().Clean()}
}

func (p Abs) Dir() Abs {
	p.assertNotZero()
	return Abs{inner: p.Path().Dir()}
}

func (p Abs) Ext() string {
	p.assertNotZero()
	return p.Path().Ext()
}

func (p Abs) Join(elem ...Path) Abs {
	p.assertNotZero()
	elemPlain := make([]string, 1, len(elem)+1)
	elemPlain[0] = p.Path().inner
	for _, e := range elem {
		elemPlain = append(elemPlain, e.inner)
	}
	return Abs{inner: Path{inner: path.Join(elemPlain...)}}
}

func (p Abs) Split() (dir Abs, file string) {
	p.assertNotZero()
	dirPlain, file := path.Split(p.Path().inner)
	return Abs{inner: Path{inner: dirPlain}}, file
}
