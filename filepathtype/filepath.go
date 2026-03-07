package filepathtype

import (
	"io/fs"
	"path/filepath"
)

type FilePath string

func (p FilePath) Abs() (FilePath, error) {
	absPlain, err := filepath.Abs(string(p))
	return FilePath(absPlain), err
}

func (p FilePath) Base() string {
	return filepath.Base(string(p))
}

func (p FilePath) Clean() FilePath {
	return FilePath(filepath.Clean(string(p)))
}

func (p FilePath) Dir() FilePath {
	return FilePath(filepath.Dir(string(p)))
}

func (p FilePath) EvalSymlinks() (FilePath, error) {
	evalPlain, err := filepath.EvalSymlinks(string(p))
	return FilePath(evalPlain), err
}

func (p FilePath) Ext() string {
	return filepath.Ext(string(p))
}

func (p FilePath) IsAbs() bool {
	return filepath.IsAbs(string(p))
}

func (p FilePath) IsLocal() bool {
	return filepath.IsLocal(string(p))
}

func (p FilePath) Join(elem ...FilePath) FilePath {
	elemPlain := make([]string, 1, len(elem)+1)
	elemPlain[0] = string(p)
	for _, e := range elem {
		elemPlain = append(elemPlain, string(e))
	}
	return FilePath(filepath.Join(elemPlain...))
}

func (p FilePath) Rel(targPath FilePath) (FilePath, error) {
	relPlain, err := filepath.Rel(string(p), string(targPath))
	return FilePath(relPlain), err
}

func (p FilePath) Split() (dir FilePath, file string) {
	dirPlain, file := filepath.Split(string(p))
	return FilePath(dirPlain), file
}

func (p FilePath) SplitList() []string {
	return filepath.SplitList(string(p))
}

func (p FilePath) VolumeName() string {
	return filepath.VolumeName(string(p))
}

func FromSlash(path string) FilePath {
	return FilePath(filepath.FromSlash(path))
}

func (p FilePath) ToSlash() string {
	return filepath.ToSlash(string(p))
}

func (p FilePath) WalkDir(fn WalkDirFunc) error {
	return filepath.WalkDir(string(p), func(path string, d fs.DirEntry, err error) error {
		return fn(FilePath(path), d, err)
	})
}

type WalkDirFunc func(path FilePath, d fs.DirEntry, err error) error
