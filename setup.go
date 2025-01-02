package entanglement

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/borghives/entanglement/concept"
	"github.com/borghives/websession"
)

//go:embed static/*
var static embed.FS

//go:embed templates/*
var templates embed.FS

func SetupEntanglementTemplates(templ *template.Template) *template.Template {
	templ = templ.Funcs(template.FuncMap{
		"entanglementframe": func(e concept.Entanglement, frame string) string {
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

	mux.Handle("GET /entanglement/static/", http.StripPrefix("/entanglement/static/", http.FileServer(http.FS(fsys))))
}

func CreateEntanglement(session *websession.Session) *concept.Entanglement {
	return &concept.Entanglement{
		SystemSession: session,
		Nonce:         websession.GetRandomHexString(),
		Token:         session.GenerateSessionToken(),
	}
}

func CreateEntanglementWithNonce(session *websession.Session, nonce string) *concept.Entanglement {
	return &concept.Entanglement{
		SystemSession: session,
		Nonce:         nonce,
		Token:         session.GenerateSessionToken(),
	}
}

func CreateEntanglementWithNonceAndToken(session *websession.Session, nonce string, token string) *concept.Entanglement {
	return &concept.Entanglement{
		SystemSession: session,
		Nonce:         nonce,
		Token:         token,
	}
}
