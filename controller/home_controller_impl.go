package controller

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HomeControllerImpl struct {
}

func NewHomeController() HomeController {
	return &HomeControllerImpl{}
}

//go:embed views/*
var views embed.FS

func (controller *HomeControllerImpl) Home(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	dirrectory, err := fs.Sub(views, "views")
	if err != nil {
		panic(err)
	}
	fileServer := http.FileServer(http.FS(dirrectory))

	t := template.Must(template.ParseFS(views, "views/*"))
	t.ExecuteTemplate(writer, "index.html", fileServer)

}
