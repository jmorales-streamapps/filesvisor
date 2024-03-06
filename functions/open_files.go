package functions

import (
	"embed"
	"html/template"
	"net/http"
)

func ServeHTML(w http.ResponseWriter, r *http.Request, filename string, content embed.FS) {
	file, err := content.ReadFile("templates/" + filename)
	if err != nil {
		http.Error(w, "No se pudo leer el archivo HTML", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.New(filename).Parse(string(file))
	if err != nil {
		http.Error(w, "Error al analizar el archivo HTML", http.StatusInternalServerError)
		return
	}

	data := struct {
		// Puedes agregar datos adicionales que quieras pasar al HTML aquí
	}{}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error al renderizar el HTML", http.StatusInternalServerError)
		return
	}
}
