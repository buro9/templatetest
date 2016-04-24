package views

import (
	"html/template"
	"sync"
)

const (
	templateHome  = "home"
	templateForum = "forum"
)

var (
	parseTemplatesOnce sync.Once
	templates          map[string]*template.Template
)

func ParseTemplates() {
	parseTemplatesOnce.Do(func() {
		templates = make(map[string]*template.Template)

		templates[templateHome] = template.Must(
			template.New("base.html").Funcs(FuncMap()).ParseFiles(
				templatePath+"base.html",
				templatePath+"home.html",
			),
		)

		templates[templateForum] = template.Must(
			template.New("base.html").Funcs(FuncMap()).ParseFiles(
				templatePath+"base.html",
				templatePath+"forum.html",
			),
		)
	})
}
