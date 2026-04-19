package entanglement

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
)

//go:embed static/*
var static embed.FS

//go:embed templates/*
var templates embed.FS

func SetupTemplateFuncs(templ *template.Template) *template.Template {
	templ = templ.Funcs(template.FuncMap{
		"entanglementframe": func(web Session, frame string) string {
			return web.CreateSubFrame(frame).GenerateToken()
		},
	})

	fsys, err := fs.Sub(templates, "templates")
	if err != nil {
		panic(err)
	}

	return template.Must(templ.ParseFS(fsys, "*.html"))
}

func SetupServeStatic(mux *http.ServeMux) {
	fsys, err := fs.Sub(static, "static")
	if err != nil {
		panic(err)
	}

	mux.Handle("GET /entanglement/static/", http.StripPrefix("/entanglement/static/", http.FileServer(http.FS(fsys))))
}
