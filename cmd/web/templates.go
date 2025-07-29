package main

import (
	"github.com/Akkshatt/go_snippet_box/internals/models"
	"html/template"
	"path/filepath"
	"time"

	"io/fs"
	"github.com/Akkshatt/go_snippet_box/ui"
)

type templateData struct {
	CurrentYear     int
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
	 User *models.User

}



func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}





var functions = template.FuncMap{
	"humanDate": humanDate,
}


// func newTemplateCache() (map[string]*template.Template, error) {
// 	cache := map[string]*template.Template{}

// 	// Load all page templates from disk
// 	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, page := range pages {
// 		name := filepath.Base(page)

// 		ts, err := template.New(name).Funcs(functions).ParseFiles(
// 			"./ui/html/base.tmpl.html",
// 			page,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}

// 		// Parse all partials into the template
// 		_, err = ts.ParseGlob("./ui/html/partials/*.tmpl.html")
// 		if err != nil {
// 			return nil, err
// 		}

// 		cache[name] = ts
// 	}

// 	return cache, nil
// }





func newTemplateCache() (map[string]*template.Template, error) {
 cache := map[string]*template.Template{}
 // Use fs.Glob() to get a slice of all filepaths in the ui.Files embedded
 
 pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl.html")
 if err != nil {
 return nil, err
 }
 for _, page := range pages {
 name := filepath.Base(page)
 

 patterns := []string{
 "html/base.tmpl.html",
 "html/partials/*.tmpl.html",
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

