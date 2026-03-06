package pathtype

import "path"

type Path string

func (p Path) Base() string {
	return path.Base(string(p))
}

func (p Path) Clean() Path {
	return Path(path.Clean(string(p)))
}

func (p Path) Dir() Path {
	return Path(path.Dir(string(p)))
}

func (p Path) Ext() string {
	return path.Ext(string(p))
}

func (p Path) IsAbs() bool {
	return path.IsAbs(string(p))
}

func (p Path) Join(elem ...Path) Path {
	elemPlain := make([]string, 1, len(elem)+1)
	elemPlain[0] = string(p)
	for _, e := range elem {
		elemPlain = append(elemPlain, string(e))
	}
	return Path(path.Join(elemPlain...))
}

func (p Path) Split() (dir Path, file string) {
	dirPlain, file := path.Split(string(p))
	return Path(dirPlain), file
}
