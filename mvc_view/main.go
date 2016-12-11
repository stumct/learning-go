package main

import (
	"net/http"
	"os"
	"text/template"

	"github.com/stumct/learn/mvc_view/controllers"
)

func main() {
	templates := populateTemplates()
	controllers.Register(templates)
	http.ListenAndServe(":8000", nil)

}

func populateTemplates() *template.Template {

	// create a new template object which will store all of our templates
	result := template.New("templates")
	// define the templates folder and open that foler for reading
	basePath := "templates"
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close() // defer the closing of the folder (cleanup!)

	// read all of the files in the folder
	templatePathsRaw, _ := templateFolder.Readdir(-1)

	// for each file read append its path to templatePaths if it is not a directory
	templatePaths := new([]string)
	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
		}
	}

	// Parse the templates in templatePaths
	result.ParseFiles(*templatePaths...)

	// return the templates
	return result
}
