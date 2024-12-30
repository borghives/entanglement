package entanglement

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/borghives/entanglement/concept"
)

//go:embed static/*
var static embed.FS

//go:embed templates/*
var templates embed.FS

func SetupEntanglmentTemplates(templ *template.Template) *template.Template {
	templ = templ.Funcs(template.FuncMap{
		"entangledsytem": func(e concept.Entanglement, frame string) string {
			return e.CreatSubFrame(frame).GenerateToken()
		},
	})

	fsys, err := fs.Sub(templates, "templates")
	if err != nil {
		panic(err)
	}

	return template.Must(templ.ParseFS(fsys, "*.html"))
}

func SetupEntanglementRoutes(mux *http.ServeMux) {
	fsys, err := fs.Sub(static, "static")
	if err != nil {
		panic(err)
	}

	mux.Handle("/entanglement/", http.StripPrefix("/entanglement/", http.FileServer(http.FS(fsys))))
}
