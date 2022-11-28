package filepackage

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

var mg map[string]FilesGneral

func GetMapDirs() map[string]FilesGneral {
	return mg
}

func ScanRootDir() (rootDir DirModel) {
	mg = make(map[string]FilesGneral)

	fmt.Println("ScanRootDir")
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	if err == nil {
		rootDir = DirModel{DirName: "root", NumfilesContent: len(files), DirRoot: "./"}
		strings.Join(strings.Split(strings.Join(strings.Split(rootDir.DirName, "/"), "-"), "."), "_")
		findInDirectory(files, &rootDir)
		// recorremos el contenido de la carpeta principal(desde donde se descargo el archivo)
		fmt.Println("\n\n\n\n\n\n")
		fmt.Println(rootDir)
	}
	return
}

func findInDirectory(filesR []fs.FileInfo, rootDir *DirModel) {
	for _, file := range filesR {
		// fmt.Println(file.Name(), file.IsDir())
		nameDir := filepath.Join(rootDir.DirRoot, file.Name())

		if file.IsDir() {

			color.Blue("name dir    " + nameDir)

			files, err := ioutil.ReadDir(nameDir)
			if err != nil {
				log.Fatal(err)
			}

			if err == nil {
				subRootDir := DirModel{
					DirRoot:         nameDir,
					DirName:         file.Name(),
					NumfilesContent: len(files),
				}
				subRootDir.GenKeyMap()
				findInDirectory(files, &subRootDir)
				mg[subRootDir.KeyMap] = &subRootDir
				rootDir.addDirM(&subRootDir)
			}

		} else {

			fileInDir := FileModel{FileName: file.Name(), Directory: nameDir}
			fileInDir.GenKeyMap()

			// m[fmt.Sprintf("%d%s", (rand.Intn(10_000-100)+100), "_dir")] = subRootDir
			mg[fileInDir.KeyMap] = &fileInDir
			rootDir.Files = append(rootDir.Files, fileInDir)
		}

	}
	fmt.Printf("%+v\n", rootDir)
}
