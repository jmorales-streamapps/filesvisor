package filepackage

import (
	"fmt"
	"math/rand"
)

type FilesGneral interface {
	GenKeyMap()
	GetDirectory() string
}

type DirModel struct {
	KeyMap          string
	DirRoot         string
	DirName         string
	NumfilesContent int
	Files           []FileModel
	DirsChils       []*DirModel
}

func (md *DirModel) addDirM(newDM *DirModel) {
	md.DirsChils = append(md.DirsChils, newDM)
}

type FileModel struct {
	KeyMap    string
	FileName  string
	Directory string
}

func (fm *FileModel) GenKeyMap() {
	fm.KeyMap = fmt.Sprintf("%d%s", (rand.Intn(10_000-100) + 100), "_file")
}

func (dm *DirModel) GenKeyMap() {
	dm.KeyMap = fmt.Sprintf("%d%s", (rand.Intn(10_000-100) + 100), "_dir")
}

func (fm *FileModel) GetDirectory() string {
	return fm.Directory
}

func (dm *DirModel) GetDirectory() string {
	return dm.DirRoot
}
