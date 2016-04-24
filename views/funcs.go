package views

import (
	"html/template"
	"strings"
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"title": strings.Title,
		"lower": strings.ToLower,
	}
}
