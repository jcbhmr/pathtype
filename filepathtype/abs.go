package filepathtype

import (
	"path/filepath"
	"runtime"
)

type Abs struct {
	inner FilePath
}

func AbsFrom(p FilePath) (Abs, bool) {
	if p.IsAbs() {
		return Abs{inner: p}, true
	} else {
		return Abs{}, false
	}
}

func (p Abs) FilePath() FilePath {
	var zero Abs
	if p == zero {
		if runtime.GOOS == "windows" {
			return FilePath("C:\\")
		} else {
			return FilePath("/")
		}
	} else {
		return p.inner
	}
}

func (p Abs) String() string {
	return string(p.FilePath())
}

func (p Abs) Bytes() []byte {
	return []byte(p.FilePath())
}

func (p Abs) Base() string {
	return p.FilePath().Base()
}

func (p Abs) Clean() Abs {
	return Abs{inner: p.FilePath().Clean()}
}

func (p Abs) Dir() Abs {
	return Abs{inner: p.FilePath().Dir()}
}

func (p Abs) EvalSymlinks() (Abs, error) {
	evalPlain, err := p.FilePath().EvalSymlinks()
	return Abs{inner: evalPlain}, err
}

func (p Abs) Ext() string {
	return p.FilePath().Ext()
}

func (p Abs) Join(elem ...FilePath) Abs {
	elemPlain := make([]string, 1, len(elem)+1)
	elemPlain[0] = string(p.FilePath())
	for _, e := range elem {
		elemPlain = append(elemPlain, string(e))
	}
	return Abs{inner: FilePath(filepath.Join(elemPlain...))}
}

func (p Abs) Rel(targPath FilePath) (FilePath, error) {
	relPlain, err := filepath.Rel(string(p.FilePath()), string(targPath))
	return FilePath(relPlain), err
}

func (p Abs) Split() (dir Abs, file string) {
	dirPlain, file := filepath.Split(string(p.FilePath()))
	return Abs{inner: FilePath(dirPlain)}, file
}

func (p Abs) SplitList() []string {
	return filepath.SplitList(string(p.FilePath()))
}

func (p Abs) VolumeName() string {
	return filepath.VolumeName(string(p.FilePath()))
}
