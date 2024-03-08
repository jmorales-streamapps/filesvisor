package functions

import (
	"math/rand"
	"time"
)

func GenString(longitud int) string {
	rand.Seed(time.Now().UnixNano())

	caracteres := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	resultado := make([]byte, longitud)
	for i := 0; i < longitud; i++ {
		resultado[i] = caracteres[rand.Intn(len(caracteres))]
	}

	return string(resultado)
}
