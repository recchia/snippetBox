package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"github.com/recchia/snippetbox/internal/models/mysql"
	"github.com/recchia/snippetbox/ui"
)

type templateData struct {
	CSRFToken       string
	CurrentYear     int
	Flash           string
	Form            any //*forms.Form
	IsAuthenticated bool
	Snippet         mysql.Snippet
	Snippets        []mysql.Snippet
}

func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.UTC().Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.html")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"html/layout.html",
			"html/partials/*.html",
			page,
		}

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}
		
		cache[name] = ts
	}

	return cache, nil
}
