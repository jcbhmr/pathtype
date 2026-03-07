package filepathtype

import (
	"path/filepath"
)

type FilePath struct {
	inner string
}

func From(s string) FilePath {
	return FilePath{inner: s}
}

func (p FilePath) Bytes() []byte {
	return []byte(p.inner)
}

func (p FilePath) Abs() (Abs, error) {
	absPlain, err := filepath.Abs(p.inner)
	if err != nil {
		return Abs{}, err
	}
	return Abs{FilePath{inner: absPlain}}, nil
}

func (p FilePath) Base() string {
	return filepath.Base(p.inner)
}

func (p FilePath) Clean() FilePath {
	return FilePath{inner: filepath.Clean(p.inner)}
}

func (p FilePath) Dir() FilePath {
	return FilePath{inner: filepath.Dir(p.inner)}
}

func (p FilePath) EvalSymlinks() (FilePath, error) {
	evalPlain, err := filepath.EvalSymlinks(p.inner)
	return FilePath{inner: evalPlain}, err
}

func (p FilePath) Ext() string {
	return filepath.Ext(p.inner)
}

func (p FilePath) IsAbs() bool {
	return filepath.IsAbs(p.inner)
}

func (p FilePath) IsLocal() bool {
	return filepath.IsLocal(p.inner)
}

func (p FilePath) Join(elem ...FilePath) FilePath {
	elemPlain := make([]string, 1, len(elem)+1)
	elemPlain[0] = p.inner
	for _, e := range elem {
		elemPlain = append(elemPlain, e.inner)
	}
	return FilePath{inner: filepath.Join(elemPlain...)}
}

func (p FilePath) Rel(targPath FilePath) (FilePath, error) {
	relPlain, err := filepath.Rel(p.inner, targPath.inner)
	return FilePath{inner: relPlain}, err
}

func (p FilePath) Split() (dir FilePath, file string) {
	dirPlain, file := filepath.Split(p.inner)
	return FilePath{inner: dirPlain}, file
}

func (p FilePath) SplitList() []string {
	return filepath.SplitList(p.inner)
}

func (p FilePath) VolumeName() string {
	return filepath.VolumeName(p.inner)
}
