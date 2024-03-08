package handlefiles

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Jon-MC-dev/files_copy/functions"
)

// Node representa un nodo en la estructura del árbol.
type DirectoryNode struct {
	Name      string          `json:"name"`
	IsDir     bool            `json:"isDir"`
	Files     []DirectoryNode `json:"children,omitempty"`
	Reference string
}

func ReeadDirectory() *DirectoryNode {
	// Ruta del directorio raíz que deseas explorar
	rootDir := "./test"

	// Llama a la función para construir el árbol
	var result DirectoryNode = DirectoryNode{Name: "Root", IsDir: true, Reference: functions.GenString(10)}

	tree, err := buildTree(rootDir)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	result.Files = append(result.Files, tree...)

	// Convierte la estructura del árbol a formato JSON
	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return nil
	}

	// Imprime el JSON resultante
	fmt.Println(string(jsonData))

	return &result

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

		child := DirectoryNode{
			Name:      file.Name(),
			IsDir:     file.IsDir(),
			Reference: functions.GenString(10),
		}

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
