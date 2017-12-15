package main

import (
	"html/template"
	"net/http"
	"sandbox"
)

type Browser struct {
	Name   string
	Engine string
}

type ViewData struct {
	Browsers []Browser
}

func main() {
	http.HandleFunc("/view/", sandbox.ViewHandler)
	http.HandleFunc("/edit/", sandbox.EditHandler)
	http.HandleFunc("/save/", sandbox.SaveHandler)
	http.HandleFunc("/range/", rangeHandler)
	http.ListenAndServe(":8080", nil)
}

func wrapBracket(s string, ltag string, rtag string) string {
	return ltag + s + rtag
}

func rangeHandler(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"bracket": func(label string) string {
			return wrapBracket(label, "{", "}")
		},
	}

	t, err := template.New("range.html").Funcs(funcMap).ParseFiles("tmpl/range.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, ViewData{
		Browsers: []Browser{
			Browser{"Google Chrome", "Blink"},
			Browser{"Internet Explorer", "Trident"},
			Browser{"Firefox", "Gecko"},
			Browser{"Safari", "WebKit"},
		},
	})
}
