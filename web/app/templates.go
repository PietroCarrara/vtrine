package app

import (
	"html/template"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
)

var templates *template.Template

func loadTemplates() {
	templates = template.New("root")
	templates.Funcs(funcmaps)

	_, err := templates.ParseGlob("web/templates/*.go.html")
	fail(err)

	if gin.Mode() != gin.ReleaseMode {
		go watchTemplates()
	}
}

func watchTemplates() {
	watcher, err := fsnotify.NewWatcher()
	fail(err)
	defer watcher.Close()

	watcher.Add("web/templates")

	for {
		select {
		case event, ok := <-watcher.Events:
			if ok {
				if event.Op&fsnotify.Write != 0 {
					loadTemplates()
				}
			}
		}
	}
}