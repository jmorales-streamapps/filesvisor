package functions

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"reflect"
	"time"
)

func GenString(longitud int) string {
	rand.Seed(time.Now().UnixNano())
	//rand.Seed(1223232)

	caracteres := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	resultado := make([]byte, longitud)
	for i := 0; i < longitud; i++ {
		resultado[i] = caracteres[rand.Intn(len(caracteres))]
	}

	return string(resultado)
}

func PrintJson(data any) {
	// Convierte la estructura del árbol a formato JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}

	// Imprime el JSON resultante
	fmt.Println(string(jsonData))
}

func ReadInfo(completeRoute string) fs.FileInfo {

	// Obtener información del archivo
	info, err := os.Stat(completeRoute)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("El archivo no existe.")
		} else {
			fmt.Println("Error al obtener información del archivo:", err)
		}
	}
	return info

	// Imprimir información del archivo
	fmt.Println("Nombre:", info.Name())
	fmt.Println("Tamaño:", info.Size(), "bytes")
	fmt.Println("Es directorio:", info.IsDir())
	fmt.Println("Modo de permisos:", info.Mode())
	fmt.Println("Tiempo de modificación:", info.ModTime())
	return info
}

func IsNumeric(value interface{}) bool {
	kind := reflect.TypeOf(value).Kind()
	return kind >= reflect.Int && kind <= reflect.Float64
}
