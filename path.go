package path

import "path"

type Path string

// Base runs [path.Base] on p
func (p Path) Base() string {
	return path.Base(string(p))
}

// Clean runs [path.Clean] on p
func (p Path) Clean() Path {
	return Path(path.Clean(string(p)))
}

// Dir runs [path.Dir] on p
func (p Path) Dir() Path {
	return Path(path.Dir(string(p)))
}

// Ext runs [path.Ext] on p
func (p Path) Ext() string {
	return path.Ext(string(p))
}

// IsAbs runs [path.IsAbs] on p
func (p Path) IsAbs() bool {
	return path.IsAbs(string(p))
}

// Join runs [path.Join] on []string{p, ...elem}
func (p Path) Join(elem ...Path) Path {
	elemPlain := make([]string, 1, len(elem)+1)
	elemPlain[0] = string(p)
	for _, e := range elem {
		elemPlain = append(elemPlain, string(e))
	}
	return Path(path.Join(elemPlain...))
}

// Split runs [path.Split] on p
func (p Path) Split() (dir Path, file string) {
	dirPlain, file := path.Split(string(p))
	return Path(dirPlain), file
}
