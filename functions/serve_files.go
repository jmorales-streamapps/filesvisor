package functions

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"

	"github.com/Masterminds/sprig"
)

func ServeHTML(w http.ResponseWriter, r *http.Request, filename string, content embed.FS, data any) {
	file, err := content.ReadFile("templates/" + filename)
	if err != nil {
		http.Error(w, "No se pudo leer el archivo HTML", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.New(filename).Delims("[[", "]]").Funcs(sprig.FuncMap()).Parse(string(file))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error al analizar el archivo HTML", http.StatusInternalServerError)
		return
	}

	// data := struct {
	// }{}

	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error al renderizar el HTML", http.StatusInternalServerError)
		return
	}
}
