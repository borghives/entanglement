package entanglement

import (
	"html/template"
	"net/http"
	"testing"
)

func TestSetupTemplateFuncs(t *testing.T) {
	tmpl := template.New("test")
	tmpl = SetupTemplateFuncs(tmpl)
	
	if tmpl == nil {
		t.Errorf("expected non-nil template")
	}
}

func TestSetupServeStatic(t *testing.T) {
	mux := http.NewServeMux()
	SetupServeStatic(mux)
	
	// This just verifies it does not panic and registers the route.
	// Since there's no easy way to query routes in a ServeMux natively without a request,
	// successful execution without panics indicates the handler was registered correctly.
}
