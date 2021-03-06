package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		//tmpl, err := template.New("test").Parse(doc)
		templates := template.New("template")
		templates.New("test").Parse(doc)
		templates.New("header").Parse(header)
		context := Context{
			[3]string{"Lemon", "Orange", "Apple"},
			"the title",
		}
		templates.Lookup("test").Execute(w, context)

	})

	http.ListenAndServe(":8000", nil)
}

const doc = `
<!DOCTYPE html>
<html>
{{template "header" .Title}}
<body>
<h1>List of Fruit</h1>
<ul>
{{range .Fruit}}
<li>{{.}}</li>
{{end}}
</ul>
</body>
</html>`

const header = `
<head><title>{{.}}</title></head>
`

type Context struct {
	Fruit [3]string
	Title string
}
