package handlefiles

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Jon-MC-dev/files_copy/functions"
)

var MapScannedFiles map[string]DirectoryNode

type DirectoryNode struct {
	Name      string          `json:"name"`
	IsDir     bool            `json:"isDir"`
	Files     []DirectoryNode `json:"children,omitempty"`
	Reference string
}

func ReeadDirectory() *DirectoryNode {
	// Inicializa el map
	MapScannedFiles = make(map[string]DirectoryNode)

	// Ruta del directorio raíz que deseas explorar
	rootDir := "./"

	// Llama a la función para construir el árbol
	newReference := functions.GenString(10)
	var mainNode DirectoryNode = DirectoryNode{Name: rootDir, IsDir: true, Reference: newReference}
	MapScannedFiles[newReference] = mainNode

	tree, err := buildTree(rootDir)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	mainNode.Files = append(mainNode.Files, tree...)

	// Convierte la estructura del árbol a formato JSON
	// jsonData, err := json.MarshalIndent(mainNode, "", "  ")
	// if err != nil {
	// 	fmt.Println("Error al convertir a JSON:", err)
	//	return nil
	//}

	// Imprime el JSON resultante
	//fmt.Println(string(jsonData))

	return &mainNode

}

func buildTree(rootPath string) ([]DirectoryNode, error) {

	// return DirectoryNode{}, nil

	// Lee el contenido del directorio
	files, err := os.ReadDir(rootPath)
	if err != nil {
		return []DirectoryNode{}, err
	}
	//   var child DirectoryNode
	// Recorre los archivos/directorios y construye el árbol
	files_and_dirs := []DirectoryNode{}

	for _, file := range files {
		childPath := filepath.Join(rootPath, file.Name())
		newReference := functions.GenString(10)

		child := DirectoryNode{
			Name:      file.Name(),
			IsDir:     file.IsDir(),
			Reference: newReference,
		}
		MapScannedFiles[newReference] = child

		// Si es un directorio, construye el árbol recursivamente
		if file.IsDir() {
			subtree, err := buildTree(childPath)
			if err != nil {
				return []DirectoryNode{}, err
			}
			child.Files = append(child.Files, subtree...)
		}

		files_and_dirs = append(files_and_dirs, child)

	}
	// nodeFather.Files = files_and_dirs

	return files_and_dirs, nil
}
