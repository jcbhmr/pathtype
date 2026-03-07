package filepathtype

import (
	"path/filepath"
)

type absInner = FilePath

type Abs struct {
	absInner
}

func (p Abs) assertNotZero() {
	if p == (Abs{}) {
		panic("zero value of Abs is not valid")
	}
}

func AbsFrom(p FilePath) (Abs, bool) {
	if p.IsAbs() {
		return Abs{p}, true
	} else {
		return Abs{}, false
	}
}

func (p Abs) FilePath() FilePath {
	p.assertNotZero()
	return p.absInner
}

func (p Abs) Bytes() []byte {
	p.assertNotZero()
	return p.FilePath().Bytes()
}

func (p Abs) Base() string {
	p.assertNotZero()
	return p.FilePath().Base()
}

func (p Abs) Clean() Abs {
	p.assertNotZero()
	return Abs{p.FilePath().Clean()}
}

func (p Abs) Dir() Abs {
	p.assertNotZero()
	return Abs{p.FilePath().Dir()}
}

func (p Abs) EvalSymlinks() (Abs, error) {
	p.assertNotZero()
	evalPlain, err := p.FilePath().EvalSymlinks()
	return Abs{evalPlain}, err
}

func (p Abs) Ext() string {
	p.assertNotZero()
	return p.FilePath().Ext()
}

func (p Abs) Join(elem ...FilePath) Abs {
	p.assertNotZero()
	elemPlain := make([]string, 1, len(elem)+1)
	elemPlain[0] = p.FilePath().inner
	for _, e := range elem {
		elemPlain = append(elemPlain, e.inner)
	}
	return Abs{FilePath{inner: filepath.Join(elemPlain...)}}
}

func (p Abs) Rel(targPath FilePath) (FilePath, error) {
	p.assertNotZero()
	relPlain, err := filepath.Rel(p.FilePath().inner, targPath.inner)
	return FilePath{inner: relPlain}, err
}

func (p Abs) Split() (dir Abs, file string) {
	p.assertNotZero()
	dirPlain, file := filepath.Split(p.FilePath().inner)
	return Abs{FilePath{inner: dirPlain}}, file
}
