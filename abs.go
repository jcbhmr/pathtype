package pathtype

import "path"

type Abs struct {
	inner Path
}

func AbsFrom(p Path) (Abs, bool) {
	if p.IsAbs() {
		return Abs{inner: p}, true
	} else {
		return Abs{}, false
	}
}

func (p Abs) Path() Path {
	var zero Abs
	if p == zero {
		return Path("/")
	} else {
		return p.inner
	}
}

func (p Abs) String() string {
	return string(p.Path())
}

func (p Abs) Bytes() []byte {
	return []byte(p.Path())
}

func (p Abs) Base() string {
	return p.Path().Base()
}

func (p Abs) Clean() Abs {
	return Abs{inner: p.Path().Clean()}
}

func (p Abs) Dir() Abs {
	return Abs{inner: p.Path().Dir()}
}

func (p Abs) Ext() string {
	return p.Path().Ext()
}

func (p Abs) Join(elem ...Path) Abs {
	elemPlain := make([]string, 1, len(elem)+1)
	elemPlain[0] = string(p.Path())
	for _, e := range elem {
		elemPlain = append(elemPlain, string(e))
	}
	return Abs{inner: Path(path.Join(elemPlain...))}
}

func (p Abs) Split() (dir Abs, file string) {
	dirPlain, file := path.Split(string(p.Path()))
	return Abs{inner: Path(dirPlain)}, file
}
