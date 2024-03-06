package handlefiles

import (
	"fmt"
	"os"
	"path/filepath"
)

// Node representa un nodo en la estructura del árbol.
type DirectoryNode struct {
	Name  string          `json:"name"`
	IsDir bool            `json:"isDir"`
	Files []DirectoryNode `json:"children,omitempty"`
}

func ReeadDirectory() *DirectoryNode {
	// Ruta del directorio raíz que deseas explorar
	rootDir := "./"

	// Llama a la función para construir el árbol
	tree, err := buildTree(rootDir)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return &tree

	// Convierte la estructura del árbol a formato JSON
	//jsonData, err := json.MarshalIndent(tree, "", "  ")
	//if err != nil {
	//	fmt.Println("Error al convertir a JSON:", err)
	//	return
	//}

	// Imprime el JSON resultante
	//fmt.Println(string(jsonData))
}

func buildTree(rootPath string) (DirectoryNode, error) {
	var result DirectoryNode

	// Lee el contenido del directorio
	files, err := os.ReadDir(rootPath)
	if err != nil {
		return result, err
	}

	// Recorre los archivos/directorios y construye el árbol
	for _, file := range files {
		childPath := filepath.Join(rootPath, file.Name())

		child := DirectoryNode{
			Name:  file.Name(),
			IsDir: file.IsDir(),
		}

		// Si es un directorio, construye el árbol recursivamente
		if file.IsDir() {
			subtree, err := buildTree(childPath)
			if err != nil {
				return result, err
			}
			child.Files = append(child.Files, subtree)
		}

		// Agrega el nodo al resultado
		result.Files = append(result.Files, child)
	}

	return result, nil
}
