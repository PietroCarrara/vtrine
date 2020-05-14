package app

import (
	"html/template"
	"os"
)

var funcmaps = template.FuncMap{
	"env": os.Getenv,
	"url": func(s string) template.URL {
		return template.URL(s)
	},
}
