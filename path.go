package pathtype

import "path"

type Path struct {
	inner string
}

func From(s string) Path {
	return Path{inner: s}
}

func (p Path) Bytes() []byte {
	return []byte(p.inner)
}

func (p Path) Base() string {
	return path.Base(p.inner)
}

func (p Path) Clean() Path {
	return Path{inner: path.Clean(p.inner)}
}

func (p Path) Dir() Path {
	return Path{inner: path.Dir(p.inner)}
}

func (p Path) Ext() string {
	return path.Ext(p.inner)
}

func (p Path) IsAbs() bool {
	return path.IsAbs(p.inner)
}

func (p Path) Join(elem ...Path) Path {
	elemPlain := make([]string, 1, len(elem)+1)
	elemPlain[0] = p.inner
	for _, e := range elem {
		elemPlain = append(elemPlain, e.inner)
	}
	return Path{inner: path.Join(elemPlain...)}
}

func (p Path) Split() (dir Path, file string) {
	dirPlain, file := path.Split(p.inner)
	return Path{inner: dirPlain}, file
}
