package app

import (
	"html/template"
	"os"
)

var funcmaps = template.FuncMap{
	"env": os.Getenv,
}
