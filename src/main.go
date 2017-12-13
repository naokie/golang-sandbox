package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

type Browser struct {
	Name   string
	Engine string
}

type ViewData struct {
	Browsers []Browser
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/range/", rangeHandler)
	http.ListenAndServe(":8080", nil)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[6:]
	p, _ := loadPage(title)
	t, err := template.ParseFiles("tmpl/view.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[6:]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	t, err := template.ParseFiles("tmpl/edit.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[6:]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func rangeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("range.html").ParseFiles("tmpl/range.html")
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
