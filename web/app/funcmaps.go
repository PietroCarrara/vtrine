package app

import (
	"html/template"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
)

var funcmaps = template.FuncMap{
	"env": os.Getenv,
	"url": func(s string) template.URL {
		return template.URL(s)
	},
	"byte": func(s uint64) string {
		return humanize.Bytes(s)
	},
	"rand": func() string {
		return strconv.Itoa(rand.Int())
	},
	"cat": func(s ...string) string {
		return strings.Join(s, "")
	},
}
